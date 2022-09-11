package main

import (
	"fmt"
	"time"
)

func main() {
	for i, b := range []byte("明月") {
		fmt.Println(i, b)
	}
	hello := []byte("Hello")
	world := "World!"
	helloWorld := append(hello, world...)
	fmt.Println(string(helloWorld))

	t := time.Now()
	fmt.Println(time.Now().Sub(t))

}
