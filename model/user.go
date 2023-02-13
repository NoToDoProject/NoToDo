package model

import "time"

// User 用户结构体
type User struct {
	Uid             int       `json:"uid" bson:"uid"`                             // 用户 ID，系统自动生成
	Username        string    `json:"username" bson:"username"`                   // 用户名，唯一，不可修改
	Password        []byte    `json:"password" bson:"password"`                   // 密码
	Nickname        string    `json:"nickname" bson:"nickname"`                   // 显示昵称
	Email           string    `json:"email" bson:"email"`                         // 邮箱，唯一
	Disabled        bool      `json:"disabled" bson:"disabled"`                   // 是否禁用，禁用后无法登录
	NeedEmailVerify bool      `json:"need_email_verify" bson:"need_email_verify"` // 是否需要邮箱验证
	EmailVerified   bool      `json:"email_verified" bson:"email_verified"`       // 邮箱是否已验证
	IsAdmin         bool      `json:"is_admin" bson:"is_admin"`                   // 是否为管理员
	RegisterTime    time.Time `json:"register_time" bson:"register_time"`         // 注册时间
}

// IsUserExist 判断用户是否存在
type IsUserExist struct {
	Username string `form:"username" bson:"username" binding:"required"` // 用户名
}

// UserLogin 用户登录结构体
type UserLogin struct {
	Username string `json:"username" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 密码
}

// UserWithPassword 用户登录结构体
type UserWithPassword struct {
	Username string `bson:"username"` // 用户名
	Password []byte `bson:"password"` // 密码
}

// UserRegister 用户注册结构体
type UserRegister struct {
	Username string `json:"username" bson:"username" binding:"required"` // 用户名
	Password string `json:"password" bson:"password" binding:"required"` // 密码
	Email    string `json:"email" bson:"email" binding:"required,email"` // 邮箱
}
