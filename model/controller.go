package model

import "github.com/gin-gonic/gin"

// Controller interface
type Controller interface {
	InitRouter(*gin.Engine) // add router
}
