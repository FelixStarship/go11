package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://127.0.0.1:8080/chunked")
	if err != nil {
		log.Fatalf("client send conn err:%+v", err)
	}
	defer res.Body.Close()

	fmt.Println(res.TransferEncoding)

	reader := bufio.NewReader(res.Body)

	for {
		line, err := reader.ReadString('\n')
		if len(line) > 0 {
			fmt.Println(line)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}
