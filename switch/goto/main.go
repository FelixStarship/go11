package main

import "fmt"

func main() {
	i := 0
Next:
	fmt.Println(i)
	i++
	if i < 5 {
		goto Next
	}
}
