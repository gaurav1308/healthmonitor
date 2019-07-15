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
	c := cron.New()
	c.AddFunc("*/1 * * * *",service.FetchAllUrl)
	c.Start()
}
func main() {

	routes.Init_routes()
	//doEvery(2*time.Second, helloworld)

}