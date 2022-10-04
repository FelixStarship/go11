package ziface

type IRoute interface {
	PreHandler(request IRequest)
	Handler(request IRequest)
	PostHandler(request IRequest)
}
