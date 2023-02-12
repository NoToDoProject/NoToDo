package server

import (
	"github.com/NoToDoProject/NoToDo/common/response"
	"github.com/gin-gonic/gin"
)

// Information 服务器信息
func Information(c *gin.Context) {
	nc := response.ContextEx{Context: c}
	nc.Success()
}
