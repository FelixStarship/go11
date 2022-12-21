package main

import (
	"fmt"
	"github.com/FelixStarship/go11/leo/leov5.0/znet"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	fmt.Println("[Leo] Client Start...")
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client start err:", err, "exit!")
		return
	}

	for {
		dp := znet.NewDataPack()
		msg, _ := dp.Pack(znet.NewMsgPackage(0, []byte("Leo v0.5 client test")))
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
