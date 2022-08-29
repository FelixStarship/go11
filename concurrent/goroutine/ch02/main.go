package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

func defaultFunc(w http.ResponseWriter, r *http.Request) {
	ch := make(chan struct{})
	select {
	case <-ch:
		fmt.Println("test")
	case <-time.After(time.Second * 3):
		fmt.Println("超时了!!")
	}
	fmt.Fprintf(w, "%v\n", "welcome to user")
}

func cancelFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println(runtime.NumGoroutine())
}

func main() {
	ch := make(chan struct{}, 1)
	<-ch
	//http.HandleFunc("/", defaultFunc)
	//http.HandleFunc("/cancel", cancelFunc)
	//if err := http.ListenAndServe(":8081", nil); err != nil {
	//	log.Fatal("ListenAndServer: ", err)
	//}
}
