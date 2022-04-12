package model

import (
	"errors"
	"fmt"
	"regexp"
	"time"
	"we-share-go/db"
)

type TUser struct {
	Id int64
	//昵称
	NickName string
	//用户名
	UserName string
	//密码
	Pwd string
	//是否已删除.0为正常，1为已删除
	Deleted int

	CreateTime time.Time
}

//用来检查参数是否正确
func (c *TUser) Check() error {
	if c.UserName == "" {
		return errors.New("用户名不能为空")
	}
	if len(c.UserName) < 5 {
		return errors.New("用户名必须大于4位")
	}
	user := &TUser{
		UserName: c.UserName,
	}
	get := db.MysqlDB.Limit(1).Find(&user)
	if get.Error != nil {
		fmt.Println(fmt.Errorf("Check Get err : %v", get.Error))
		return errors.New("服务器繁忙请稍后重试")
	}
	if get != nil {
		return errors.New("用户名已存在")
	}

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
