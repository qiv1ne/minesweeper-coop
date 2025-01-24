package main

import (
	"github.com/Sinoverg/game-svc/pkg"
	"github.com/Sinoverg/game-svc/service"
)

func main(){
	gamesvc := service.NewGameService()
	board, err :=gamesvc.CreateMineBoard(pkg.BoardConfig{
		Mines: 1000,
		Width: 10,
		Height: 10,
	})
	if err != nil {
		panic(err)
	}
	board.
	pkg.PrintBroadGracefully(board.RealBoard)
}
