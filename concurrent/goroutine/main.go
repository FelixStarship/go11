package main

import (
	"fmt"
	"log"
	"time"
)

func main()  {

	log.SetFlags(30)
	go exam(1)

	go func() {
		fmt.Println(exam(2))
	}()


	time.Sleep(5*time.Second)
}

func exam(i int) int  {
	log.Println("test")
	return i
}