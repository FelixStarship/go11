package main

import (
	"github.com/FelixStarship/go11/leo/leov7.0/server/route"
	"github.com/FelixStarship/go11/leo/leov7.0/znet"
)

func main() {

	s := znet.NewServer("Leo v7.0")

	s.SetOnConnStart(route.DoConnectionBegin)

	s.SetOnConnStop(route.DoConnectionLost)

	s.AddRoute(0, &route.PingRoute{})

	s.Server()

}
