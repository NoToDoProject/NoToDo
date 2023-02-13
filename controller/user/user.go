package user

import (
	"fmt"
	"github.com/NoToDoProject/NoToDo/common"
	"github.com/NoToDoProject/NoToDo/common/response"
	"github.com/NoToDoProject/NoToDo/database"
	"github.com/NoToDoProject/NoToDo/database/user"
	"github.com/NoToDoProject/NoToDo/model"
	"github.com/gin-gonic/gin"
	"time"
)

// IsUserExist 判断用户是否存在
func IsUserExist(c *gin.Context) {
	nc := response.ContextEx{Context: c}

	// 获取参数
	var isExist model.IsUserExist
	_ = nc.BindQuery(&isExist)

	nc.Success(user.IsUserExist(isExist))
}

// Login 登录
func Login(c *gin.Context) {
	nc := response.ContextEx{Context: c}

	// 获取登录信息
	var loginInfo model.UserLogin
	_ = nc.BindJSON(&loginInfo)

	// 检查用户是否存在
	var userExist model.IsUserExist
	common.CopyStruct(&loginInfo, &userExist)
	userGotten, err := user.GetUser(userExist)
	if err != nil {
		nc.LoginError()
		return
	}

	// 检查密码是否正确
	password := makeNewPassword(loginInfo)
	if !common.ComparePassword(userGotten.Password, password) {
		nc.LoginError()
		return
	}

	// 登录成功
	nc.Success("sample_token")
}

// Register 注册
func Register(c *gin.Context) {
	nc := response.ContextEx{Context: c}

	// 检查是否允许注册
	if !database.Config.CanRegister {
		nc.RegisterDisabled()
		return
	}

	// 获取注册信息
	var registerInfo model.UserRegister
	_ = nc.BindJSON(&registerInfo)

	// 检查用户是否存在
	var isExistInfo model.IsUserExist
	common.CopyStruct(&registerInfo, &isExistInfo)
	if user.IsUserExist(isExistInfo) {
		nc.Failure(response.Error, "User already exists")
		return
	}

	// 注册成功
	var newUser model.User
	common.CopyStruct(&registerInfo, &newUser)
	var userLogin model.UserLogin
	common.CopyStruct(&registerInfo, &userLogin)
	newUser.Uid = -1 // 验证通过后分配uid
	newUser.Password = common.EncryptPassword(makeNewPassword(userLogin))
	newUser.Nickname = newUser.Username
	newUser.Disabled = false
	newUser.NeedEmailVerify = database.Config.NeedRegisterEmailVerification
	newUser.EmailVerified = false
	newUser.IsAdmin = false
	newUser.RegisterTime = time.Now()

	if !user.AddUser(newUser) {
		nc.Failure(response.Error, "Add user failed")
		return
	}

	// todo 邮箱唯一约束
	nc.Success("register success")
}

// makeNewPassword 生成新密码
func makeNewPassword(user model.UserLogin) string {
	return fmt.Sprintf("%s%s%s", user.Password, "notodo", user.Username)
}
