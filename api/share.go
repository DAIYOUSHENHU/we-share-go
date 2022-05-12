package api

import (
	"fmt"
	"net/http"
	"time"
	"we-share-go/db"
	"we-share-go/model"

	"github.com/gin-gonic/gin"
)

func AddShareGood(c *gin.Context) {
	s := &model.Share{}
	err := c.BindJSON(s)
	if err != nil {
		fmt.Println(fmt.Errorf("good BindJSON err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "数据格式不正确",
		})
		return
	}
	fmt.Println(s)
	err = s.CheckShare()
	if err != nil {
		fmt.Println(fmt.Errorf("good Check err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	s.CreateTime = time.Now()
	result := db.MysqlDB.Create(s)
	if result.Error != nil {
		fmt.Println(fmt.Errorf("good Insert err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
