package ziface

/*
  连接管理模块
*/
type IConnManger interface {
	Add(conn IConnection)
	Remove(conn IConnection)
	Get(connID uint32) (IConnection, error)
	Len() int
	ClearConn()
}
