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

func NewGameService(config minesweeper.BoardConfig) (MinesweeperService, error) {
	slog.Info("creating new game service")
	board, err := minesweeper.NewMineBoard(config)
	return minesweeperService{
		Board: board,
	}, err
}

func (svc minesweeperService) OpenCell(x, y int) (int, error) {
	slog.Info("opening cell", "x", x, "y", y)
	return svc.Board.OpenCell(x, y)
}
func (svc minesweeperService) PlaceFlag(x, y int) (int, error) {
	slog.Info("placing flag", "x", x, "y", y)
	return svc.Board.PlaceFlag(x, y)
}
func (svc minesweeperService) NewSeed() int64 {
	slog.Info("generating new seed")
	return minesweeper.NewSeed()
}
