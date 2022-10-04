package main

import (
	"fmt"
	"net"
	"time"

	"github.com/FelixStarship/go11/leo/leov2.0/ziface"
	"github.com/FelixStarship/go11/leo/leov2.0/znet"
)

func main() {

	go func() {
		s := znet.NewServer("[Leo v1.0]")

		s.AddRoute(&PingRoute{})

		s.Server()
	}()

	time.Sleep(time.Second * 3)
	fmt.Println("[Leo] Client Start...")
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
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
		time.Sleep(time.Second * 1)
	}

}

type PingRoute struct {
	znet.BaseRoute
}

func (p *PingRoute) PreHandler(req ziface.IRequest) {
	_, err := req.GetConnection().GetTCPConnection().Write([]byte("before ping..."))
	if err != nil {
		fmt.Println("ping call err", err)
	}
}
