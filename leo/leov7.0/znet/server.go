package znet

import (
	"fmt"
	"net"

	"github.com/FelixStarship/go11/leo/leov7.0/ziface"
)

type Server struct {
	Name       string
	IPVersion  string
	IP         string
	Port       int
	msgHandler ziface.IMsgHandle
	ConnMgr    ziface.IConnManger
	// 函数一等公民
	OnConnStart func(ziface.IConnection)
	OnConnStop  func(ziface.IConnection)
}

func (s *Server) Start() {
	fmt.Printf("[START] Server listener at IP:%s,Port:%d,is starting\n", s.IP, s.Port)

	go func() {

		// 启动工作池
		s.msgHandler.StartWorkerPool()

		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err:", err)
			return
		}
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}
		fmt.Println("start Leo server ", s.Name, "success,now listening...")
		var cid uint32
		cid = 0
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}

			/*
				设置服务器最大连接数限制、超过最大连接数直接关闭当前连接
			*/
			if s.GetConnMgr().Len() > 3 {
				conn.Close()
				continue
			}
			go NewConnection(s, conn, cid, s.msgHandler).Start()
			cid++
		}
	}()
}

func (s *Server) Server() {
	s.Start()
	select {}
}

func (s *Server) Stop() {
	fmt.Println("[STOP] Leo server,name", s.Name)

	//清理连接信息
	s.GetConnMgr().ClearConn()
}

func (s *Server) AddRoute(msgId uint32, route ziface.IRoute) {
	s.msgHandler.AddRoute(msgId, route)
}

func (s *Server) GetConnMgr() ziface.IConnManger {
	return s.ConnMgr
}

func (s *Server) SetOnConnStart(hookFunc func(ziface.IConnection)) {
	s.OnConnStart = hookFunc
}

func (s *Server) SetOnConnStop(hookFunc func(ziface.IConnection)) {
	s.OnConnStop = hookFunc
}

func (s *Server) CallOnConnStart(conn ziface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("------> CallOnConnStart...")
		s.OnConnStart(conn)
	}
}

func (s *Server) CallOnConnStop(conn ziface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("-----> CallOnConnStop...")
		s.OnConnStop(conn)
	}
}

func NewServer(name string) ziface.IServer {
	return &Server{
		Name:       name,
		IPVersion:  "tcp4",
		IP:         "0.0.0.0",
		Port:       7777,
		msgHandler: NewMsgHandle(),
		ConnMgr:    NewConnManger(),
	}
}
