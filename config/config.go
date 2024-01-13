package config

import (
	"fmt"

	"example.com/ginexample/helpers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	host     = "127.0.0.1"
	port     = 3306
	user     = "root"
	password = "amri"
	dbName   = "db_hris"
)

func DatabaseConnection() *gorm.DB {
	// db, err = sql.Open("mysql", "root:amri@tcp(127.0.0.1:3306)/db_hris")
	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(sqlInfo), &gorm.Config{})
	helpers.ErrorPanic(err)

	return db
}
