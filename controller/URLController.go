package controller

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	//"../model"
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"healthmonitor/resource"
	//"github.com/robfig/cron"
	//"net/http"
	//"time"
	"healthmonitor/service"
)

func CreateUrl(c *gin.Context) {

	service.CreateUrl(c)

}

func FetchData(c *gin.Context){
	service.FetchData(c)
}


