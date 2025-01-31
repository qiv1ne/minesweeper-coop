package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/qiv1ne/minesweeper-coop/services/redis-svc/endpoints"
	"github.com/qiv1ne/minesweeper-coop/services/redis-svc/middleware"
	"github.com/qiv1ne/minesweeper-coop/services/redis-svc/service"
	"github.com/qiv1ne/minesweeper-coop/services/redis-svc/transport"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	klog "github.com/go-kit/log"
)

func main() {
	svc, err := service.NewRedisService(os.Getenv("REDIS_ADDRES"), os.Getenv("REDIS_PASSWORD"))
	if err != nil {
		slog.Error(err.Error(), "REDIS_ADDRES", os.Getenv("REDIS_ADDRES"), "REDIS_PASSWORD", os.Getenv("REDIS_PASSWORD"))
		panic(err)
	}

	logger := klog.NewLogfmtLogger(os.Stderr)

	var getGame endpoint.Endpoint
	getGame = endpoints.MakeGetGameEndpoint(svc)
	getGame = middleware.LoggingMiddleware(klog.With(logger, "method", "get game"))(getGame)

	getGameHandler := httptransport.NewServer(
		getGame,
		transport.DecodeGetGameRequest,
		transport.EncodeResponse,
	)

	var saveGame endpoint.Endpoint
	saveGame = endpoints.MakeSaveGameEndpoint(svc)
	saveGame = middleware.LoggingMiddleware(klog.With(logger, "method", "set game"))(saveGame)

	saveGameHandler := httptransport.NewServer(
		saveGame,
		transport.DecodeSetGameRequest,
		transport.EncodeResponse,
	)
	http.Handle("GET /game", getGameHandler)
	http.Handle("POST /game", saveGameHandler)
	slog.Info("server is listen", "port", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
