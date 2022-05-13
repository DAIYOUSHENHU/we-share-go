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

func AddLog(c *gin.Context) {
	l := &model.Loginfo{}
	err := c.BindJSON(l)
	if err != nil {
		fmt.Println(fmt.Errorf("good BindJSON err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "数据格式不正确",
		})
		return
	}
	err = l.CheckLoginfo()
	if err != nil {
		fmt.Println(fmt.Errorf("LogInfo Check err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	l.CreateTime = time.Now()
	result := db.MysqlDB.Create(l)
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

func GetLog(c *gin.Context) {
	var err error
	l := &model.Loginfo{}

	var logs []model.Loginfo
	// 获取管理的物资，0表示正常使用
	logs, err = l.GetAllLogs()
	if err != nil {
		fmt.Println(fmt.Errorf("GetAllusers err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	logsMarshal, err := json.Marshal(logs)
	if err != nil {
		fmt.Println(fmt.Errorf("Marshal good err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "ok",
		"data": string(logsMarshal),
	})
}
