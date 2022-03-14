package global

import (
	"github.com/go-redis/redis"
)

var DB redis.Client

func ConnectToDB() {

	client := redis.NewClient(&redis.Options{
		Addr:     REDIS_URL,
		Password: "",
		DB:       0,
	})

	DB = *client
}
