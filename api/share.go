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

func GetBorrow(c *gin.Context) {
	var err error
	u := &model.TUser{}
	err = c.BindJSON(u)
	if err != nil {
		fmt.Println(fmt.Errorf("user BindJSON err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "数据格式不正确",
		})
		return
	}
	s := &model.Share{}
	var shares []model.Share
	shares, err = s.GetSharesBorrow(u.ID)
	if err != nil {
		fmt.Println(fmt.Errorf("GetAllShares err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	SharesMarshal, err := json.Marshal(shares)
	if err != nil {
		fmt.Println(fmt.Errorf("Marshal askhelp err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	fmt.Println(string(SharesMarshal))
	c.JSON(http.StatusOK, gin.H{
		"msg":  "ok",
		"data": string(SharesMarshal),
	})

}
