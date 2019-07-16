package controller

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"healthmonitor/service"
)

func CreateUrl(c *gin.Context) {

	service.CreateUrl(c)

}

func ReadUrl(c *gin.Context) {

	service.ReadUrl(c)

}
func FetchData(c *gin.Context){
	service.FetchData(c)
}


