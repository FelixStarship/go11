package main

import "fmt"

func main()  {
	//exam()
	//fmt.Println(Triple(5))
	//exam1()
	exam2()

}

func exam()  {
	defer fmt.Println("9")
	fmt.Println("0")
	defer fmt.Println("8")
	fmt.Println("1")
	if false{
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
		r+=n
	}()
	return n+n
}

func exam1()  {
	func(){
		for i:=0;i<3;i++ {
			defer fmt.Println("a:",i)
		}
	}()

	fmt.Println()
	func(){
		for i:=0;i<3;i++ {
			defer func(){
				fmt.Println("b:",i)
			}()
		}
	}()

}

func exam2()  {
	func(){
		for i:=0;i<3;i++ {
			defer fmt.Println("a:",i)
		}
	}()

	fmt.Println()
	func(){
		for i:=0;i<3;i++ {
			defer func(i int){
				fmt.Println("b:",i)
			}(i)
		}
	}()
}