package main

import (
	//"fmt"
	//"github.com/gin-gonic/gin"
	"healthmonitor/resource"

	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/robfig/cron"
	"healthmonitor/routes"
	//"net/http"
	//"time"
	//"./controller"
	"healthmonitor/service"
)
func init() {
	//open a db connection

	resource.IniDb()
	refreshTime := cron.New()    //Refresh after 10 minutes
	refreshTime.AddFunc("*/10 * * * *",service.FetchAllUrl)
	refreshTime.Start()
}
func main() {

	routes.Init_routes()
	//doEvery(2*time.Second, helloworld)

}