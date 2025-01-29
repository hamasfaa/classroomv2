package config

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := "root:pass@tcp(127.0.0.1:3306)/classroom?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if db != nil {
		fmt.Println("Connected to database")
	}

	if err != nil {
		fmt.Println("Failed to connect to database")
	}

	return db, err
}
