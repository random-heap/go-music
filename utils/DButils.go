package utils

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)



//定义数据库对象
var DB *sqlx.DB
var err error

func init() {
	fmt.Println("初始化数据库")

	//定义mysql数据源，配置数据库地址，帐号以及密码
	dsn := "root:root1243@tcp(localhost:3306)/go_music?charset=utf8&parseTime=True&loc=Local"

	//根据数据源dsn和mysql驱动, 创建数据库对象
	DB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

}
