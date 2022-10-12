package ziface

type IRequest interface {
	GetConnection() IConnection
	GetData() []byte
}
