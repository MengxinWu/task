package driver

import (
	"context"
	"fmt"
	"task/models"

	"github.com/bwmarrin/snowflake"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

var (
	engine       *xorm.Engine
	client       *redis.Client
	node         *snowflake.Node
	dispatchConn *kafka.Conn
	executeConn  *kafka.Conn
)

// InitEngine init engine.
func InitEngine() {
	var err error
	user := "root"
	passwd := "Server.Sues.112"
	host := "mysql"
	port := 3306
	db := "task"
	masterDSN := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4", user, passwd, host, port, db)
	if engine, err = xorm.NewEngine("mysql", masterDSN); err != nil {
		log.Panic(err)
	}
	if err = engine.Ping(); err != nil {
		log.Panic(err)
	}
	log.Printf("init engine success...")
	return
}

// InitRedis init redis.
func InitRedis() {
	var (
		ctx    = context.Background()
		result string
		err    error
	)
	redisHost := "redis"
	redisPort := "6379"
	passwd := "Server.Sues.112"
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: passwd,
		DB:       0,
	})
	if result, err = client.Ping(ctx).Result(); err != nil {
		log.Panic(err)
	}
	log.Printf("redis ping result: %s", result)
	log.Printf("init redis success...")
	return
}

// InitSnowNode init snow node.
func InitSnowNode() {
	var err error
	if node, err = snowflake.NewNode(0); err != nil {
		log.Panic(err)
	}
	log.Printf("init snow node success...")
	return
}

// GetSnowFlakeId get snow id.
func GetSnowFlakeId() int64 {
	return node.Generate().Int64()
}

// InitDispatchConn init dispatch conn.
func InitDispatchConn() {
	var (
		ctx = context.Background()
		err error
	)
	if dispatchConn, err = kafka.DialLeader(ctx, "tcp", models.KafkaAddress, models.KafkaTopicDispatch, 0); err != nil {
		log.Panic(err)
	}
	log.Printf("init dispatch conn success...")
	return
}

// InitExecuteConn init execute conn.
func InitExecuteConn() {
	var (
		ctx = context.Background()
		err error
	)
	if executeConn, err = kafka.DialLeader(ctx, "tcp", models.KafkaAddress, models.KafkaTopicExecute, 0); err != nil {
		log.Panic(err)
	}
	log.Printf("init execute conn success...")
	return
}
