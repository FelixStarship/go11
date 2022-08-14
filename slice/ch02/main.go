package main

import (
	"fmt"
)

func main() {
	m := map[string]int{"Go": 2007}
	m["C"] = 1972
	m["Java"] = 1995
	fmt.Println(m)
	m["Go"] = 2009
	delete(m, "Java")
	fmt.Println(m)
}
