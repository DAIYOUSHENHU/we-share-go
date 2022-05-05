package model

import (
	"errors"
	"fmt"
	"time"
	"we-share-go/db"
)

type TAskGood struct {
	ID int64 `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	//用户名
	GoodName string `gorm:"type:varchar(20);not null" json:"goodname"`
	//用户ID
	UserId int64 `gorm:"type:int;not null" json:"userid"`
	//用户电话
	UserPhone string `gorm:"type:varchar(20);not null" json:"userphone"`
	//物品描述
	Desc string `gorm:"type:varchar(20)" json:"desc"`
	//是否已删除.0为正常，1为已删除
	Deleted int `gorm:"type:int;not null" json:"delete"`

	CreateTime time.Time
}

func (g *TAskGood) CheckAskGood() error {
	db.MysqlDB.AutoMigrate(&TAskGood{})
	if !db.MysqlDB.HasTable(&TAskGood{}) {
		if db.MysqlDB.HasTable(&TAskGood{}) {
			fmt.Println("user表创建成功")
		} else {
			fmt.Println("user表创建失败")
		}
	} else {
		fmt.Println("表已存在")
	}
	if g.GoodName == "" {
		return errors.New("物品名不能为空")
	}

	if g.UserPhone == "" {
		return errors.New("联系电话不能为空")
	}

	return nil
}

//查询所有
func (g *TAskGood) GetAllGoods() (askgoods []TAskGood, err error) {
	//select * from TAskGood
	db.MysqlDB.Find(&TAskGood{})
	return
}

func (g *TAskGood) GetGoodsByName(goodName string) (askgoods []TAskGood, err error) {
	//select * from TAskGood
	db.MysqlDB.Where("good_name LIKE ?", goodName).Find(askgoods)
	return
}
