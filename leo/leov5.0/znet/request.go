package znet

import "github.com/FelixStarship/go11/leo/leov5.0/ziface"

type Request struct {
	conn ziface.IConnection
	msg  ziface.IMessage
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgID()
}
