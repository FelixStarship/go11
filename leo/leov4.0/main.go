package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/FelixStarship/go11/leo/leov4.0/ziface"
	"github.com/FelixStarship/go11/leo/leov4.0/znet"
)

func main() {

	go func() {
		s := znet.NewServer("Leo v4.0")

		s.AddRoute(0, &PingRoute{})

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
		dp := znet.NewDataPack()
		msg, _ := dp.Pack(znet.NewMsgPackage(0, []byte("Leo v0.4 client test")))
		_, err := conn.Write(msg)
		if err != nil {
			log.Fatalf("write error,err:%+v", err)
		}

		headData := make([]byte, dp.GetHeadLen())
		_, err = io.ReadFull(conn, headData)
		if err != nil {
			log.Fatalf("read head error:%+v", err)
		}

		msgHead, err := dp.UnPack(headData)
		if err != nil {
			log.Fatalf("server unpack error:%+v", err)
		}

		if msgHead.GetDataLen() > 0 {
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetDataLen())

			_, err = io.ReadFull(conn, msg.Data)
			if err != nil {
				log.Fatalf("server unpack data error:%+v", err)
			}

			fmt.Println("==> Recv Msg:ID=", msg.ID, " ,len=", msg.DataLen, ",data=", string(msg.Data))
		}

		time.Sleep(time.Second * 1)
	}

}

type PingRoute struct {
	znet.BaseRoute
}

func (p *PingRoute) PreHandler(req ziface.IRequest) {
	//读取客户端数据、在响应给客户端
	fmt.Println("recv from client:msgId=", req.GetMsgID(), "msgData=", string(req.GetData()))
	err := req.GetConnection().SendMsg(0, []byte("ping....ping...."))
	if err != nil {
		fmt.Println("ping call err", err)
	}
}
