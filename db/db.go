package db

import (
	"fmt"
	"runtime/debug"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	MysqlDB *gorm.DB
	err     error
)

type Model struct {
	ID         int       `gorm:"primary_key" json:"id"`
	CreatedOn  time.Time `json:"created_on"`
	ModifiedOn time.Time `json:"modified_on"`
}

func init() {
	//防止程序崩溃
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(fmt.Errorf("InitMysql:%+v", err))
			fmt.Println(fmt.Errorf("%s", string(debug.Stack())))
		}
	}()

	MysqlDB, err = gorm.Open("mysql", "root:xdwd@(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("init db ok!")
	MysqlDB.SingularTable(true)

}

func CloseDB() {
	defer MysqlDB.Close()
}
