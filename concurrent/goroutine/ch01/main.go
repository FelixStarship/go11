package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func job() error {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*1)
	done := make(chan struct{}, 1)
	go func() {
		time.Sleep(time.Millisecond * 2000)
		done <- struct{}{}
	}()
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return fmt.Errorf("超时了")
	}
}

func main() {
	for i := 0; i <= 50; i++ {
		go job()
	}
	for {
		time.Sleep(time.Second)
		fmt.Println(runtime.NumGoroutine())
	}
}
