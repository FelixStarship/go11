package main

import "fmt"

func main() {
	var c chan struct{} //nil
	select {
	case <-c: //阻塞
	case c <- struct{}{}: //阻塞
	default:
		fmt.Println("Go here.")
	}
}
