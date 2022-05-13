package model

import (
	"errors"
	"fmt"
	"time"
	"we-share-go/db"
)

type Askhelp struct {
	ID int64 `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	//用户id
	UserId      int64  `gorm:"type:int;not null" json:"userid"`
	GoodName    string `gorm:"type:varchar(20);not null" json:"goodname"`
	Desc        string `gorm:"type:varchar(100)" json:"desc"`
	UserPhone   string `gorm:"type:varchar(20);not null" json:"userphone"`
	UserAddress string `gorm:"type:varchar(100)" json:"usersddress"`
	CreateTime  time.Time
}

//用来检查参数是否正确
func (c *Askhelp) CheckAskhelp() error {
	db.MysqlDB.AutoMigrate(&Askhelp{})
	if !db.MysqlDB.HasTable(&Askhelp{}) {
		if db.MysqlDB.HasTable(&Askhelp{}) {
			fmt.Println("askhelp表创建成功")
		} else {
			fmt.Println("askhelp表创建失败")
		}
	} else {
		fmt.Println("表已存在")
	}
	if c.GoodName == "" {
		return errors.New("物资名不能为空")
	}
	if c.UserPhone == "" {
		return errors.New("联系电话不能为空")
	}
	if c.UserAddress == "" {
		return errors.New("联系电话不能为空")
	}

	return nil
}

//查询所有
func (as *Askhelp) GetAllaskhelps() (askhelps []Askhelp, err error) {
	//select * from Askhelp
	db.MysqlDB.Find(&askhelps)
	return
}

//查询所有物资（系统）
func (as *Askhelp) GetAskhelpTotal() (total int64, err error) {
	//select * from Good
	var askhelp Askhelp
	db.MysqlDB.Model(&askhelp).Count(&total)
	return
}
