package main

import "net/http"

// Suggestions from golang-nuts
// http://play.golang.org/p/Ctg3_AQisl

//func doEvery(d time.Duration, f func(time.Time)) {
//	for x := range time.Tick(d) {
//		f(x)
//	}
//}
//
//func helloworld(t time.Time) {
//	fmt.Printf("%v: Hello, World!\n", t)
//}

func main() {
	//doEvery(2*time.Second, helloworld)
	resp, err := http.Get("http://google.com/")
	if err != nil {
		print("k"+err.Error())
	} else {
		print("h"+string(resp.StatusCode) + resp.Status)
	}
}