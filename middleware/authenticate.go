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

var identityKey = common.IdentityKeyInContext // key for user info in context
var JWTMiddleware = &jwt.GinJWTMiddleware{
	Realm:            "notodo",
	SigningAlgorithm: "HS256",
	Key:              []byte("0xk9013kum0uh0cfu0nv-1v2-ic0--cu812u-bc-182c916et2yx87t1b2"),
	//KeyFunc:               nil, 			// signature key function
	Timeout:       time.Hour * 24, // token valid time
	MaxRefresh:    time.Hour * 24, // can refresh token time
	Authenticator: Login,          // login function
	Authorizator: func(data any, c *gin.Context) bool { // check user can access this resource
		if user, ok := data.(model.User); ok && user.Uid != 0 {
			c.Set(identityKey, user)
			return true // authorized once login, can use RBAC or ABAC if needed
		}
		return false
	},
	PayloadFunc: func(data any) jwt.MapClaims { // store user info in jwt payload
		user := data.(model.User)
		return jwt.MapClaims{
			"uid": user.Uid,
		}
	},
	Unauthorized: func(c *gin.Context, code int, message string) { // response for unauthorized, message is error message
		nc := response.ContextEx{Context: c}
		log.Warnf("Unauthorized: %+v", message)
		nc.Unauthorized()
	},
	LoginResponse: func(c *gin.Context, code int, message string, time time.Time) { // response for login
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
	//LogoutResponse:        nil, // response for logout
	RefreshResponse: func(c *gin.Context, code int, message string, time time.Time) { // response for refresh token
		nc := response.ContextEx{Context: c}
		if code != http.StatusOK {
			log.Panicf("RefreshResponse: %+v", message)
		}
		nc.Success(gin.H{
			"token":  message,
			"expire": time.Format("2006-01-02 15:04:05"),
		})
	},
	IdentityHandler: func(context *gin.Context) any { // get user info from context
		data := jwt.ExtractClaims(context)
		uid := int(data["uid"].(float64)) // WHY float64?
		user, err := userDb.GetUserByUid(uid)
		if err != nil {
			log.Panicf("IdentityHandler: %+v", err)
		}
		return user
	},
	IdentityKey:   identityKey,             // where to store user info in context
	TokenLookup:   "header: Authorization", // where to get token
	TokenHeadName: "Bearer",                // token suffix
	TimeFunc:      time.Now,                // func to get current time
	//HTTPStatusMessageFunc: nil, // func to get http status message
	SendCookie:        false, // set cookie
	SendAuthorization: false, // set Authorization header
	DisabledAbort:     false, // will not abort if the token is expired
}
var AuthMiddleware, err = jwt.New(JWTMiddleware)
var MiddleFunc = AuthMiddleware.MiddlewareFunc()

func init() {
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
}

// Login user login
func Login(c *gin.Context) (any, error) {
	nc := response.ContextEx{Context: c}

	// get login info
	var loginInfo model.UserLogin
	_ = nc.BindJSON(&loginInfo)

	// check user exist
	var userExist model.IsUserExist
	common.CopyStruct(&loginInfo, &userExist)
	userGotten, err := userDb.GetUser(userExist)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	// check user disabled
	if userGotten.Disabled {
		return nil, fmt.Errorf("user disabled")
	}

	// check password
	password := common.MakeNewPassword(loginInfo)
	if !common.ComparePassword(userGotten.Password, password) {
		return nil, fmt.Errorf("password error")
	}

	return userGotten, nil
}

// IsAdminMiddleware check if user is admin
func IsAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		nc := response.ContextEx{Context: c}
		user := nc.GetUser()
		if !user.IsAdmin {
			nc.Unauthorized()
			c.Abort()
			return
		}
		c.Next()
	}
}
