package main

import (
	"fmt"
	"gin-web/database"
	"gin-web/routers"
)

func main() {
	fmt.Println("version:", Version)
	fmt.Println("branch:", Branch)
	fmt.Println("revision:", Revision)
	fmt.Println("build_date:", BuildDate)

	database.InitMysql()
	router := routers.InitRouter()

	router.Static("/static", "./static")
	router.Run("0.0.0.0:8081")
}
