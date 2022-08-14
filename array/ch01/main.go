package main

import "fmt"

type st struct {
	Name string
	Agt  []int
}

func main() {
	array := [...]bool{true, false}
	fmt.Println(array)

	array1 := [1]int{}
	fmt.Println(array1)
	st1 := st{}
	fmt.Println(st1)

	pm := &map[string]int{"C": 1}
	fmt.Printf("%T\n", pm)

	st2 := st{Name: "felix", Agt: []int{1, 2, 3}}
	fmt.Println(st2)
	st2.oprtation()
	fmt.Println(st2)
}

func (s *st) oprtation() {
	s.Name = "指针"
	s.Agt = []int{4, 5, 6}
}
