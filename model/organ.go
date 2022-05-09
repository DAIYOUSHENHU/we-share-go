package model

import (
	"errors"
	"fmt"
	"time"
	"we-share-go/db"
)

type Organ struct {
	ID        int64  `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	OrganName string `gorm:"type:varchar(20);not null" json:"organname"`
	//用户id
	UserId  string `gorm:"type:varchar(20);not null" json:"userid"`
	Deleted int    `gorm:"type:int;not null" json:"delete"`

	CreateTime time.Time
}

//用来检查参数是否正确
func (c *Organ) CheckOrgan() error {
	db.MysqlDB.AutoMigrate(&Organ{})
	if !db.MysqlDB.HasTable(&Organ{}) {
		if db.MysqlDB.HasTable(&Organ{}) {
			fmt.Println("organ表创建成功")
		} else {
			fmt.Println("organ表创建失败")
		}
	} else {
		fmt.Println("表已存在")
	}
	if c.OrganName == "" {
		return errors.New("用户名不能为空")
	}

	var or Organ
	var count int
	// 根据用户名查询
	db.MysqlDB.Where("organ_name=?", c.OrganName).First(&or).Count(&count)

	if count != 0 {
		return errors.New("组织名已存在")
	}

	return nil
}
