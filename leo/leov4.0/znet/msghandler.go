package znet

import (
	"github.com/FelixStarship/go11/leo/leov4.0/ziface"
	"strconv"
)

type MsgHandler struct {
	Apis map[uint32]ziface.IRoute
}

func NewMsgHandle() *MsgHandler {
	return &MsgHandler{
		Apis: make(map[uint32]ziface.IRoute),
	}
}

func (s *MsgHandler) AddRoute(msgId uint32, route ziface.IRoute) {
	if _, ok := s.Apis[msgId]; ok {
		panic("repeated api , msgId =" + strconv.Itoa(int(msgId)))
	}
	s.Apis[msgId] = route
}

func (s *MsgHandler) DoMsgHandler(request ziface.IRequest) {
	handler, ok := s.Apis[request.GetMsgID()]
	if !ok {
		return
	}
	handler.PreHandler(request)
	handler.Handler(request)
	handler.PostHandler(request)
}
