package ziface

type IMsgHandle interface {
	DoMsgHandler(request IRequest)
	AddRoute(msgId uint32, route IRoute)
	StartWorkerPool()
	SendMsgToTaskQueue(request IRequest)
}
