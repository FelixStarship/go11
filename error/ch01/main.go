package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
)

type Point struct {
	Longitude     string
	Latitude      string
	Distance      string
	ElevationGain string
}

func parse(r io.Reader) (*Point, error) {
	var p Point
	if err := binary.Read(r, binary.BigEndian, &p.Longitude); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.Latitude); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.Distance); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.ElevationGain); err != nil {
		return nil, err
	}
	return &p, nil
}

func parse1(r io.Reader) (*Point, error) {
	var (
		p   Point
		err error
	)
	read := func(data interface{}) {
		if err != nil {
			return
		}
		err = binary.Read(r, binary.BigEndian, data)
	}
	read(&p.Longitude)
	read(&p.Latitude)
	read(&p.Distance)
	read(&p.ElevationGain)

	if err != nil {
		return &p, err
	}
	return &p, err
}

type Person struct {
	Name   [10]byte
	Age    uint8
	Weight uint8
	err    error
}

func (p *Person) read(data interface{}) {
	if p.err == nil {
		p.err = errors.Wrap(binary.Read(bytes.NewReader([]byte("hxz")), binary.BigEndian, data), "read failed")
	}
}

func (p *Person) ReadName() *Person {
	p.read(&p.Name)
	return p
}

func (p *Person) ReadAge() *Person {
	p.read(&p.Age)
	return p
}

func (p *Person) ReadWeight() *Person {
	p.read(&p.Weight)
	return p
}

func (p *Person) Print() *Person {
	if p.err == nil {
		fmt.Printf("Name=%s, Age=%d, Weight=%d\n", p.Name, p.Age, p.Weight)
	}
	return p
}

func main() {
	p := Person{}
	p.ReadName().
		ReadAge().
		ReadWeight().
		Print()
	log.Fatalf("err:%+v", p.err)
}
