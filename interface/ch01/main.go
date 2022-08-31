package main

import "fmt"

type Country struct {
	Name string
}

type City struct {
	Name string
}

type Stringable interface {
	ToString() string
}

func (c City) ToString() string {
	return "City=" + c.Name
}

func (c Country) ToString() string {
	return "Country=" + c.Name
}

func PrintStr(p ...Stringable) {
	for _, s := range p {
		fmt.Println(s.ToString())
	}
}

func main() {
	PrintStr(Country{"USA"}, City{"西雅图"})
}
