package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task/http/di"
)

func GetUserList(c *gin.Context) {
	userService:= di.NewUserService()
	c.String(http.StatusOK, userService.GetUserList())
}
