package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"healthmonitor/model"
	"io/ioutil"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"healthmonitor/resource"
	"net/http"
	"sync"
	"time"
)


///funtion to take data from postman_____________________________________________________________________

var wg sync.WaitGroup
func CreateUrl(c *gin.Context) {

	var urls []model.UrlModel
	c.Bind(&urls)


	Tests(urls)


}


///funtion to take data from files


func ReadUrl(c *gin.Context) {

	var urls []model.UrlModel
	p, _ := ioutil.ReadFile(c.Query("path"))
	json.Unmarshal(p, &urls)

	Tests(urls)

}


func Tests(urls []model.UrlModel){
	wg.Add(len(urls))
	for i := 0; i < len(urls); i++ {
		go Test (urls[i])
	}
	wg.Wait()
}

func Test(urls model.UrlModel){
	var count int
	var u model.UrlModel
	resource.Db.Model(&model.UrlModel{}).Where("url = ?", urls.URL).Count(&count)
	if count == 0 {
		resource.Db.Save(&urls)
	} else {
		resource.Db.Where("url = ?", urls.URL).First(&u)
		u.Frequency = urls.Frequency
		u.Crawl_timeout = urls.Crawl_timeout
		u.Failure_thresold = urls.Failure_thresold
		resource.Db.Save(&u)
	}
	wg.Done()
}
///funtion to check health of url in a specific tries________________________________

func FetchData(c *gin.Context){
	id:=c.Param("id")
	tries:=c.Param("tries")
	var data model.UrlData
	resource.Db.Model(&model.UrlModel{}).Where("url_id = ?", id).Where("total_attempts =? ",tries).First(&data)
	if(data.URLID!=0) {
		c.JSON(http.StatusOK,gin.H{"Health":data.Health})
	}else {
		c.JSON(http.StatusOK,gin.H{"Message":"Not exist"})
	}
}


///funtion to fetch all urls and call them after some refreshing time_________________________________________


func FetchAllUrl() {
	var urls []model.UrlModel
	resource.Db.Find(&urls)
	for _, item := range urls {
		go testing(item)


	}
}


///funtion to check status of a url____________________________________________________________________________________


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

			} else {
				udata.Health = "BAD"
				resource.Db.Save(&udata)
				time.Sleep(time.Duration(item.Frequency)*time.Second)

			}
		}

	}
}