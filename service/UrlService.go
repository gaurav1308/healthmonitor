package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"healthmonitor/model"
	"io/ioutil"

	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"healthmonitor/resource"
	//"github.com/robfig/cron"
	"net/http"
	"time"
)

func CreateUrl(c *gin.Context) {

	var urls []model.UrlModel
	c.Bind(&urls)
	for i := 0; i < len(urls); i++ {
		var count int
		var u model.UrlModel
		resource.Db.Model(&model.UrlModel{}).Where("url = ?", urls[i].URL).Count(&count)
		if count == 0 {
			//fmt.Println("count == 0")
			resource.Db.Save(&urls[i])
		} else {
			//fmt.Println("count==1")
			resource.Db.Where("url = ?", urls[i].URL).First(&u)
			u.Frequency=urls[i].Frequency
			u.Crawl_timeout=urls[i].Crawl_timeout
			u.Failure_thresold=urls[i].Failure_thresold
			resource.Db.Save(&u)
		}

	}


}



func ReadUrl(c *gin.Context) {

	var urls []model.UrlModel
	p, _ := ioutil.ReadFile(c.Query("path"))
	json.Unmarshal(p, &urls)
	for i := 0; i < len(urls); i++ {
		var count int
		var u model.UrlModel
		resource.Db.Model(&model.UrlModel{}).Where("url = ?", urls[i].URL).Count(&count)
		if count == 0 {
			//fmt.Println("count == 0")
			resource.Db.Save(&urls[i])
		} else {
			//fmt.Println("count==1")
			resource.Db.Where("url = ?", urls[i].URL).First(&u)
			u.Frequency = urls[i].Frequency
			u.Crawl_timeout = urls[i].Crawl_timeout
			u.Failure_thresold = urls[i].Failure_thresold
			resource.Db.Save(&u)
		}

	}

}



func FetchData(c *gin.Context){
	id:=c.Param("id")
	tries:=c.Param("tries")
	var data model.UrlData
	//db.Find(&data)
	//var health int
	resource.Db.Model(&model.UrlModel{}).Where("url_id = ?", id).Where("total_attempts =? ",tries).First(&data)
	if(data.URLID!=0) {
		//fmt.Println(data.Health)
		c.JSON(http.StatusOK,gin.H{"Health":data.Health})
	}else {
		c.JSON(http.StatusOK,gin.H{"Message":"Not exist"})
		//fmt.Println("faltu")
	}
}
func FetchAllUrl() {
	var urls []model.UrlModel
	resource.Db.Find(&urls)

	//if len(urls) <= 0 {
	//	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No URL found!"})
	//	return
	//}

	for _, item := range urls {
		go testing(item)


	}
}

func testing(item model.UrlModel){
	for i:=0;i<item.Failure_thresold;i++ {
		var udata model.UrlData
		udata.URLID=item.ID
		udata.Attempts=i+1
		udata.Total_attempts=item.Tries+1
		//var up urlModel
		item.Tries=item.Tries+1;
		resource.Db.Save(&item)
		timeout := time.Duration(item.Crawl_timeout) * time.Millisecond
		client := http.Client{
			Timeout: timeout,
		}
		resp, err := client.Get(item.URL)
		if err != nil {
			fmt.Println(err.Error())
			udata.Health = "BAD"
			resource.Db.Save(&udata)
			time.Sleep(time.Duration(item.Frequency)*time.Second)
		} else {
			if resp.StatusCode == 200 {
				udata.Health = "GOOD"
				resource.Db.Save(&udata)
				return
				break

			} else {
				udata.Health = "BAD"
				resource.Db.Save(&udata)
				time.Sleep(time.Duration(item.Frequency)*time.Second)

			}
		}

	}
}