package ziface

type IServer interface {
	Start()
	Stop()
	Server()
	AddRoute(msgId uint32, route IRoute)
	GetConnMgr() IConnManger
	// 设置该Server的连接创建时Hook函数
	SetOnConnStart(func(IConnection))
	// 设置该Server的连接断开时Hook函数
	SetOnConnStop(func(IConnection))
	// 调用连接OnConnStart Hook函数
	CallOnConnStart(IConnection)
	// 调用连接OnConnStop Hook函数
	CallOnConnStop(IConnection)
}
