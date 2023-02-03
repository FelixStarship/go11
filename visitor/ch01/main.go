package main

import "fmt"

type Info struct {
	Namespace   string
	Name        string
	OtherThings string
}

type VisitorFunc func(*Info, error) error

type Visitor interface {
	Visit(VisitorFunc) error
}

func (info *Info) Visit(fn VisitorFunc) error {
	return fn(info, nil)
}

type NameVisitor struct {
	visitor Visitor
}

func (v *NameVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("NameVisitor() before call function!")
		err = fn(info, err)
		if err == nil {
			fmt.Printf("====>Name=%s,Namespace=%s\n", info.Name, info.Namespace)
		}
		fmt.Println("NameVisitor() after call function!")
		return err
	})
}

type LogVisitor struct {
	visitor Visitor
}

func (v *LogVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("LogVisitor() before call function")
		err = fn(info, err)
		fmt.Println("LogVisitor() after call function")
		return err
	})
}

func main() {

	info := &Info{}
	var v Visitor = info
	v = &LogVisitor{v}
	v = &NameVisitor{v}

	v.Visit(func(info *Info, err error) error {
		info.Name = "leo"
		info.Namespace = "打gopher"
		info.OtherThings = "We are running as remote team."
		return nil
	})

}
