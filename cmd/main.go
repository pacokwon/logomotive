package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pacokwon/logomotive/cmd/routers"
	// "github.com/pacokwon/logomotive/cmd/store"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=127.0.0.1 user=root password=kbcat dbname=logomotive port=5432 TimeZone=Asia/Seoul"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// env := store.Env{DB: db}
	// if err != nil {
	// 	return
	// }

	fmt.Println("Connected To Database")

	router := gin.Default()
	routers.Init(router)

	router.Run(":8080")
}
