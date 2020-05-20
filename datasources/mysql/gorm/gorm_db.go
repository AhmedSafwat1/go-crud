package gorm

import (
	"fmt"
	"github.com/AhmedSafwat1/go-crud/domain/posts"
	"github.com/AhmedSafwat1/go-crud/utils/errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)
const (
	mysql_user_name     = "mysql_user_name"
	mysql_user_password ="mysql_user_password"
	mysql_user_host		= "mysql_user_host"
	mysql_user_shema	= "mysql_user_schema"
)

var DB *gorm.DB


func init() {
	var err * errors.RestError
	DB , err := Connect()
	if err != nil {
		log.Fatalln(err)
	}
	DB.AutoMigrate(&posts.Post{})
	fmt.Println(DB, "in init");
}


func Connect() (*gorm.DB, *errors.RestError) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"root",
		"",
		"127.0.0.1:3306",
		"users_db",
	)
	db, err := gorm.Open("mysql", dataSourceName)

	if err != nil {
		return nil, errors.NewInteenalServerError("database error")
	}

	log.Println("database connect hi ..... ")
	return db, nil;
}