package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"time"
)

var db *gorm.DB

type Websites struct{
	Websites []Website`json:"websites"`
}


type  Website struct{
	URL     string `json:"url"`
	Crawl_timeout int    `json:"crawl_timeout"`
	Frequency	int `json:"frequency"`
	Failure_thresold int`json:"failure_thresold"'`
}
func init() {
	//open a db connection
	var err error
	db, err = gorm.Open("mysql", "root:gaurav@/healthmonitor?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
	db.AutoMigrate(&urlModel{})
	db.AutoMigrate(&urlData{})
}


func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}


//func helloworld(t time.Time) {
//	resp, err := http.Get("http://google.com/")
//
//}


func main() {

	router := gin.Default()




	v1 := router.Group("/healthmonitor")
	{
		v1.POST("/send", createUrl)
		v1.GET("/health", fetchAllUrl)
	}
	router.Run()
	//doEvery(2*time.Second, helloworld)

}

type (
	// todoModel describes a todoModel type
	urlModel struct {
		gorm.Model
		URL     string `json:"url"`
		Crawl_timeout int    `json:"crawl_timeout"`
		Frequency	int `json:"frequency"`
		Failure_thresold int`json:"failure_thresold"'`
	}

	// transformedTodo represents a formatted todo
	transformedModel struct {
		ID        uint   `json:"id"`
		URL     string `json:"url"`
		Crawl_timeout int    `json:"crawl_timeout"`
		Frequency	int `json:"frequency"`
		Failure_thresold int`json:"failure_thresold"'`
	}

	urlData struct {
		gorm.Model
		RID       uint `db:"rid"`
		Attempts    int    `db:"attempts"`
		Health int    `db:"health"`
	}

)

func createUrl(c *gin.Context) {

	var urls []urlModel
	c.Bind(&urls)
	for i := 0; i < len(urls); i++ {
		var count int
		var u urlModel
		db.Model(&urlModel{}).Where("url = ?", urls[i].URL).Count(&count)
		if count == 0 {
			//fmt.Println("count == 0")
			db.Save(&urls[i])
		} else {
			//fmt.Println("count==1")
			db.Where("url = ?", urls[i].URL).First(&u)
			u.Frequency=urls[i].Frequency
			u.Crawl_timeout=urls[i].Crawl_timeout
			u.Failure_thresold=urls[i].Failure_thresold
			db.Save(&u)
		}

		}

}


func fetchAllUrl(c *gin.Context) {
	var urls []urlModel
	//fmt.Println("hi1")
	db.Find(&urls)

	if len(urls) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No URL found!"})
		return
	}

	for _, item := range urls {
		//fmt.Println("hi2")
		go testing(item)
			//for i:=0;i<item.Failure_thresold;i++ {
			//	//fmt.Println("hi3")
			//	var udata urlData
			//	udata.RID=item.ID
			//	udata.Attempts=i+1
			//
			//	timeout := time.Duration(item.Crawl_timeout) * time.Millisecond
			//	client := http.Client{
			//		Timeout: timeout,
			//	}
			//	resp, err := client.Get(item.URL)
			//	if err != nil {
			//		fmt.Println(err.Error())
			//		udata.Health = 0
			//		db.Save(&udata)
			//		time.Sleep(time.Duration(item.Frequency)*time.Second)
			//	} else {
			//		if resp.StatusCode == 200 {
			//			udata.Health = 1
			//			db.Save(&udata)
			//			//fmt.Println("hi4")
			//			return
			//			break
			//
			//		} else {
			//			udata.Health = 0
			//			db.Save(&udata)
			//			time.Sleep(time.Duration(item.Frequency)*time.Second)
			//
			//		}
			//	}
			//
			//}


	}
}

func testing(item urlModel){
	for i:=0;i<item.Failure_thresold;i++ {
		var udata urlData
		udata.RID=item.ID
		udata.Attempts=i+1

		timeout := time.Duration(item.Crawl_timeout) * time.Millisecond
		client := http.Client{
			Timeout: timeout,
		}
		resp, err := client.Get(item.URL)
		if err != nil {
			fmt.Println(err.Error())
			udata.Health = 0
			db.Save(&udata)
			time.Sleep(time.Duration(item.Frequency)*time.Second)
		} else {
			if resp.StatusCode == 200 {
				udata.Health = 1
				db.Save(&udata)
				return
				break

			} else {
				udata.Health = 0
				db.Save(&udata)
				time.Sleep(time.Duration(item.Frequency)*time.Second)

			}
		}

	}
}