package route

import (
	"fmt"
	"github.com/FelixStarship/go11/leo/leov7.0/ziface"
	"github.com/FelixStarship/go11/leo/leov7.0/znet"
)

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

func DoConnectionBegin(conn ziface.IConnection) {
	fmt.Println("DoConnectionBegin is Called...")
	if err := conn.SendMsg(1, []byte("DoConnectionBegin BEGIN...")); err != nil {
		fmt.Println(err)
	}
}

func DoConnectionLost(conn ziface.IConnection) {
	fmt.Println("DoConnectionLost is Called ...")
}
