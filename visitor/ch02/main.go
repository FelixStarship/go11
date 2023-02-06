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

type DecoratedVisitor struct {
	visitor    Visitor
	decorators []VisitorFunc
}

func NewDecoratedVisitor(v Visitor, fn ...VisitorFunc) Visitor {
	if len(fn) == 0 {
		return v
	}
	return &DecoratedVisitor{visitor: v, decorators: fn}
}

func (v *DecoratedVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		if err != nil {
			return err
		}
		if err := fn(info, err); err != nil {
			return err
		}
		for i := range v.decorators {
			if err := v.decorators[i](info, nil); err != nil {
				return err
			}
		}
		return nil
	})
}

func main() {
	//var v Visitor = new(Info)
	//v = NewDecoratedVisitor(v, NameVisitor, LogVisitor)
}
