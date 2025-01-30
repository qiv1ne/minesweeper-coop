package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/qiv1ne/minesweeper-coop/services/redis-svc/service"
	"github.com/qiv1ne/minesweeper-coop/services/redis-svc/structs"
)

func MakeGetGameEndpoint(svc service.RedisService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(structs.GetGameRequest)
		game, err := svc.GetGame(req.Id)
		if err != nil {
			return structs.GetGameResponse{
				Game: game,
				Err:  err,
			}, nil
		}
		return structs.GetGameResponse{
			Game: game,
			Err:  err,
		}, nil
	}
}

func MakeSaveGameEndpoint(svc service.RedisService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(structs.SaveGameRequest)
		err := svc.SaveGame(req.Game)
		if err != nil {
			return structs.SaveGameResponse{
				Err: err,
			}, nil
		}
		return structs.SaveGameResponse{
			Err: err,
		}, nil
	}
}
