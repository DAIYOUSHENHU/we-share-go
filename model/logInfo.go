package model

import (
	"errors"
	"fmt"
	"time"
	"we-share-go/db"
)

type Loginfo struct {
	ID int64 `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	//组织状态 0为正常，1为禁用
	UserId int64 `gorm:"type:int;not null" json:"userid"`
	//用户名
	UserName string `gorm:"type:varchar(20);not null" json:"username"`
	//角色，0为普通用户，1为组织，2为管理员
	Role int `gorm:"type:int;not null" json:"role"`
	//操作描述
	Desc string `gorm:"type:varchar(100)" json:"desc"`
	//请求类型
	ReqType string `gorm:"type:varchar(100)" json:"reqtype"`

	CreateTime time.Time `json:"createtime"`
}

func (l *Loginfo) CheckLoginfo() error {
	db.MysqlDB.AutoMigrate(&Loginfo{})
	if !db.MysqlDB.HasTable(&Loginfo{}) {
		if db.MysqlDB.HasTable(&Loginfo{}) {
			fmt.Println("loginfo表创建成功")
		} else {
			fmt.Println("loginfo表创建失败")
		}
	} else {
		fmt.Println("表已存在")
	}
	if l.UserId == 0 {
		return errors.New("用户id不能为空")
	}
	if l.UserName == "" {
		return errors.New("用户名不能为空")
	}

	if l.Desc == "" {
		return errors.New("操作描述不能为空")
	}
	return nil
}

func (l *Loginfo) GetAllLogs() (logs []Loginfo, err error) {
	db.MysqlDB.Limit(100).Order("create_time desc").Find(&logs)
	return
}
