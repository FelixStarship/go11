package main

import (
	"fmt"
)

func main() {

	for i:=90;i<100;i++ {
		n:=exam2(i)
		fmt.Print("最小的比",i,"最大的素数为",n)
		fmt.Println()
	}
}

func exam3()  {
	i := 0
Next:
	fmt.Println(i)
	i++
	if i < 5 {
		goto Next
	}
}

func exam() {

	i := 0
Next:
	if i > 5 {
		goto Exit
	}

	{
		k := i + 1
		fmt.Println(k)
	}

	i++
	goto Next

Exit:
	fmt.Println("exit")
}

func exam1()  {
	var k int
	i := 0
Next:
	if i > 5 {
		goto Exit
	}


	k = i + 1
	fmt.Println(k)


	i++
	goto Next

Exit:
	fmt.Println("exit")
}

func exam2(n int) int  {
Outer:
	for ;;n++{
		for i:=2;;i++ {
			switch  {
			case i*i>n:
				break Outer
			case n%i==0:
				continue Outer
			}
		}
	}
	return n
}