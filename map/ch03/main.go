package main

import "fmt"

func main() {
	type Ta []int
	type Tb []int
	dest := Ta{1, 2, 3}
	src := Tb{5, 6, 7, 8, 9}
	n := copy(dest, src)
	fmt.Println(n, src, dest)
	n = copy(dest[1:], dest)
	fmt.Println(n, dest)

	//遍历一个nil映射和nil切片是允许的
	var s []int
	for _, v := range s {
		fmt.Println(v)
	}

	//遍历元素
	type Person struct {
		name string
		age  int
	}

	person := [2]Person{{"Alice", 28}, {"Bob", 25}}
	for i, p := range person {
		fmt.Println(i, p)
		person[1].name = "Jack"
		p.age = 31
	}
	fmt.Println(person)

	person1 := []Person{{"Alice", 28}, {"Bob", 25}}
	for i, p := range person1 {
		fmt.Println(i, p)
		person1[1].name = "Jack"
		p.age = 31
	}
	fmt.Println(person1)

	langs := map[struct{ dynamic, strong bool }]map[string]int{
		{true, false}:  {"JavaScrips": 1995},
		{false, true}:  {"Go": 2009},
		{false, false}: {"C": 1972},
	}

	m0 := map[*struct{ dynamic, strong bool }]*map[string]int{}
	for category, langInfo := range langs {
		m0[&category] = &langInfo
		category.dynamic, category.strong = true, true
	}

	for category, langInfo := range langs {
		fmt.Println(category, langInfo)
	}

	for category, langInfo := range m0 {
		fmt.Println(*category, *langInfo)
	}

}
