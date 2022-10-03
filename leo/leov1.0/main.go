package main

import (
	"fmt"
	"github.com/FelixStarship/go11/leo/leov1.0/znet"
	"net"
	"time"
)

func main() {

	go znet.NewServer("[Leo v1.0]").Start()

	time.Sleep(time.Second * 3)
	fmt.Println("[Leo] Client Start...")
	conn, err := net.Dial("tcp", "127.0.0.1:777")
	if err != nil {
		fmt.Println("client start err:", err, "exit!")
		return
	}

	for {
		_, err := conn.Write([]byte("bibi"))
		if err != nil {
			fmt.Println("write error err", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error")
			return
		}
		fmt.Printf("server call back:%s,cnt:%d\n", buf, cnt)
	}

}
