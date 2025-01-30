package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/qiv1ne/minesweeper-coop/services/redis-svc/structs"
)

func DecodeGetGameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request structs.GetGameRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeSetGameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request structs.SaveGameRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
