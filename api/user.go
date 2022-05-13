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
	err = u.CheckRegister()
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

func Login(c *gin.Context) {
	u := &model.TUser{}
	err := c.BindJSON(u)
	//检查数据是否成功绑定
	if err != nil {
		fmt.Println(fmt.Errorf("login BindJSON err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "数据格式不正确",
		})
		return
	}
	// 检查login数据是否合法
	err = u.CheckLogin()
	if err != nil {
		fmt.Println(fmt.Errorf("login Check err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	//检查角色
	role := u.CheckRole()

	userInfo := u.GetUserInfo()
	// 用户禁用
	if userInfo.State == 1 {
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "Forbidden",
		})
		return
	}
	fmt.Println("80")
	userInfo.Pwd = "******"
	userInfoMarshal, err := json.Marshal(userInfo)
	if err != nil {
		fmt.Println(fmt.Errorf("Marshal userInfo err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	// 根据角色返回数据
	c.JSON(http.StatusOK, gin.H{
		"msg":      "ok",
		"token":    "login",
		"role":     role,
		"userInfo": string(userInfoMarshal),
	})

}

//获取用户（管理）
func GetUser(c *gin.Context) {
	var err error
	u := &model.TUser{}

	var users []model.TUser
	// 获取管理的物资，0表示正常使用
	users, err = u.GetAllUsersManage()
	if err != nil {
		fmt.Println(fmt.Errorf("GetAllusers err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	usersMarshal, err := json.Marshal(users)
	if err != nil {
		fmt.Println(fmt.Errorf("Marshal good err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	fmt.Println(string(usersMarshal))
	c.JSON(http.StatusOK, gin.H{
		"msg":  "ok",
		"data": string(usersMarshal),
	})

}

//禁用用户（管理）
func BanUser(c *gin.Context) {
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
	err = u.UpdateState(1)
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

type SysInfoRes struct {
	UserTotal  int64 `json:"usertotal"`
	OrganTotal int64 `json:"organtotal"`
	GoodTotal  int64 `json:"goodtotal"`
	ShareTotal int64 `json:"sharetotal"`
}

//系统信息
func SysInfo(c *gin.Context) {
	var err error
	u := &model.TUser{}
	userTotal, err := u.GetUserTotal()

	o := &model.Organ{}
	organTotal, err := o.GetOrganTotal()

	g := &model.Good{}
	goodTotal, err := g.GetGoodTotal()

	s := &model.Share{}
	shareTotal, err := s.GetShareTotal()

	var sysinfo SysInfoRes
	sysinfo.UserTotal = userTotal
	sysinfo.OrganTotal = organTotal
	sysinfo.GoodTotal = goodTotal
	sysinfo.ShareTotal = shareTotal

	sysinfoMarshal, err := json.Marshal(sysinfo)
	if err != nil {
		fmt.Println(fmt.Errorf("Marshal good err : %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	fmt.Println(string(sysinfoMarshal))
	c.JSON(http.StatusOK, gin.H{
		"msg":  "ok",
		"data": string(sysinfoMarshal),
	})
}
