package main

import (
	"fmt"
	"time"
)

func main() {
	ch02()
}

func ch01() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	c <- 3 //产生恐慌
	close(c)
	c <- 4
}

func ch02() {
	var ball = make(chan string)
	kickBall := func(playerName string) {
		for {
			fmt.Println(<-ball, "传球")
			time.Sleep(time.Second)
			ball <- playerName
		}
	}
	go kickBall("张三")
	go kickBall("李四")
	go kickBall("王二麻子")
	go kickBall("刘大")
	ball <- "裁判"
	var c chan bool //一个零值nil通道、通道阻塞
	<-c
}
