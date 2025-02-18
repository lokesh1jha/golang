package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	db *gorm.DB
)

func Connect() {
	database, err := gorm.Open("mysql", "root:password@/bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect to database!")
	}
	db = database
}

func GetDB() *gorm.DB {
	return db
}
