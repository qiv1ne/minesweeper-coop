package structs

import (
	"time"

	"github.com/qiv1ne/minesweeper"
)

type GameSession struct {
	BoardId   string                `json:"board_id"`
	Board     minesweeper.MineBoard `json:"board"`
	CreatedAt time.Time             `json:"created_at"`
}
