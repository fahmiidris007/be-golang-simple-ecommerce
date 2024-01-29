package handlers

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	dsn := "root:210120@tcp(127.0.0.1:3306)/s-ecom?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// db.AutoMigrate(&models.TransactionItem{}) // Migrate model yang diperlukan
}
