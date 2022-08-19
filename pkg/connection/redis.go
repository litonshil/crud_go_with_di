package connection

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/litonshil/crud_go_echo/config"
)

var redisClient *redis.Client

func ConnectRedis() {

	redisClient = redis.NewClient(&redis.Options{
		Addr:     config.GetConfig().RedisHost + ":" + config.GetConfig().RedisPort,
		Password: config.GetConfig().RedisPass,
		DB:       0,
	})
	red, err := redisClient.Ping(context.Background()).Result()
	fmt.Println(red)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("redis connection successful...")
}

func Redis() *redis.Client {
	return redisClient
}
