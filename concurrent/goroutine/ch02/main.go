package main

import (
	"fmt"
	_ "net/http"
	"time"
)

func defaultFunc() {
	ch := make(chan struct{})
	select {
	case <-ch:
		fmt.Println("test")
	case <-time.After(time.Second * 3):
		fmt.Println("超时了!!")
	}
}
func main() {
	ch := make(chan struct{})
	<-ch
}
