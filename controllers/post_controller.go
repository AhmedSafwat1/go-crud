package controllers

import (
	"github.com/AhmedSafwat1/go-crud/domain/posts"
	"github.com/AhmedSafwat1/go-crud/services"
	"github.com/AhmedSafwat1/go-crud/utils/errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Index(c *gin.Context){
	c.String(http.StatusNotImplemented, "not implemnt")
}

func Store(c *gin.Context){
	var (
		post  posts.Post

	)
	errBind := c.ShouldBindJSON(&post)
	if errBind != nil {
		err := errors.ErrorValidation(errBind.Error())
		log.Println(errBind)
		c.JSON(err.Code, err)
		return
	}
	result , err := services.Create(post)
	if err != nil {
		c.JSON(err.Code, err);
		return
	}
	c.JSON(http.StatusCreated, result)
}

func Edit(c *gin.Context){
	c.String(http.StatusNotImplemented, "not implemnt")
}

func Delete(c *gin.Context){
	c.String(http.StatusNotImplemented, "not implemnt")
}

func Show(c *gin.Context){
	c.String(http.StatusNotImplemented, "not implemnt")
}