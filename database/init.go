package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initdb() *gorm.DB {
	dsn := "root:11220518@tcp(127.0.0.1:3306)/restaurant?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database", err.Error())
	}
	err0 := db.AutoMigrate(&Staff{})
	if err0 != nil {
		fmt.Println("failed to migrate staff", err0.Error())
	}
	err1 := db.AutoMigrate(&Customer{})
	if err1 != nil {
		fmt.Println("failed to migrate customer", err1.Error())
	}
	err2 := db.AutoMigrate(&Table{})
	if err2 != nil {
		fmt.Println("failed to migrate customer", err2.Error())
	}
	err3 := db.AutoMigrate(&Dish{})
	if err3 != nil {
		fmt.Println("failed to migrate customer", err3.Error())
	}
	err4 := db.AutoMigrate(&Order{}, &OrderItem{})
	if err4 != nil {
		fmt.Println("failed to migrate customer", err4.Error())
	}
	DB = db
	return db
}
func Getdb() *gorm.DB {
	return DB
}
