package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log/slog"
	"net/http"

	"github.com/qiv1ne/minesweeper-coop/services/redis-svc/service"
	"github.com/qiv1ne/minesweeper-coop/services/redis-svc/structs"
)

func GetGameHandler(svc service.RedisService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			slog.Error("can't read request body", "error", err, "body", r.Body)
			http.Error(w, "can't read request body: "+err.Error(), http.StatusBadRequest)
			return
		}
		var getGameRequest structs.GetGameRequest
		json.Unmarshal(body, &getGameRequest)
		if err != nil {
			slog.Error("invalid json format", "error", err)
			http.Error(w, "invalid json format: "+err.Error(), http.StatusBadRequest)
			return
		}

		game, err := svc.GetGame(getGameRequest.Id)
		getGameResponse := structs.GetGameResponse{
			Game: game,
			Err:  err,
		}
		response, err := json.Marshal(getGameResponse)
		if err != nil {
			slog.Error("can't marshal response", "error", err)
			http.Error(w, "can't marshal response: "+err.Error(), http.StatusBadRequest)
			return
		}
		w.Write(response)
	}
}

func SetGameHandler(svc service.RedisService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			slog.Error("can't read request body", "error", err, "body", r.Body)
			http.Error(w, "can't read request body: "+err.Error(), http.StatusBadRequest)
			return
		}

	}
}
