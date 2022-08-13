//go:build wireinject
// +build wireinject

package example

import "github.com/google/wire"

func InitializeEvent() Event {
	wire.Build(NewGreeter, NewEvent, NewMessage)
	return Event{}
}
