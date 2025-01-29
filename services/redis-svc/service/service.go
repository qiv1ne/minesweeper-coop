package service

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/qiv1ne/minesweeper-coop/services/minesweeper-svc/structs"
	"github.com/redis/go-redis/v9"
)

const (
	// It's namespaces for redis database. See docs.
	userSpace = "sessions:"
	gameSpace = "games:"
)

type RedisService interface {
	SaveGame(session structs.GameSession) error
	DeleteSession(id string) error
	GetGame(id string) (structs.GameSession, error)
}

type redisService struct {
	client *redis.Client
}

func NewRedisService(addr, password string) (*redis.Client, error) {
	slog.Info("creating new redis service", "addres", addr)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
	})
	return client, client.Ping(context.Background()).Err()
}

func (r redisService) DeleteGame(id string) error {

}

func (r redisService) SaveGame(session structs.GameSession) error {
	data, err := json.Marshal(session)
	if err != nil {
		return err
	}
	return r.client.Set(context.Background(), gameSpace+session.BoardId, data, time.Hour).Err()
}

func (r redisService) GetGame(id string) (structs.GameSession, error) {
	session := structs.GameSession{}
	result, err := r.client.Get(context.Background(), gameSpace+id)
	if err != nil {
		return session, err
	}
	err = json.Unmarshal(result, session)
	if err != nil {
		return session, err
	}
	return session, err
}
