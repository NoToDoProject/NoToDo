package user

import (
	"github.com/NoToDoProject/NoToDo/common"
	"github.com/NoToDoProject/NoToDo/common/response"
	"github.com/NoToDoProject/NoToDo/database"
	userDb "github.com/NoToDoProject/NoToDo/database/user"
	"github.com/NoToDoProject/NoToDo/model"
	"github.com/gin-gonic/gin"
	"time"
)

// isUserExist check user exist
func isUserExist(c *gin.Context) {
	nc := response.ContextEx{Context: c}

	// get params
	var isExist model.IsUserExist
	_ = nc.BindQuery(&isExist)

	nc.Success(userDb.IsUserExist(isExist))
}

// register user register
func register(c *gin.Context) {
	nc := response.ContextEx{Context: c}

	// check register enable
	if !database.Config.CanRegister {
		nc.RegisterDisabled()
		return
	}

	// todo wip
	if database.Config.NeedRegisterEmailVerification {
		nc.Failure(response.Error, "Register need email verification")
		return
	}

	// get params
	var registerInfo model.UserRegister
	_ = nc.BindJSON(&registerInfo)

	// check user exist
	var isExistInfo model.IsUserExist
	common.CopyStruct(&registerInfo, &isExistInfo)
	if userDb.IsUserExist(isExistInfo) {
		nc.Failure(response.Error, "User already exists")
		return
	}

	// check email exist
	if userDb.IsEmailExist(registerInfo.Email) {
		nc.Failure(response.Error, "Email already exists")
		return
	}

	// save user
	var newUser model.User
	common.CopyStruct(&registerInfo, &newUser)
	var userLogin model.UserLogin
	common.CopyStruct(&registerInfo, &userLogin)
	newUser.Uid = -1 // allocate after email verify
	newUser.Password = common.EncryptPassword(common.MakeNewPassword(userLogin))
	newUser.Nickname = newUser.Username
	newUser.Disabled = false
	newUser.NeedEmailVerify = database.Config.NeedRegisterEmailVerification
	newUser.EmailVerified = false
	newUser.IsAdmin = false
	newUser.RegisterTime = time.Now()

	if !userDb.AddUser(newUser) {
		nc.Failure(response.Error, "Add user failed")
		return
	}

	if !userDb.AllocateUid(newUser) {
		nc.Failure(response.Error, "Allocate uid failed")
		return
	}

	nc.Success("register success")
}

// info get user info
func info(c *gin.Context) {
	nc := response.ContextEx{Context: c}
	user := nc.GetUser()
	nc.Success(user)
}
