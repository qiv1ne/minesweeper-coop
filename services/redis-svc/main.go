package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/qiv1ne/minesweeper-coop/services/redis-svc/endpoints"
	"github.com/qiv1ne/minesweeper-coop/services/redis-svc/service"
	"github.com/qiv1ne/minesweeper-coop/services/redis-svc/transport"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	svc, err := service.NewRedisService(os.Getenv("REDIS_ADDRES"), os.Getenv("REDIS_PASSWORD"))
	if err != nil {
		slog.Error(err.Error(), "REDIS_ADDRES", os.Getenv("REDIS_ADDRES"), "REDIS_PASSWORD", os.Getenv("REDIS_PASSWORD"))
		panic(err)
	}

	getGameHandler := httptransport.NewServer(
		endpoints.MakeGetGameEndpoint(svc),
		transport.DecodeGetGameRequest,
		transport.EncodeResponse,
	)

	saveGameHandler := httptransport.NewServer(
		endpoints.MakeSaveGameEndpoint(svc),
		transport.DecodeSetGameRequest,
		transport.EncodeResponse,
	)
	http.Handle("GET /game", getGameHandler)
	http.Handle("POST /game", saveGameHandler)
	slog.Info("server is listen", "port", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
