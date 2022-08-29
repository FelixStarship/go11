package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"runtime"
	"time"
)

func main() {

	s := time.Minute * 60 * 24 * 7

	fmt.Println(s)

	var g errgroup.Group // 声明一个group实例
	var urls = []string{
		"https://segmentfault.com/",
		"http://2",
		"http://3",
	}
	for _, url := range urls { // 分别获取网站内容
		url := url          // url是局部变量，for循环中对多个协程传递值时，需要重新进行赋值
		g.Go(func() error { // group 的go方法，启一个协程去执行代码
			// Fetch the URL.
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
		fmt.Println(runtime.NumGoroutine())
	}
	if err := g.Wait(); err == nil { // group 的wait方法，等待上面的 g.go的协程执行完成，并且可以接受错误
		fmt.Println("Successfully fetched all URLs.")
	} else {
		fmt.Println(err)
	}
}
