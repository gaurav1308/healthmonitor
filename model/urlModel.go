package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
type (
UrlModel struct {
gorm.Model
URL     string `json:"url"`
Crawl_timeout int    `json:"crawl_timeout"`
Frequency	int `json:"frequency"`
Failure_thresold int`json:"failure_thresold"'`
Tries int  `json:"tries" gorm:"default:0"`
}

//// transformedTodo represents a formatted todo
//transformedModel struct {
//ID        uint   `json:"id"`
//URL     string `json:"url"`
//Crawl_timeout int    `json:"crawl_timeout"`
//Frequency	int `json:"frequency"`
//Failure_thresold int`json:"failure_thresold"'`
//}


)

