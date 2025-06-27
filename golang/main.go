package main

import (
	product_handler "bewell-test/pkg/product/handlers"
	"bewell-test/utils"

	"fmt"
	"net/http"
	"os"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "2606" // ถ้าไม่มีค่ากำหนดไว้ใน env ก็ใช้ default
	}

	utils.InitDB()

	if utils.DB == nil {
		fmt.Println("DB is nil")
		return
	}
	utils.DB = utils.DB.Session(&gorm.Session{
		Logger: utils.DB.Logger.LogMode(logger.Info),
	})
	fmt.Println("✅ DB is connected")

	product_handler.Call()

	fmt.Println("Server started at :" + port)
	http.ListenAndServe(":"+port, nil)
}
