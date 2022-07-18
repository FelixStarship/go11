package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	switch n := rand.Intn(100) % 5; n {
	case 0, 1, 2, 3, 4, 5:
		fmt.Println("n=", n)
		fallthrough
	case 6, 7, 8:
		n := 99
		fmt.Println("n=", n)
	}

	switch n := rand.Intn(3); n {
	case 0:
		fmt.Println("n == 0")
	default:
		fmt.Println("n == 2")
	case 1:
		fmt.Println("n == 1")
	}
}
