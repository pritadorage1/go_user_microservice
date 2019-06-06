package model

import (
	"fmt"
	"log"

	config "go_user_microservice/config"

	"github.com/jinzhu/gorm"

	//we can change this as per your database
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//CreateConnection : create mysql
func CreateConnection(conf config.DBConfig) (*gorm.DB, error) {

	// MYSQL Connection string
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.Username, conf.Password, conf.Host, conf.Port, conf.DbName)
	log.Println("Database Connected")
	return gorm.Open("mysql", connectionString)

}
