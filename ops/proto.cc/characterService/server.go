package main

import (
	"context"
	"strconv"

	"github.com/dtherhtun/Learning-go/ops/proto.cc/characterService/character"
	pb "github.com/dtherhtun/Learning-go/ops/proto.cc/go/character"
)

type CharacterServer struct {
	pb.UnimplementedCharacterServiceServer
}

func newServer() *CharacterServer {
	return &CharacterServer{}
}

func (cs *CharacterServer) GetCharacters(ctx context.Context, request *pb.AllCharactersRequest) (*pb.AllCharactersResponse, error) {
	characters, err := character.GetCharacters()
	if err != nil {
		return nil, err
	}

	var results []*pb.Result

	for _, character := range characters {
		i, _ := strconv.ParseInt(character.ID, 10, 32)

		result := &pb.Result{
			Character: &pb.Character{
				Id:          int32(i),
				Name:        character.Name,
				Category:    character.Category,
				Bio:         character.Bio,
				Description: character.Description,
			},
		}

		results = append(results, result)
	}

	return &pb.AllCharactersResponse{
		Header:  request.GetHeader(),
		Results: results,
	}, nil
}

func (cs *CharacterServer) GetCharacterById(ctx context.Context, request *pb.GetCharacterRequest) (*pb.GetCharacterResponse, error) {
	response, err := character.GetCharacterById(strconv.Itoa(int(request.GetCharacterId())))
	if err != nil {
		return nil, err
	}

	i, _ := strconv.ParseInt(response.ID, 10, 32)

	return &pb.GetCharacterResponse{
		Header: request.GetHeader(),
		Result: &pb.Result{
			Character: &pb.Character{
				Id:          int32(i),
				Name:        response.Name,
				Category:    response.Category,
				Bio:         response.Bio,
				Description: response.Description,
			},
		},
	}, nil
}
