package model

import (
	"errors"
	"fmt"
	"regexp"
	"time"
	"we-share-go/db"
)

type TUser struct {
	ID int64 `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	//昵称
	NickName string `gorm:"type:varchar(20);not null" json:"nickname"`
	//用户名
	UserName string `gorm:"type:varchar(20);not null" json:"username"`
	//密码
	Pwd string `gorm:"type:varchar(20);not null" json:"pwd"`
	//是否已删除.0为正常，1为已删除
	Deleted int `gorm:"type:int;not null" json:"delete"`

	CreateTime time.Time
}

//用来检查参数是否正确
func (c *TUser) Check() error {
	if !db.MysqlDB.HasTable(&TUser{}) {
		db.MysqlDB.AutoMigrate(&TUser{})
		if db.MysqlDB.HasTable(&TUser{}) {
			fmt.Println("user表创建成功")
		} else {
			fmt.Println("user表创建失败")
		}
	} else {
		fmt.Println("表已存在")
	}
	if c.UserName == "" {
		return errors.New("用户名不能为空")
	}
	if len(c.UserName) < 5 {
		return errors.New("用户名必须大于4位")
	}
	fmt.Println("check user")
	fmt.Println(c)
	fmt.Println(c.UserName)

	var user TUser
	var count int
	fmt.Println("49: ", user)
	fmt.Println("username: ", user.UserName)
	db.MysqlDB.Where("user_name=?", c.UserName).First(&user).Count(&count)
	fmt.Println("52: ", user)
	fmt.Println("55: ", count)

	if count != 0 {
		return errors.New("用户名已存在")
	}

	// db.MysqlDB.Where("username=?", c.UserName).First(&user)

	if c.Pwd == "" {
		return errors.New("密码不能为空")
	}

	compile, err := regexp.Compile(`^(.{6,16}[^0-9]*[^A-Z]*[^a-z]*[a-zA-Z0-9]*)$`)
	if err != nil {
		fmt.Println(fmt.Errorf("Check regexp err : %v", err))
		return errors.New("服务器繁忙请稍后重试")
	}
	b := compile.MatchString(c.Pwd)
	if !b {
		return errors.New("密码不符合规则！密码应为6-16位（可以包含字母、数字、下划线）")
	}

	return nil
}
