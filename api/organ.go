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

func AddOrgan(c *gin.Context) {
	//此处我试用的json格式传输数据，通常情况下这里应该是form形式的提交，这里是我个人的喜好，比较方便，可灵活变换
	or := &model.Organ{}
	err := c.BindJSON(or)
	if err != nil {
		fmt.Println(fmt.Errorf("organ BindJSON err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "数据格式不正确",
		})
		return
	}
	err = or.CheckOrgan()
	if err != nil {
		fmt.Println(fmt.Errorf("organ Check err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	or.CreateTime = time.Now()
	result := db.MysqlDB.Create(or)
	if result.Error != nil {
		fmt.Println(fmt.Errorf("organ Insert err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func GetOrganApproveing(c *gin.Context) {
	var err error
	or := &model.Organ{}
	err = c.BindJSON(or)
	if err != nil {
		fmt.Println(fmt.Errorf("organ BindJSON err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "数据格式不正确",
		})
		return
	}
	organName := or.OrganName
	var organs []model.Organ
	if organName == "" {
		organs, err = or.GetAllOrgans(0)
		if err != nil {
			fmt.Println(fmt.Errorf("GetAllorgans err : %v", err))
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
			return
		}
	} else {
		organs, err = or.GetOrgansByName(0, organName)
		if err != nil {
			fmt.Println(fmt.Errorf("GetAllorgansByName err : %v", err))
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
			return
		}
	}
	organsMarshal, err := json.Marshal(organs)
	if err != nil {
		fmt.Println(fmt.Errorf("Marshal organ err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	fmt.Println(string(organsMarshal))
	c.JSON(http.StatusOK, gin.H{
		"msg":  "ok",
		"data": string(organsMarshal),
	})

}

func GetOrganApproved(c *gin.Context) {
	var err error
	or := &model.Organ{}
	err = c.BindJSON(or)
	if err != nil {
		fmt.Println(fmt.Errorf("organ BindJSON err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "数据格式不正确",
		})
		return
	}
	organName := or.OrganName
	var organs []model.Organ
	if organName == "" {
		organs, err = or.GetAllOrgans(1)
		if err != nil {
			fmt.Println(fmt.Errorf("GetAllorgans err : %v", err))
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
			return
		}
	} else {
		organs, err = or.GetOrgansByName(1, organName)
		if err != nil {
			fmt.Println(fmt.Errorf("GetAllorgansByName err : %v", err))
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
			return
		}
	}
	organsMarshal, err := json.Marshal(organs)
	if err != nil {
		fmt.Println(fmt.Errorf("Marshal organ err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	fmt.Println(string(organsMarshal))
	c.JSON(http.StatusOK, gin.H{
		"msg":  "ok",
		"data": string(organsMarshal),
	})

}

// 接受组织
func AcceptOrgan(c *gin.Context) {
	//此处我试用的json格式传输数据，通常情况下这里应该是form形式的提交，这里是我个人的喜好，比较方便，可灵活变换
	or := &model.Organ{}
	err := c.BindJSON(or)
	if err != nil {
		fmt.Println(fmt.Errorf("organ BindJSON err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "数据格式不正确",
		})
		return
	}
	//更改用户表用户角色
	u := &model.TUser{}
	err = u.UpdateRole(or.UserId)
	if err != nil {
		fmt.Println(fmt.Errorf("organ update approve err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	//更改审核状态
	err = or.UpdateApprove(or.ID, 1)
	if err != nil {
		fmt.Println(fmt.Errorf("organ update approve err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

// 拒绝组织
func RefuseOrgan(c *gin.Context) {
	//此处我试用的json格式传输数据，通常情况下这里应该是form形式的提交，这里是我个人的喜好，比较方便，可灵活变换
	or := &model.Organ{}
	err := c.BindJSON(or)
	if err != nil {
		fmt.Println(fmt.Errorf("organ BindJSON err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "数据格式不正确",
		})
		return
	}
	//更改审核状态
	err = or.UpdateApprove(or.ID, 2)
	if err != nil {
		fmt.Println(fmt.Errorf("organ update approve err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

//获取组织（管理）
func GetOrgan(c *gin.Context) {
	var err error
	o := &model.Organ{}

	var organs []model.Organ
	// 获取管理的物资，0表示正常使用
	organs, err = o.GetAllOrgansManage()
	if err != nil {
		fmt.Println(fmt.Errorf("GetAllorgans err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	organsMarshal, err := json.Marshal(organs)
	if err != nil {
		fmt.Println(fmt.Errorf("Marshal good err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	fmt.Println(string(organsMarshal))
	c.JSON(http.StatusOK, gin.H{
		"msg":  "ok",
		"data": string(organsMarshal),
	})

}

//禁用组织（管理）
func BanOrgan(c *gin.Context) {
	var err error
	o := &model.Organ{}
	err = c.BindJSON(o)
	if err != nil {
		fmt.Println(fmt.Errorf("organ BindJSON err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "数据格式不正确",
		})
		return
	}
	err = o.UpdateState(1)
	if err != nil {
		fmt.Println(fmt.Errorf("Share update approve err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
