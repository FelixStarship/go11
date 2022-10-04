package znet

import "github.com/FelixStarship/go11/leo/leov2.0/ziface"

type BaseRoute struct{}

func (b *BaseRoute) PreHandler(req ziface.IRequest) {}

func (b BaseRoute) Handler(req ziface.IRequest) {}

func (b BaseRoute) PostHandler(req ziface.IRequest) {}
