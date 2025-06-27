package utils

import (
	"log"
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	db := os.Getenv("DB_SQLITE")
	if db == "" {
		db = "bewell-test" // ถ้าไม่มีค่ากำหนดไว้ใน env ก็ใช้ default
	}
	DB, err = gorm.Open(sqlite.Open("../sqlite/"+db+".db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("❎ failed to connect database: %v", err)
	}
}
