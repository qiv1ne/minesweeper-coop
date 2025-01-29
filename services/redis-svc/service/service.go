package service

import (
	minesweeperSvc "github.com/qiv1ne/minesweeper-coop/services/minesweeper-svc/structs"
)

type RedisService interface {
	SaveGame(session minesweeperSvc.GameSession)
}
