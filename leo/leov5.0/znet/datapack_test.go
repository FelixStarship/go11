package znet

import (
	"fmt"
	"io"
	"log"
	"net"
	"runtime"
	"testing"
)

func TestDataPack(t *testing.T) {
	listen, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		log.Fatalf("server listen err:%+v", err)
	}

	go func() {
		for {
			conn, err := listen.Accept()
			if err != nil {
				log.Fatalf("server accept err:%+v", err)
			}
			go func(conn net.Conn) {
				dp := NewDataPack()
				for {
					headData := make([]byte, dp.GetHeadLen())
					_, err := io.ReadFull(conn, headData)
					if err != nil {
						log.Fatalf("read head err:%+v", err)
					}

					msgHead, err := dp.UnPack(headData)

					if msgHead.GetDataLen() > 0 {
						msg := msgHead.(*Message)
						msg.Data = make([]byte, msg.GetDataLen())

						_, err := io.ReadFull(conn, msg.Data)
						if err != nil {
							log.Fatalf("read data err:%+v", err)
						}
						fmt.Println("===> Recv Msg: ID=", msg.ID, ",len=", msg.DataLen, ",Data=", string(msg.Data))
					}

				}
			}(conn)

			fmt.Println(runtime.NumGoroutine())
		}
	}()

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		log.Fatalf("client dial err:%+v", err)
	}

	dp := NewDataPack()

	msg1 := &Message{
		ID:      0,
		DataLen: 3,
		Data:    []byte{'l', 'e', 'o'},
	}

	sendData1, err := dp.Pack(msg1)

	msg2 := &Message{
		ID:      1,
		DataLen: 10,
		Data:    []byte{'h', 'e', 'l', 'l', 'o'},
	}

	sendData2, err := dp.Pack(msg2)

	conn.Write(append(sendData1, sendData2...))

	select {}
}
