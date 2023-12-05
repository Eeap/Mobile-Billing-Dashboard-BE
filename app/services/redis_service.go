package services

import (
	"context"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
	"main/app/models"
	"main/platform/cache"
)

func SetAlertMessage(message *models.AlertMessage, email string) error {
	connRedis, err := cache.RedisConnection()
	if err != nil {
		// Return status 500 and Redis connection error.
		return err
	}
	messageData, err := json.Marshal(message)
	if err != nil {
		return err
	}
	// 먼저 get 요청을 하고 있으면 추가하고 없으면 하나 만드는 방식으로 가야할듯
	err = connRedis.LRange(context.Background(), email, 0, -1).Err()
	if err != nil {
		return err
	}
	if err == redis.Nil {
		// set new key
		err = connRedis.LPush(context.Background(), email, string(messageData)).Err()
		if err != nil {
			// Return Redis connection error.
			return err
		}
	} else if err != nil {
		return err
	} else {
		// update key
		err = connRedis.RPush(context.Background(), email, string(messageData)).Err()
		if err != nil {
			// Return Redis connection error.
			return err
		}
	}
	return nil
}

func GetAlertMessages(email string) (*[]models.AlertMessage, error) {
	connRedis, err := cache.RedisConnection()
	if err != nil {
		// Return status 500 and Redis connection error.
		return nil, err
	}
	res, err := connRedis.LRange(context.Background(), email, 0, -1).Result()
	if err != nil {
		return nil, err
	}
	resData := []models.AlertMessage{}
	for _, v := range res {
		var data models.AlertMessage
		err = json.Unmarshal([]byte(v), &data)
		if err != nil {
			return nil, err
		}
		resData = append(resData, data)
	}
	return &resData, nil
}

func DeleteAlertMessages(email string) error {
	connRedis, err := cache.RedisConnection()
	if err != nil {
		log.Println(err)
		// Return status 500 and Redis connection error.
		return err
	}
	err = connRedis.Del(context.Background(), email).Err()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
