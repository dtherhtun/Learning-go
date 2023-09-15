package main

import "github.com/dtherhtun/Learning-go/ops/proto.cc/characterService/character"

func main() {
	a := character.App{}
	a.Initialize()
	a.Run()
}
