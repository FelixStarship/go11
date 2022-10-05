package znet

type Message struct {
	ID      uint32
	DataLen uint32
	Data    []byte
}

func NewMsgPackage(id uint32, data []byte) *Message {
	return &Message{
		ID:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

func (m *Message) GetDataLen() uint32 {
	return m.DataLen
}

func (m *Message) GetMsgID() uint32 {
	return m.ID
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetDataLen(len uint32) {
	m.DataLen = len
}

func (m *Message) SetMsgID(msgID uint32) {
	m.ID = msgID
}

func (m *Message) SetData(data []byte) {
	m.Data = data
}
