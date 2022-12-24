package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		// Handle error
	}

	for i := 0; i <= 1000; i++ {

		// 创建带缓冲的写入器
		writer := bufio.NewWriter(conn)

		// 通过写入器写入数据，这里写入了两个字符串，之间没有间隔
		writer.WriteString("Hello")
		// 数据包之间没有边界、导致服务端读取数据包一次读出
		writer.WriteString("World")

		// 将数据刷到网络连接上
		writer.Flush()

		//conn.Write([]byte("[这里才是一个完整的数据包!]"))
		readbuf := make([]byte, 512)
		dataLen, err := conn.Read(readbuf)
		if err != nil {
			log.Fatal(err)
		}

		if dataLen > 0 {
			fmt.Println(string(readbuf))
		}
	}

}
