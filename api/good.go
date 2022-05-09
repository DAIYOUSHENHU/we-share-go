package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"we-share-go/db"
	"we-share-go/model"

	"github.com/gin-gonic/gin"
)

func AddShareGood(c *gin.Context) {
	g := &model.Good{}
	err := c.BindJSON(g)
	if err != nil {
		fmt.Println(fmt.Errorf("askgood BindJSON err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "数据格式不正确",
		})
		return
	}
	err = g.CheckGood()
	if err != nil {
		fmt.Println(fmt.Errorf("askgood Check err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	g.CreateTime = time.Now()
	result := db.MysqlDB.Create(g)
	if result.Error != nil {
		fmt.Println(fmt.Errorf("Register Insert err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func GetShareGood(c *gin.Context) {
	var err error
	g := &model.Good{}
	err = c.BindJSON(g)
	if err != nil {
		fmt.Println(fmt.Errorf("askgood BindJSON err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "数据格式不正确",
		})
		return
	}

	goodName := g.GoodName
	var goods []model.Good

	if goodName == "" {
		goods, err = g.GetAllGoods()
		if err != nil {
			fmt.Println(fmt.Errorf("GetAllGoods err : %v", err))
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
			return
		}
	} else {
		goods, err = g.GetGoodsByName(goodName)
		if err != nil {
			fmt.Println(fmt.Errorf("GetAllGoodsByName err : %v", err))
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
			return
		}
	}
	goodsMarshal, err := json.Marshal(goods)
	if err != nil {
		fmt.Println(fmt.Errorf("Marshal good err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	fmt.Println(string(goodsMarshal))
	c.JSON(http.StatusOK, gin.H{
		"msg":  "ok",
		"data": string(goodsMarshal),
	})

}
