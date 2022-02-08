package models

import (
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

type SaveURL struct {
	Id    string `json:"id" redis:"id"`
	URL   string `json:"link" redis:"link"`
	Count int    `json:"count" redis:"count"`
	// Expires string `json:"expires" redis:"expires"`
}

func SaveURLToDB(id string, url string) error {
	_, err := db.Get(ctx, id).Result()
	switch {
	//key doesn't exist
	case err == redis.Nil:
		obj := SaveURL{
			Id:    id,
			URL:   url,
			Count: 0,
		}
		p, err := json.Marshal(&obj)
		if err != nil {
			return err
		}
		err = db.Set(ctx, id, p, 0).Err()
		if err != nil {
			return err
		}
		// key exists, or some other error
	case err != nil:
		return err
	}
	return nil
}
