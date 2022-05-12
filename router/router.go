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
	//新增请求物资信息
	u.POST("/addAskhelp", api.Askhelp)
	u.POST("/getAskhelp", api.GetAskhelp)

	// 组织api接口
	or := g.Group("/organ")
	//中间件
	or.Use(mid.MidCors)
	or.POST("/addOrgan", api.AddOrgan)
	or.POST("/getOrganApproveing", api.GetOrganApproveing)
	or.POST("/getOrganApproved", api.GetOrganApproved)
	// or.GET("/getOrganReject", api.GetOrganReject)
	or.POST("/acceptOrgan", api.AcceptOrgan)
	or.POST("/refuseOrgan", api.RefuseOrgan)

	// 物资api接口
	g1 := g.Group("/good")
	//中间件
	g1.Use(mid.MidCors)
	g1.POST("/addGood", api.AddGood)
	g1.POST("/getGoodApproveing", api.GetGoodApproveing)
	g1.POST("/getGoodApproved", api.GetGoodApproved)
	// g1.POST("/getGood", api.GetGood)
	g1.POST("/acceptGood", api.AcceptGood)
	g1.POST("/refuseGood", api.RefuseGood)

	g1.POST("/addShareGood", api.AddShareGood)
	g1.POST("/getShareGood", api.GetShareGood)
	g1.POST("/getShareApproveing", api.GetShareApproveing)
	g1.POST("/getShareApproved", api.GetShareApproved)
	g1.POST("/acceptShare", api.AcceptShare)
	g1.POST("/refuseShare", api.RefuseShare)
}
