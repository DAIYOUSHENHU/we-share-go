package api

import (
	"fmt"
	"net/http"
	"time"

	"we-share-go/db"
	"we-share-go/model"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	//此处我试用的json格式传输数据，通常情况下这里应该是form形式的提交，这里是我个人的喜好，比较方便，可灵活变换
	u := &model.TUser{}
	err := c.BindJSON(u)
	if err != nil {
		fmt.Println(fmt.Errorf("Register BindJSON err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "数据格式不正确",
		})
		return
	}
	err = u.Check()
	if err != nil {
		fmt.Println(fmt.Errorf("Register Check err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	u.CreateTime = time.Now()
	result := db.MysqlDB.Create(u)
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
