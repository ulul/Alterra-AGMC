package config

import (
	"fmt"
	"testing/models"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var JWTSecret = []byte("t151sS3cr3t")
var JWTExp = time.Now().Add(time.Hour * 24).Unix()

func InitDB() {
	config := map[string]string{
		"DB_USERNAME": "local",
		"DB_PASSWORD": "s3cr3t",
		"DB_PORT":     "3306",
		"DB_HOST":     "localhost",
		"DB_NAME":     "agmc",
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_USERNAME"],
		config["DB_PASSWORD"],
		config["DB_HOST"],
		config["DB_PORT"],
		config["DB_NAME"])

	var e error
	DB, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
}
