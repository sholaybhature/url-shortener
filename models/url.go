package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type URLObj struct {
	Id       string        `json:"id"`
	URL      string        `json:"link"`
	Count    int           `json:"count"`
	Visitors []VisitorsObj `json:"visitors"`
}

type VisitorsObj struct {
	Ip     string    `json:"ip"`
	Time   time.Time `json:"time"`
	Device string    `json:"device"`
}

func SaveURLToDB(id string, url string) error {
	_, err := db.Get(ctx, id).Result()
	switch {
	//key doesn't exist
	case err == redis.Nil:
		obj := URLObj{
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

func GetURLFromDB(id string) (*URLObj, error) {
	val, err := db.Get(ctx, id).Result()

	if err != nil {
		return nil, err
	}
	obj := URLObj{}
	json.Unmarshal([]byte(val), &obj)
	return &obj, nil
}

func UpdateURLInDB(id string, ip string, time time.Time, device string) (*URLObj, error) {
	val, err := db.Get(ctx, id).Result()
	fmt.Println(ip, time)
	if err != nil {
		return nil, err
	}
	obj := URLObj{}
	visitorInfo := VisitorsObj{
		Ip:     ip,
		Time:   time,
		Device: device,
	}
	json.Unmarshal([]byte(val), &obj)
	obj.Count += 1
	obj.Visitors = append(obj.Visitors, visitorInfo)
	p, err := json.Marshal(&obj)
	err = db.Set(ctx, id, p, 0).Err()
	if err != nil {
		return nil, err
	}

	return &obj, nil
}
