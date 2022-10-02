package main

import (
	"fmt"
	"time"
)

func main() {
	//exam()
	//fmt.Println(Triple(5))
	//exam1()
	//exam2()
	//exam3()
	//exam4()
	exam3()

}

func exam() {
	defer fmt.Println("9")
	fmt.Println("0")
	defer fmt.Println("8")
	fmt.Println("1")
	if false {
		defer fmt.Println("not reachable")
	}
	defer func() {
		defer fmt.Println("7")
		fmt.Println("3")
		defer func() {
			fmt.Println("5")
			fmt.Println("6")
		}()
		fmt.Println("4")
	}()
	fmt.Println("2")
	return
	defer fmt.Println("not reachable")
}

func Triple(n int) (r int) {
	defer func() {
		r += n
	}()
	return n + n
}

func exam1() {
	func() {
		for i := 0; i < 3; i++ {
			defer fmt.Println("a:", i)
		}
	}()

	fmt.Println()
	func() {
		for i := 0; i < 3; i++ {
			defer func() {
				fmt.Println("b:", i)
			}()
		}
	}()

}

func exam2() {
	func() {
		for i := 0; i < 3; i++ {
			defer fmt.Println("a:", i)
		}
	}()

	fmt.Println()
	func() {
		for i := 0; i < 3; i++ {
			defer func(i int) {
				fmt.Println("b:", i)
			}(i)
		}
	}()
}

func exam3() {
	defer func() {
		fmt.Println("正常退出")
		panic("恐慌!")
	}()
	fmt.Println("嗨")
	defer func() {
		v := recover()
		fmt.Println("恐慌被恢复了：", v)
	}()
	panic("拜拜!")
	fmt.Println("执行不到这里")
}

func exam4() {
	fmt.Println("hi")
	go func() {
		time.Sleep(time.Second)
		panic(123)
	}()

	for {
		time.Sleep(time.Second)
	}
}

// 匿名函数的值何时被初始化
//【一个匿名函数体内的表达式是在此函数被执行的时候才会被逐渐估值的，不管此函数是被普通调 用还是延迟/协程调用】
func exam5() {
	for i := 0; i < 5; i++ {
		a := i
		go func(i int) {
			fmt.Println("i=", i)
		}(a)
	}
}
