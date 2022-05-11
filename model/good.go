package model

import (
	"errors"
	"fmt"
	"time"
	"we-share-go/db"
)

type Good struct {
	ID int64 `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	//物资名
	GoodName string `gorm:"type:varchar(20);not null" json:"goodname"`
	//物品描述
	Desc string `gorm:"type:varchar(100)" json:"desc"`
	//用户ID
	UserId int64 `gorm:"type:int;not null" json:"userid"`
	//用户电话
	UserPhone string `gorm:"type:varchar(20);not null" json:"userphone"`
	//组织ID
	OrganId int64 `gorm:"type:int;not null" json:"organid"`
	//组织状态 0为审核中，1为已通过
	Approve int64 `gorm:"type:int;not null" json:"approve"`
	//物资共享状态 0为审核阶段  1为共享状态
	ShareState int64 `gorm:"type:int;not null" json:"sharestate"`
	//是否已删除.0为正常，1为已删除
	Deleted int `gorm:"type:int;not null" json:"delete"`

	CreateTime time.Time
}

func (g *Good) CheckGood() error {
	db.MysqlDB.AutoMigrate(&Good{})
	if !db.MysqlDB.HasTable(&Good{}) {
		if db.MysqlDB.HasTable(&Good{}) {
			fmt.Println("good表创建成功")
		} else {
			fmt.Println("good表创建失败")
		}
	} else {
		fmt.Println("表已存在")
	}
	if g.GoodName == "" {
		return errors.New("物品名不能为空")
	}
	if g.UserId == 0 {
		return errors.New("用户ID不能为空")
	}
	if g.UserPhone == "" {
		return errors.New("联系电话不能为空")
	}
	if g.OrganId == 0 {
		return errors.New("组织ID不能为空")
	}

	return nil
}

//查询所有
func (g *Good) GetAllGoods() (goods []Good, err error) {
	//select * from Good
	db.MysqlDB.Find(&goods)
	return
}

func (g *Good) GetGoodsByName(goodName string) (goods []Good, err error) {
	//select * from Good
	db.MysqlDB.Where("good_name LIKE ?", goodName).Find(&goods)
	return
}
