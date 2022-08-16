package driver

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	engine *xorm.Engine
	client *redis.Client
)

func init() {
	initDB()
	initRedis()
}

func initDB() {
	var err error
	user := "root"
	passwd := "Server.Sues.112"
	host := "mysql"
	port := 3306
	db := "task"
	masterDSN := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4", user, passwd, host, port, db)
	if engine, err = xorm.NewEngine("mysql", masterDSN); err != nil {
		fmt.Println(err)
		panic(err)
	}
	if err = engine.Ping(); err != nil {
		fmt.Println(err)
		panic(err)
	}
	return
}

func initRedis() {
	var (
		res string
		err error
	)
	redisHost := "redis"
	redisPort := "6379"
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: "Server.Sues.112",
		DB:       0,
	})
	if res, err = client.Ping(context.Background()).Result(); err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(res)
	return
}
