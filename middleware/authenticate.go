package middleware

import (
	"fmt"
	"github.com/NoToDoProject/NoToDo/common"
	"github.com/NoToDoProject/NoToDo/common/response"
	userDb "github.com/NoToDoProject/NoToDo/database/user"
	"github.com/NoToDoProject/NoToDo/model"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var identityKey = common.IdentityKeyInContext // jwt的payload中的key
var JWTMiddleware = &jwt.GinJWTMiddleware{
	Realm:            "notodo",                                                             // 作用域
	SigningAlgorithm: "HS256",                                                              // 签名算法
	Key:              []byte("0xk9013kum0uh0cfu0nv-1v2-ic0--cu812u-bc-182c916et2yx87t1b2"), // 签名
	//KeyFunc:               nil, 			// 签名函数
	Timeout:       time.Hour * 24, // 过期时间
	MaxRefresh:    time.Hour * 24, // 允许刷新时间
	Authenticator: Login,          // 登录函数
	Authorizator: func(data any, c *gin.Context) bool { // 判断用户是否有权限
		if user, ok := data.(model.User); ok && user.Uid != 0 {
			c.Set(identityKey, user)
			return true // 登录就有权限
		}
		return false
	},
	PayloadFunc: func(data any) jwt.MapClaims { // 设置jwt的payload，data是登录函数返回的数据
		user := data.(model.User)
		return jwt.MapClaims{
			"uid": user.Uid,
		}
	},
	Unauthorized: func(c *gin.Context, code int, message string) { // 未登录时返回的信息，message是登录函数返回的err信息
		nc := response.ContextEx{Context: c}
		log.Warnf("Unauthorized: %+v", message)
		nc.Unauthorized()
	},
	LoginResponse: func(c *gin.Context, code int, message string, time time.Time) { // 登录响应
		nc := response.ContextEx{Context: c}
		if code != http.StatusOK {
			log.Panicf("LoginResponse: %+v", message)
		}
		log.Debugf("claims: %+v", jwt.ExtractClaims(c))
		nc.Success(gin.H{
			"token":  message,
			"expire": time.Format("2006-01-02 15:04:05"),
		})
	},
	//LogoutResponse:        nil, // 登出响应函数
	RefreshResponse: func(c *gin.Context, code int, message string, time time.Time) { // 刷新响应
		nc := response.ContextEx{Context: c}
		if code != http.StatusOK {
			log.Panicf("RefreshResponse: %+v", message)
		}
		nc.Success(gin.H{
			"token":  message,
			"expire": time.Format("2006-01-02 15:04:05"),
		})
	},
	IdentityHandler: func(context *gin.Context) any { // 请求中获取用户信息
		data := jwt.ExtractClaims(context)
		uid := int(data["uid"].(float64)) // 这里不知道为什么是float64
		user, err := userDb.GetUserByUid(uid)
		if err != nil {
			log.Panicf("IdentityHandler: %+v", err)
		}
		return user
	},
	IdentityKey:   identityKey,             // payload在context中的key
	TokenLookup:   "header: Authorization", // 从哪里获取token
	TokenHeadName: "Bearer",                // token前缀
	TimeFunc:      time.Now,                // 获取当前时间的函数
	//HTTPStatusMessageFunc: nil, // 设置http状态码对应的信息
	SendCookie:        false, // 是否发送cookie
	SendAuthorization: false, // 在header中设置Authorization
	DisabledAbort:     false, // 是否禁止context.Abort
}
var AuthMiddleware, err = jwt.New(JWTMiddleware)
var MiddleFunc = AuthMiddleware.MiddlewareFunc()

func init() {
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
}

// Login 登录
func Login(c *gin.Context) (any, error) {
	nc := response.ContextEx{Context: c}

	// 获取登录信息
	var loginInfo model.UserLogin
	_ = nc.BindJSON(&loginInfo)

	// 检查用户是否存在
	var userExist model.IsUserExist
	common.CopyStruct(&loginInfo, &userExist)
	userGotten, err := userDb.GetUser(userExist)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	// 检查用户是否被禁用
	if userGotten.Disabled {
		return nil, fmt.Errorf("user disabled")
	}

	// 检查密码是否正确
	password := common.MakeNewPassword(loginInfo)
	if !common.ComparePassword(userGotten.Password, password) {
		return nil, fmt.Errorf("password error")
	}

	// 登录成功
	return userGotten, nil
}
