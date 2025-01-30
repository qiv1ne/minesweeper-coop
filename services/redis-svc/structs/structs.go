package structs

import (
	"github.com/qiv1ne/minesweeper-coop/services/minesweeper-svc/structs"
)

type GetGameRequest struct {
	Id string `json:"id"`
}

type GetGameResponse struct {
	Game structs.GameSession `json:"game"`
	Err  error               `json:"err,omitempty"`
}

type SaveGameRequest struct {
	Game structs.GameSession `json:"game"`
}

type SaveGameResponse struct {
	Err error `json:"err,omitempty"`
}
