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

func AddGood(c *gin.Context) {
	g := &model.Good{}
	err := c.BindJSON(g)
	if err != nil {
		fmt.Println(fmt.Errorf("good BindJSON err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "数据格式不正确",
		})
		return
	}
	err = g.CheckGood()
	if err != nil {
		fmt.Println(fmt.Errorf("good Check err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	g.CreateTime = time.Now()
	result := db.MysqlDB.Create(g)
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

// func GetShareGood(c *gin.Context) {
// 	var err error
// 	g := &model.Good{}
// 	err = c.BindJSON(g)
// 	if err != nil {
// 		fmt.Println(fmt.Errorf("good BindJSON err : %v", err))
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"msg": "数据格式不正确",
// 		})
// 		return
// 	}

// 	goodName := g.GoodName
// 	var goods []model.Good

// 	if goodName == "" {
// 		goods, err = g.GetAllGoods()
// 		if err != nil {
// 			fmt.Println(fmt.Errorf("GetAllGoods err : %v", err))
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"msg": err.Error(),
// 			})
// 			return
// 		}
// 	} else {
// 		goods, err = g.GetGoodsByName(goodName)
// 		if err != nil {
// 			fmt.Println(fmt.Errorf("GetAllGoodsByName err : %v", err))
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"msg": err.Error(),
// 			})
// 			return
// 		}
// 	}
// 	goodsMarshal, err := json.Marshal(goods)
// 	if err != nil {
// 		fmt.Println(fmt.Errorf("Marshal good err : %v", err))
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"msg": err.Error(),
// 		})
// 		return
// 	}
// 	fmt.Println(string(goodsMarshal))
// 	c.JSON(http.StatusOK, gin.H{
// 		"msg":  "ok",
// 		"data": string(goodsMarshal),
// 	})

// }

func GetGoodApproveing(c *gin.Context) {
	var err error
	u := &model.TUser{}
	err = c.BindJSON(u)
	if err != nil {
		fmt.Println(fmt.Errorf("organ BindJSON err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "数据格式不正确",
		})
		return
	}
	o := &model.Organ{}
	organ_id := o.GetOrganId(u.ID)
	fmt.Println(organ_id)
	g := &model.Good{}
	var goods []model.Good
	// 获取审核中的物资，0表示待审核
	goods, err = g.GetAllGoods(organ_id, 0)
	if err != nil {
		fmt.Println(fmt.Errorf("GetAllgoods err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
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

func GetGoodApproved(c *gin.Context) {
	var err error
	u := &model.TUser{}
	err = c.BindJSON(u)
	if err != nil {
		fmt.Println(fmt.Errorf("organ BindJSON err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "数据格式不正确",
		})
		return
	}
	o := &model.Organ{}
	organ_id := o.GetOrganId(u.ID)
	g := &model.Good{}
	var goods []model.Good
	// 获取审核中的物资，1表示审核通过
	goods, err = g.GetAllGoods(organ_id, 1)
	if err != nil {
		fmt.Println(fmt.Errorf("GetAllgoods err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
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

// 接受物资
func AcceptGood(c *gin.Context) {
	//此处我试用的json格式传输数据，通常情况下这里应该是form形式的提交，这里是我个人的喜好，比较方便，可灵活变换
	g := &model.Good{}
	err := c.BindJSON(g)
	if err != nil {
		fmt.Println(fmt.Errorf("Good BindJSON err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "数据格式不正确",
		})
		return
	}

	//更改审核状态
	err = g.UpdateApprove(g.ID, 1)
	if err != nil {
		fmt.Println(fmt.Errorf("Good update approve err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

// 拒绝物资
func RefuseGood(c *gin.Context) {
	//此处我试用的json格式传输数据，通常情况下这里应该是form形式的提交，这里是我个人的喜好，比较方便，可灵活变换
	g := &model.Good{}
	err := c.BindJSON(g)
	if err != nil {
		fmt.Println(fmt.Errorf("Good BindJSON err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "数据格式不正确",
		})
		return
	}
	//更改审核状态
	err = g.UpdateApprove(g.ID, 2)
	if err != nil {
		fmt.Println(fmt.Errorf("Good update approve err : %v", err))
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
		fmt.Println(fmt.Errorf("organ BindJSON err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "数据格式不正确",
		})
		return
	}
	var goods []model.Good
	// 获取审核中的物资，0表示待审核
	goods, err = g.GetAllShareGoods()
	if err != nil {
		fmt.Println(fmt.Errorf("GetAllgoods err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
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

func GetShareApproveing(c *gin.Context) {
	var err error
	u := &model.TUser{}
	err = c.BindJSON(u)
	if err != nil {
		fmt.Println(fmt.Errorf("organ BindJSON err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "数据格式不正确",
		})
		return
	}
	o := &model.Organ{}
	organ_id := o.GetOrganId(u.ID)
	fmt.Println(organ_id)
	s := &model.Share{}
	var shares []model.Share
	// 获取审核中的共享物资，0表示待审核
	shares, err = s.GetAllShares(organ_id, 0)
	if err != nil {
		fmt.Println(fmt.Errorf("GetAllgoods err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	sharesMarshal, err := json.Marshal(shares)
	if err != nil {
		fmt.Println(fmt.Errorf("Marshal good err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	fmt.Println(string(sharesMarshal))
	c.JSON(http.StatusOK, gin.H{
		"msg":  "ok",
		"data": string(sharesMarshal),
	})

}

func GetShareApproved(c *gin.Context) {
	var err error
	u := &model.TUser{}
	err = c.BindJSON(u)
	if err != nil {
		fmt.Println(fmt.Errorf("organ BindJSON err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "数据格式不正确",
		})
		return
	}
	o := &model.Organ{}
	organ_id := o.GetOrganId(u.ID)
	fmt.Println(organ_id)
	s := &model.Share{}
	var shares []model.Share
	// 获取审核中的共享物资，1表示审核通过
	shares, err = s.GetAllShares(o.ID, 1)
	if err != nil {
		fmt.Println(fmt.Errorf("GetAllgoods err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	sharesMarshal, err := json.Marshal(shares)
	if err != nil {
		fmt.Println(fmt.Errorf("Marshal good err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	fmt.Println(string(sharesMarshal))
	c.JSON(http.StatusOK, gin.H{
		"msg":  "ok",
		"data": string(sharesMarshal),
	})

}

// 接受共享
func AcceptShare(c *gin.Context) {
	//此处我试用的json格式传输数据，通常情况下这里应该是form形式的提交，这里是我个人的喜好，比较方便，可灵活变换
	s := &model.Share{}
	err := c.BindJSON(s)
	if err != nil {
		fmt.Println(fmt.Errorf("Share BindJSON err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "数据格式不正确",
		})
		return
	}

	//更改审核状态
	err = s.UpdateApprove(s.ID, 1)
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

// 拒绝共享
func RefuseShare(c *gin.Context) {
	s := &model.Share{}
	err := c.BindJSON(s)
	if err != nil {
		fmt.Println(fmt.Errorf("Share BindJSON err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "数据格式不正确",
		})
		return
	}
	//更改审核状态
	err = s.UpdateApprove(s.ID, 2)
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
