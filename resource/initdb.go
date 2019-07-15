package resource

import (
	"github.com/jinzhu/gorm"
	"healthmonitor/model"
)
var Db *gorm.DB

func IniDb(){
	var err error
	Db, err = gorm.Open("mysql", "root:gaurav@/healthmonitor?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
	Db.AutoMigrate(&model.UrlModel{})
	Db.AutoMigrate(&model.UrlData{})
}