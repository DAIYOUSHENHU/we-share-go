package model

import (
	"errors"
	"fmt"
	"time"
	"we-share-go/db"
)

type Share struct {
	ID int64 `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	//物资ID
	GoodId int64 `gorm:"type:int;not null" json:"goodid"`
	//物资名
	GoodName string `gorm:"type:varchar(20);not null" json:"goodname"`
	//用户ID
	UserId int64 `gorm:"type:int;not null" json:"userid"`
	//用户名
	UserName string `gorm:"type:varchar(20);not null" json:"username"`
	//组织ID
	OrganId int64 `gorm:"type:int;not null" json:"organid"`
	//用户电话
	UserPhone string `gorm:"type:varchar(20);not null" json:"userphone"`
	//物品描述
	Note string `gorm:"type:varchar(100)" json:"note"`
	//物资审核状态 0为审核中，1为已通过
	Approve int64 `gorm:"type:int;not null" json:"approve"`
	//是否已删除.0为正常，1为已删除
	Deleted int `gorm:"type:int;not null" json:"delete"`

	CreateTime time.Time
}

//用来检查参数是否正确
func (s *Share) CheckShare() error {
	db.MysqlDB.AutoMigrate(Share{})
	if !db.MysqlDB.HasTable(Share{}) {
		if db.MysqlDB.HasTable(Share{}) {
			fmt.Println("Share表创建成功")
		} else {
			fmt.Println("Share表创建失败")
		}
	} else {
		fmt.Println("表已存在")
	}
	if s.GoodId == 0 {
		return errors.New("物资id不能为空")
	}
	if s.UserId == 0 {
		return errors.New("用户id不能为空")
	}
	if s.OrganId == 0 {
		return errors.New("用户id不能为空")
	}
	if s.UserPhone == "" {
		return errors.New("用户联系电话不能为空")
	}

	return nil
}

//查询所有物资（共享）
func (s *Share) GetAllShares(organ_id int64, code int) (shares []Share, err error) {
	//select * from Good
	db.MysqlDB.Where("organ_id=? AND approve=?", organ_id, code).Find(&shares)
	return
}

//更新审核状态
func (s *Share) UpdateApprove(id int64, code int64) error {
	var share Share
	if s.Approve != 0 {
		return errors.New("该物资已经过审核")
	}
	db.MysqlDB.Model(&share).Where("id=?", id).Update("approve", code)
	return nil
}

//查询借用的物资
func (s *Share) GetSharesBorrow(user_id int64) (shares []Share, err error) {
	//select * from Good
	db.MysqlDB.Where("user_id=? AND approve=?", user_id, 1).Find(&shares)
	return
}

//查询所有共享（系统）
func (s *Share) GetShareTotal() (total int64, err error) {
	//select * from Good
	var share Share
	db.MysqlDB.Model(&share).Where("approve=?", 1).Count(&total)
	return
}
