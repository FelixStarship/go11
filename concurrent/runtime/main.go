package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main()  {

	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.GOMAXPROCS(10))
	fmt.Println(runtime.GOMAXPROCS(100))
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Second*1000000)
		wg.Wait()
	}()
	wg.Wait()
}
