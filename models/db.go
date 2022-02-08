package models

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var db *redis.Client
var ctx = context.Background()

func NewDatabase(address string) error {
	db = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})
	if err := db.Ping(ctx).Err(); err != nil {
		return err
	}
	return nil
}
