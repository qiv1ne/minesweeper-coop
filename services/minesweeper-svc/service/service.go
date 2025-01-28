// File service.go define microservice data for creating gRPC endpoints
package service

import (
	"log/slog"

	"github.com/qiv1ne/minesweeper"
)

// GameService interface represent main service
type MinesweeperService interface {
	OpenCell(x, y int) (int, error)
	PlaceFlag(x, y int) (int, error)
	NewSeed() int64
}

type minesweeperService struct {
	Board *minesweeper.MineBoard
}

func NewGameService() MinesweeperService {
	slog.Info("creating new game service")
	return gameService{}
}

func (gameService) CreateMineBoard(config minesweeper.BoardConfig) (*minesweeper.MineBoard, error) {
	slog.Info("creating new mine board")
	board, err := minesweeper.NewMineBoard(config)
	return board, err
}

func (gameService) OpenCell(x, y int) (int, error) {
	return minesweeper
}
func (gameService) PlaceFlag(x, y int) (int, error) {

}
func (gameService) NewSeed() int64 {

}
