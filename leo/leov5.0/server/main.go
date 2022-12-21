package main

import (
	"github.com/FelixStarship/go11/leo/leov5.0/server/route"
	"github.com/FelixStarship/go11/leo/leov5.0/znet"
)

func main() {

	s := znet.NewServer("Leo v5.0")

	s.AddRoute(0, &route.PingRoute{})

	s.Server()

}
