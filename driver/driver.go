package driver

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func init() {
	var err error
	user := "root"
	passwd := "Server.Sues.112"
	host := "127.0.0.1"
	port := 43306
	db := "task"
	masterDSN := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4", user, passwd, host, port, db)
	if engine, err = xorm.NewEngine("mysql", masterDSN); err != nil {
		//panic(err)
		log.Printf("error: %v\n", err)
	}
	if err = engine.Ping(); err != nil {
		//panic(err)
		log.Printf("error: %v\n", err)
	}
}
