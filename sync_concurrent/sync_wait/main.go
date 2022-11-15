package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	const N = 5
	var values [N]int32
	var wgA, wgB sync.WaitGroup
	wgA.Add(N)
	wgB.Add(1)

	for i := 0; i < N; i++ {
		i := i
		go func() {
			wgB.Wait()
			log.Printf("values[%v]=%v \n", i, values[i])
			wgA.Done()
		}()
	}

	for i := 0; i < N; i++ {
		values[i] = 50 + rand.Int31n(50)
	}

	wgB.Done()
	wgA.Wait()
}
