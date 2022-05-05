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
	//用户名
	UserName string `gorm:"type:varchar(20);not null" json:"username"`
	//密码
	Pwd string `gorm:"type:varchar(20);not null" json:"pwd"`
	//角色，0为普通用户，1为组织，2为管理员
	Role int `gorm:"type:int;not null" json:"role"`
	//是否已删除.0为正常，1为已删除
	Deleted int `gorm:"type:int;not null" json:"delete"`

	CreateTime time.Time
}

//用来检查参数是否正确
func (c *TUser) CheckRegister() error {
	db.MysqlDB.AutoMigrate(&TUser{})
	if !db.MysqlDB.HasTable(&TUser{}) {
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

	if c.Pwd == "" {
		return errors.New("密码不能为空")
	}

	var user TUser
	var count int
	// 根据用户名查询
	db.MysqlDB.Where("user_name=?", c.UserName).First(&user).Count(&count)

	if count != 0 {
		return errors.New("用户名已存在")
	}
	// 检验密码合法性
	compile, err := regexp.Compile(`^(.{6,16}[^0-9]*[^A-Z]*[^a-z]*[a-zA-Z0-9]*)$`)
	if err != nil {
		return errors.New("服务器繁忙请稍后重试")
	}
	b := compile.MatchString(c.Pwd)
	if !b {
		return errors.New("密码不符合规则！密码应为6-16位（可以包含字母、数字、下划线）")
	}

	return nil
}

//用来检查参数是否正确
func (c *TUser) CheckLogin() error {
	// user表自动迁移
	db.MysqlDB.AutoMigrate(&TUser{})
	if !db.MysqlDB.HasTable(&TUser{}) {
		fmt.Println("表不存在")
	}
	if c.UserName == "" {
		return errors.New("用户名不能为空")
	}
	if c.Pwd == "" {
		return errors.New("密码不能为空")
	}

	var user TUser
	var count int
	db.MysqlDB.Where("user_name=?", c.UserName).First(&user).Count(&count)

	if count == 0 {
		return errors.New("用户不存在")
	}

	if user.UserName != c.UserName || user.Pwd != c.Pwd {
		return errors.New("密码错误")
	}

	return nil
}

func (c *TUser) CheckRole() int {
	var user TUser
	var count int
	db.MysqlDB.Where("user_name=?", c.UserName).First(&user).Count(&count)
	// 返回用户角色
	return user.Role
}
