package main

import (
	"fmt"
	"github.com/robfig/cron"

	//"net/http"
	"time"
)

// Suggestions from golang-nuts
// http://play.golang.org/p/Ctg3_AQisl

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func helloworld() {
	fmt.Printf(" Hello, World!\n",)
}

func main() {
	//doEvery(2*time.Second, helloworld)
	//resp, err := http.Get("http://google.com/")
	//if err != nil {
	//	print("k"+err.Error())
	//} else {
	//	print("h"+string(resp.StatusCode) + resp.Status)
	//}
	c := cron.New()
	c.AddFunc("*/1 * * * *",helloworld)
	c.Start()
}