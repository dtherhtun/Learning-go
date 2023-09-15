package main

import (
	pb "github.com/dtherhtun/Learning-go/ops/proto.cc/go/character"
)

type CharacterServer struct {
	pb.UnimplementedCharacterServiceServer
}
