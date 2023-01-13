package znet

import (
	"fmt"
	"github.com/FelixStarship/go11/leo/leov5.0/ziface"
	"strconv"
)

type MsgHandler struct {
	Apis           map[uint32]ziface.IRoute
	WorkerPoolSize uint32
	TaskQueue      []chan ziface.IRequest
}

func NewMsgHandle() *MsgHandler {
	return &MsgHandler{
		Apis:           make(map[uint32]ziface.IRoute),
		WorkerPoolSize: 10,
		TaskQueue:      make([]chan ziface.IRequest, 10),
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

func (s *MsgHandler) StartOneWorker(workerID int, taskQueue chan ziface.IRequest) {
	fmt.Println("Worker ID =,", workerID, "is started.")
	for {
		select {
		case request := <-taskQueue:
			s.DoMsgHandler(request)
		}
	}
}

func (s *MsgHandler) StartWorkerPool() {
	for i := 0; i < int(s.WorkerPoolSize); i++ {
		s.TaskQueue[i] = make(chan ziface.IRequest, 1024)
		go s.StartOneWorker(i, s.TaskQueue[i])
	}
}

func (s *MsgHandler) SendMsgToTaskQueue(request ziface.IRequest) {
	workerID := request.GetConnection().GetConnID() % s.WorkerPoolSize
	fmt.Println("Add ConnID=", request.GetConnection().GetConnID(), ",request msgID=", request.GetMsgID(), ",to workerID=", workerID)
	s.TaskQueue[workerID] <- request
}
