package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main()  {

	//log.SetFlags(30)
	//go exam(1)
	//
	//go func() {
	//	fmt.Println(exam(2))
	//}()
	//
	//
	//time.Sleep(5*time.Second)

	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
	wg.Add(2)
	go SayGreetings("hi",10)
	go SayGreetings("hello!",10)
	wg.Wait()
}

func exam(i int) int  {
	log.Println("test")
	return i
}

func exam1(greeting string,times int)  {
	for i:=0;i<times;i++{
		log.Println(greeting)
		d:=time.Second*time.Duration(rand.Intn(5))
		log.Println(d)
		time.Sleep(d)
	}
}

func SayGreetings(greeting string,times int)  {
	for i:=0;i<times;i++ {
		log.Println(greeting)
		//d:=time.Second*time.Duration(rand.Intn(5))
		//time.Sleep(d)
	}
	wg.Done()
}