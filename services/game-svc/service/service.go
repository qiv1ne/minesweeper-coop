// File service.go define microservice data for creating gRPC endpoints
package service

import (
	"github.com/Sinoverg/game-svc/pkg"
	"github.com/rs/zerolog/log"
)

// Define the GameService interface for go-kit.
type GameService interface {
	// CreateMineMap function will create the MineMap struct.
	// It should accept mine's count, height of board and width of board
	CreateMineBoard(opts pkg.BoardConfig) (*pkg.MineBoard, error)
	RevealAll(board [][]pkg.Cell) ([][]pkg.Cell, error)
	Reveal(i, j int) ()
}

type gameService struct{}

func NewGameService() GameService {
	log.Info().Msg("Creating new game service.")
	return &gameService{}
}

func (gameService) CreateMineBoard(opts pkg.BoardConfig) (*pkg.MineBoard, error) {
	log.Info().Msg("Creating new minesweeper board")
	board, err := pkg.CreateBoard(opts)
	if err != nil {
		return nil, err
	}
	realBoard, err := pkg.RevealAll(board)
	if err != nil {
		return nil, err
	}
	mb := &pkg.MineBoard{
		RealBoard: realBoard,
		UserBoard: board,
	}
	return mb, nil
}

func (gameService) RevealAll(b [][]pkg.Cell) ([][]pkg.Cell, error) {
	return pkg.RevealAll(b)
}
