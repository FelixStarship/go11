package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.RWMutex
	go func() {
		m.RLock()
		defer m.RUnlock()
		fmt.Println("a")
		time.Sleep(time.Second)
	}()

	go func() {
		time.Sleep(time.Second * 1 / 4)
		m.Lock()
		defer m.Unlock()
		fmt.Println("b")
		time.Sleep(time.Second)
	}()

	go func() {
		time.Sleep(time.Second * 2 / 4)
		m.Lock()
		defer m.Unlock()
		fmt.Println("c")
	}()

	go func() {
		time.Sleep(time.Second * 3 / 4)
		m.RLock()
		defer m.RUnlock()
		fmt.Println("d")
	}()

	time.Sleep(time.Second * 3)
	fmt.Println()
}
