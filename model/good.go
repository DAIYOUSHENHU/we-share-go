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
	//物资审核状态 0为审核中，1为已通过
	Approve int64 `gorm:"type:int;not null" json:"approve"`
	//物资共享状态 0为审核阶段  1为共享状态
	ShareState int64 `gorm:"type:int;not null" json:"sharestate"`
	//是否已删除.0为正常，1为已删除
	Deleted int `gorm:"type:int;not null" json:"delete"`

	CreateTime time.Time
}

type ShowShare struct {
	ID int64 `json:"id"`
	//物资名
	GoodName string `json:"goodname"`
	//物品描述
	Desc string `json:"desc"`
	//用户ID
	UserId int64 `json:"userid"`
	//用户名
	UserName string `json:"username"`
	//用户电话
	UserPhone string `json:"userphone"`
	//组织ID
	OrganId int64 `json:"organid"`
	//组织名
	OrganName string `json:"organname"`
	//组织地址
	OrganAddress string `json:"organaddress"`
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

//查询所有物资（共享）
// func (g *Good) GetAllShareGoods() (goods []Good, err error) {
// 	//select * from Good
// 	db.MysqlDB.Where("approve=?", 1).Find(&goods)
// 	return
// }
//查询所有物资（共享）
func (g *Good) GetAllShareGoods() (res []ShowShare, err error) {
	//select * from Good
	// res := &[]ShowShare{}
	db.MysqlDB.Table("good").Select("good.id, good.good_name, good.desc, good.user_id, good.organ_id, organ.organ_name, organ.organ_phone, organ.organ_address").Joins("left join organ ON organ.id = good.organ_id").Where("good.approve=? AND good.share_state=? AND organ.approve=?", 1, 0, 1).Scan(&res)
	return
}

//查询所有物资（审核）
func (g *Good) GetAllGoods(organ_id int64, code int) (goods []Good, err error) {
	//select * from Good
	db.MysqlDB.Where("organ_id=? AND approve=?", organ_id, code).Find(&goods)
	return
}

func (g *Good) GetGoodsByName(goodName string) (goods []Good, err error) {
	//select * from Good
	db.MysqlDB.Where("good_name LIKE ?", goodName).Find(&goods)
	return
}

//更新审核状态
func (g *Good) UpdateApprove(id int64, code int64) error {
	var good Good
	if g.Approve != 0 {
		return errors.New("该物资已经过审核")
	}
	db.MysqlDB.Model(&good).Where("id=?", id).Update("approve", code)
	return nil
}

//更新审核状态
func (g *Good) UpdateShareState(id int64, code int64) error {
	var good Good
	if g.Approve != 0 {
		return errors.New("该物资已共享")
	}
	db.MysqlDB.Model(&good).Where("id=?", id).Update("share_state", code)
	return nil
}

//查询借出的物资
func (g *Good) GetGoodsLend(user_id int64) (goods []Good, err error) {
	//select * from Good
	db.MysqlDB.Where("user_id=? AND approve=?", user_id, 1).Find(&goods)
	return
}
