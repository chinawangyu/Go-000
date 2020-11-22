package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 野协程panic问题示范
func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/error", func(c *gin.Context) {

		//！！！此处同样存在越界错误, 因为是一个新的协程环境, panic后直接退出进程, ping接口无法再响应
		go func() {
			/*defer func() {  //只要在新启动的gorouting中声明recover捕获即可正常运行
				if err := recover(); err != nil {
					fmt.Println(reflect.TypeOf(err))
				}
			}()*/

			var errStr []string
			fmt.Println(errStr[1])
		}()

		//当此处存在错误时, 因为存在recovery,所以不会影响ping的正常使用
		var errStr []string
		c.JSON(200, gin.H{
			"message": errStr[1],
		})
	})
	r.Run(":8080")

}
