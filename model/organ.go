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
	UserId int64 `gorm:"not null" json:"userid"`
	//组织电话
	OrganAddress string `gorm:"type:varchar(100);not null" json:"organaddress"`
	//组织电话
	OrganPhone string `gorm:"type:varchar(20);not null" json:"organphone"`
	//组织描述
	Desc string `gorm:"type:varchar(100)" json:"desc"`
	//组织状态 0为审核中，1为已通过
	Approve int64 `gorm:"type:int;not null" json:"approve"`
	//组织状态 0为正常，1为禁用
	State   int64 `gorm:"type:int;not null" json:"state"`
	Deleted int   `gorm:"type:int;not null" json:"delete"`

	CreateTime time.Time
}

//用来检查参数是否正确
func (c *Organ) CheckOrgan() error {
	db.MysqlDB.AutoMigrate(Organ{})
	if !db.MysqlDB.HasTable(Organ{}) {
		if db.MysqlDB.HasTable(Organ{}) {
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
	db.MysqlDB.Where("organ_name=? AND approve=?", c.OrganName, 1).First(&or).Count(&count)

	if count != 0 {
		return errors.New("组织名已存在")
	}

	return nil
}

func (g *Organ) GetAllOrgans(code int) (organs []Organ, err error) {
	//select * from Organ
	db.MysqlDB.Where("approve=?", code).Find(&organs)
	return
}

func (g *Organ) GetOrgansByName(code int, organName string) (organs []Organ, err error) {
	//select * from Organ
	db.MysqlDB.Where("code = ? AND organ_name LIKE ?", code, organName).Find(&organs)
	return
}

//更新审核状态
func (o *Organ) UpdateApprove(id int64, code int64) error {
	var organ Organ
	if o.Approve != 0 {
		return errors.New("该用户已经过审核")
	}
	db.MysqlDB.Model(&organ).Where("id=?", id).Update("approve", code)
	return nil
}

//根据 user_id 获取 organ_id
func (o *Organ) GetOrganId(user_id int64) int64 {
	var organ Organ
	db.MysqlDB.Where("user_id=?", user_id).First(&organ)
	return organ.ID
}
