package main

import (
	"healthmonitor/resource"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/robfig/cron"
	"healthmonitor/routes"
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

}