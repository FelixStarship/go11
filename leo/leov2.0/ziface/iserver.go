package ziface

type IServer interface {
	Start()
	Stop()
	Server()
	AddRoute(route IRoute)
}
