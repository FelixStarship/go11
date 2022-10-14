package ziface

type IServer interface {
	Start()
	Stop()
	Server()
	AddRoute(msgId uint32, route IRoute)
}
