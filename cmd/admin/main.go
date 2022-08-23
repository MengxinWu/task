package main

import (
	"task/driver"
	"task/service/admin"
)

func main() {
	// 初始化资源
	driver.InitEngine()
	driver.InitSnowNode()
	driver.InitDispatchWriter()

	// 创建并运行gin服务
	router := service.NewRouter()
	_ = router.Run(":8080")
}
