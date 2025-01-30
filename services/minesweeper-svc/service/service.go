// File service.go define microservice data for creating gRPC endpoints
package service

import (
	"log/slog"
	"time"

	"github.com/google/uuid"

	"github.com/qiv1ne/minesweeper"
	"github.com/qiv1ne/minesweeper-coop/services/minesweeper-svc/structs"
)

// GameService interface represent main service
type MinesweeperService interface {
	NewMineBoard(config minesweeper.BoardConfig) (minesweeper.MineBoard, error)
	OpenCell(x, y int) (int, error)
	PlaceFlag(x, y int) (int, error)
	NewSeed() int64
}

type minesweeperService struct{}

func NewGameService(config minesweeper.BoardConfig) MinesweeperService {
	slog.Info("creating new game service")
	return minesweeperService{}
}

func (svc minesweeperService) NewGame(config minesweeper.BoardConfig) (minesweeper.MineBoard, error) {
	// creating new board
	board, err := minesweeper.NewMineBoard(config)

	// generating id for this game session
	id := uuid.NewString()

	// define when game session is created
	t := time.Now().Format(time.TimeOnly)

	// create session struct
	session := structs.GameSession{
		BoardId:   id,
		Board:     *board,
		CreatedAt: t,
	}

	return *board, err
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
