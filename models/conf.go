package models

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func init() {
	db = GetDB()
	db.AutoMigrate(&User{})
}

func GetDB() *gorm.DB {
	db, err = gorm.Open(sqlite.Open("db.sqlite3"), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
