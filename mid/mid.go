package mid

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func MidAuth(c *gin.Context) {
// 	//获取请求的ip地址
// 	ip := c.ClientIP()
// 	//如果请求地址来源不正确，那么阻止这个请求继续
// 	if ip != "baidu.com" {
// 		println("ip 地址不正确")
// 		c.Abort()
// 		return
// 	}
// 	// 处理请求
// 	c.Next()
// }

func MidCors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Credentials", "true")

	//放行所有OPTIONS方法
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	// 处理请求
	c.Next()
}
