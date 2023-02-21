package model

import "time"

type User struct {
	Uid             int       `json:"uid" bson:"uid"`           // unique, auto increment
	Username        string    `json:"username" bson:"username"` // unique, can't be changed
	Password        []byte    `json:"-" bson:"password"`
	Nickname        string    `json:"nickname" bson:"nickname"` // string for display
	Email           string    `json:"email" bson:"email"`       // unique
	Disabled        bool      `json:"-" bson:"disabled"`        // unable to login if disabled
	NeedEmailVerify bool      `json:"-" bson:"need_email_verify"`
	EmailVerified   bool      `json:"-" bson:"email_verified"`
	IsAdmin         bool      `json:"-" bson:"is_admin"` // admin can do anything
	RegisterTime    time.Time `json:"registerTime" bson:"register_time"`
}

// IsUserExist check if user exist query
type IsUserExist struct {
	Username string `form:"username" bson:"username" binding:"required"`
}

// UserLogin login request body
type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserWithPassword username and password
type UserWithPassword struct {
	Username string `bson:"username"`
	Password []byte `bson:"password"`
}

// UserRegister user register request body
type UserRegister struct {
	Username string `json:"username" bson:"username" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
	Email    string `json:"email" bson:"email" binding:"required,email"`
}
