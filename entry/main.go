package main

import "github.com/FelixStarship/go11/entry/example"

func main() {
	event := example.InitializeEvent()
	event.Start()
}
