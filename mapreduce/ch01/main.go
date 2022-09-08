package main

import "fmt"

func main() {
	old := EmployeeCountIf(list, func(e *Employee) bool {
		return e.Age > 10
	})
	fmt.Println(old)
}

type Employee struct {
	Name     string
	Age      int
	Vacation int
	Salary   int
}

var list = []Employee{
	{"扯淡1号", 10, 0, 100},
	{"扯淡2号", 10, 1, 100},
	{"扯淡3号", 10, 2, 100},
	{"扯淡4号", 10, 3, 100},
	{"扯淡5号", 10, 4, 100},
}

func EmployeeCountIf(list []Employee, fn func(e *Employee) bool) int {
	count := 0
	for i := range list {
		if fn(&list[i]) {
			count += 1
		}
	}
	return count
}
