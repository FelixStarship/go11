package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Visitor func(shape Shape)

type Shape interface {
	accept(Visitor)
}

type Circle struct {
	Radius int
}

func (c *Circle) accept(v Visitor) {
	v(c)
}

type Rectangle struct {
	Width, Height int
}

func (r *Rectangle) accept(v Visitor) {
	v(r)
}

func JsonVisitor(shape Shape) {
	bytes, err := json.Marshal(shape)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func XmlVisitor(shape Shape) {
	bytes, err := xml.Marshal(shape)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func main() {

	shapes := []Shape{&Circle{10}, &Rectangle{100, 200}}

	for _, s := range shapes {
		s.accept(JsonVisitor)
		s.accept(XmlVisitor)
	}
}
