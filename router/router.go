package router

import (
	"we-share-go/api"
	"we-share-go/mid"

	"github.com/gin-gonic/gin"
)

func RegRouter(g *gin.Engine) {
	//第一组api接口 例如：http://localhost:8080/g1/hello1
	g1 := g.Group("/api/v1")

	//中间件
	g1.Use(mid.MidCors)

	g1.GET("/hello1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "Hello g1",
		})
	})

	g1.POST("/register", api.Register)
	g1.POST("/login", api.Login)

}
