package main

import "task/service/admin"

func main() {
	router := service.NewRouter()
	_ = router.Run(":8080")
}
