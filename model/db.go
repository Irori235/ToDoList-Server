package model

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "gorm.io/driver/mysql"
)

var db *gorm.DB

func DBConnection() (*gorm.DB, error) {
	_db, err := gorm.Open(GetDBConfig())
	if err != nil {
		return nil, err
	}
	db = _db
	return db, err
}

func GetDBConfig() (string, string) {
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	hostname := os.Getenv("DB_HOSTNAME")
	dbname := os.Getenv("DB_DBNAME")

	connect := user + ":" + password + "@" + "tcp(" + hostname + ":3306)" + "/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(connect)
	return "mysql", connect
}

func CreateDB() *gorm.DB {
	db, err := DBConnection()
	if err != nil {
		panic(fmt.Errorf("DB Error: %w", err))
	}
	CreateTable(db)
	return db
}

func CreateTable(db *gorm.DB) {
	db.AutoMigrate(&Task{})
}
