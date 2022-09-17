package main

import "fmt"

func main() {
	ch02()
}

func ch02() {
	c := make(chan struct{})
	close(c)
	select {
	case c <- struct{}{}:
	case <-c:

	}
}

func ch01() {
	c := make(chan string, 2)
	trySend := func(v string) {
		select {
		case c <- v:
		default:

		}
	}
	tryReceive := func() string {
		select {
		case v := <-c:
			return v
		default:
			return "-"
		}
	}
	trySend("Hello!")
	trySend("Hi!")
	trySend("Bye!")
	fmt.Println(tryReceive())
	fmt.Println(tryReceive())
	fmt.Println(tryReceive())
}
