package model

import "github.com/gin-gonic/gin"

type Controller interface {
	InitRouter(*gin.Engine)
}
