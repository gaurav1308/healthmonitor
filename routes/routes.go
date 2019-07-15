package routes
import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

import "healthmonitor/controller"

func Init_routes(){

	router := gin.Default()

	v1 := router.Group("/healthmonitor")
	{
	v1.POST("/send", controller.CreateUrl)
	v1.GET("/health/:id/:tries", controller.FetchData)
	}
	router.Run()
}