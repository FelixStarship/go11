package main

import "fmt"

type Inter1 interface {
	fun1()
	fun2()
}

type Worker struct {
	Name string
	Age  int
	// 内嵌接口类型
	Inter1
}

func (w Worker) fun1() {
	fmt.Println("Worker fun1")
}

type Salary struct {
	Money int
}

func (s Salary) fun1() {
	fmt.Println("Salary fun1")
}
func (s Salary) fun2() {
	fmt.Println("Salary fun2")
}

func main() {
	s := Salary{}
	w := Worker{Inter1: s}

	//w.Age
	//w.Name
	w.fun1()
	w.fun2()
	//w.Inter1
	w.Inter1.fun1()
	w.Inter1.fun2()
	// 断言访问
	fmt.Println(w.Inter1.(Salary).Money)
}
