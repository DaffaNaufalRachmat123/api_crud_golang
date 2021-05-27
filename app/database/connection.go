package database

import (
	"fmt"
	"api_crud/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db *gorm.DB
var err error
func init_database(){
	dbType := config.LoadEnv("DB_CONNECTION")
	var connectionString string
	if dbType == "mysql" {
		connectionString = fmt.Sprintf(
			"%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.LoadEnv("DB_USERNAME"), config.LoadEnv("DB_PASSWORD"), config.LoadEnv("DB_HOST"), config.LoadEnv("DB_NAME"))
	} else if dbType == "postgres" {
		connectionString = fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
			config.LoadEnv("DB_HOST"), config.LoadEnv("DB_PORT"), config.LoadEnv("DB_USERNAME"), config.LoadEnv("DB_NAME"), config.LoadEnv("DB_PASSWORD"))
	}
	db , err = gorm.Open(dbType , connectionString)
	if err != nil {
		fmt.Print(err.Error())
	}

}

func GetDatabase() *gorm.DB {
	init_database()
	return db
}

func CloseDatabase(){
	db.Close()
}