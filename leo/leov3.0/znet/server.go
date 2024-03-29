package znet

import (
	"fmt"
	"github.com/FelixStarship/go11/leo/leov2.0/ziface"
	"net"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
	Route     ziface.IRoute
}

func (s *Server) Start() {
	fmt.Printf("[START] Server listener at IP:%s,Port:%d,is starting\n", s.IP, s.Port)

	go func() {
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
			go NewConnection(conn, cid, s.Route).Start()
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
}

func (s *Server) AddRoute(route ziface.IRoute) {
	s.Route = route
}

func NewServer(name string) ziface.IServer {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      7777,
	}
}
