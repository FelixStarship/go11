# golang中结构体嵌套接口
在golang中结构体A嵌套另一个结构体B见的很多，可以扩展A的能力。

A不仅拥有了B的属性，还拥有了B的方法，这里面还有一个字段提升的概念。

```go
package main

import "fmt"

type Worker struct {
	Name string
	Age int
	// 结构体类型嵌套
	Salary
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
	w := Worker{Salary: s}

	//w.Name
	//w.Age
	//w.Money
	//w.Salary
	//w.fun1()
	//w.fun2()
	//w.Salary.fun1()
	//w.Salary.fun2()
}
```

***很明显现在 Worker 强依赖与 Salary ，有时候我们希望 Worker 只依赖于一个接口，这样只要实现了此接口的对象都可以传递进来。***

优化之后：

```go
package main

import "fmt"

type Inter1 interface {
	fun1()
	fun2()
}

type Worker struct {
	Name string
	Age int
	// 内嵌接口类型：对外多态实现，只要实现了Inter1接口的对象都可以【赋值】
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
	//w.fun1()
	//w.fun2() 隐式调用Inter1对应的成员方法
	//w.Inter1
	//w.Inter1.fun1()
	//w.Inter1.fun2()
	// 无法访问 Money 属性，可以增加方法来实现
}
```

