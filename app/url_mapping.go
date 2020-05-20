package app

import (
	"fmt"
	"github.com/AhmedSafwat1/go-crud/controllers"
	"github.com/AhmedSafwat1/go-crud/datasources/mysql/gorm"
)

func mapUrl()  {
	db  := gorm.DB
	fmt.Println(db)
	//defer db.Close()

	router.GET("/ping", controllers.Ping)

	router.GET("/posts", controllers.Index)

	router.GET("/posts/:id", controllers.Show)

	router.POST("posts", controllers.Store)

	router.PATCH("/posts/:id", controllers.Edit)

	router.DELETE("/posts/:id", controllers.Delete)




}
