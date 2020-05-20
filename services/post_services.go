package services

import (
	"github.com/AhmedSafwat1/go-crud/datasources/mysql/gorm"
	"github.com/AhmedSafwat1/go-crud/domain/posts"
	"github.com/AhmedSafwat1/go-crud/utils/errors"
	"log"
)


func Create(post posts.Post) (* posts.Post, *errors.RestError){
	db , err := gorm.Connect()
	if err != nil {
		log.Fatalln(err)
		return  nil , errors.NewInteenalServerError("error in database")
	}
	db.Create(&post)
	return &post, nil
}