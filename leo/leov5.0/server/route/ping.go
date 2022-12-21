package route

import (
	"fmt"
	"github.com/FelixStarship/go11/leo/leov5.0/ziface"
	"github.com/FelixStarship/go11/leo/leov5.0/znet"
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
