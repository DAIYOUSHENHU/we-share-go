package router

import (
	"we-share-go/api"
	"we-share-go/mid"

	"github.com/gin-gonic/gin"
)

func RegRouter(g *gin.Engine) {
	// 用户api接口
	u := g.Group("/api/v1")

	//中间件
	u.Use(mid.MidCors)

	u.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "Hello g1",
		})
	})

	u.POST("/register", api.Register)
	u.POST("/login", api.Login)

	// 物资api接口
	g1 := g.Group("/good")
	//中间件
	g1.Use(mid.MidCors)

	g1.POST("/addShareGood", api.AddShareGood)
	g1.POST("/getShareGood", api.GetShareGood)

}
