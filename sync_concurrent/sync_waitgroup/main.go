package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	const N = 5
	var values [N]int32

	var wg sync.WaitGroup

	for i := 0; i < N; i++ {
		wg.Add(1)
		i := i
		go func() {
			values[i] = 50 + rand.Int31n(50)
			fmt.Println("Done:", i)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("values:", values)
}
