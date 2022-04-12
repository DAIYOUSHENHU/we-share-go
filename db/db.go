package db

import (
	"fmt"
	"runtime/debug"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var MysqlDB *gorm.DB

func InitDB() error {
	//防止程序崩溃
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(fmt.Errorf("InitMysql:%+v", err))
			fmt.Println(fmt.Errorf("%s", string(debug.Stack())))
		}
	}()

	db, err := gorm.Open("mysql", "root:xdwd@(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	MysqlDB = db
	fmt.Println("init db ok!")
	return nil
}
