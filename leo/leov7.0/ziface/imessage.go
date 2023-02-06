package ziface

type IMessage interface {
	//TLV 格式数据包
	GetDataLen() uint32
	GetMsgID() uint32
	GetData() []byte

	SetMsgID(uint32)
	SetData([]byte)
	SetDataLen(uint32)
}
