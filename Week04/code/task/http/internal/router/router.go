package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"task/http/internal/controller"
)

func Init(g *gin.Engine) (err error) {
	if g == nil {
		err = fmt.Errorf("nil gin engine")
		return err
	}

	//g.Use(middle.Log)

	user := g.Group("/user")
	{
		user.GET("/list", controller.GetUserList)
	}

	return nil
}
