package database

import (
	"encoding/json"
	"fmt"
	"time"

	proto "github.com/Soyaka/microlearn-user/api/proto/gen"
	"github.com/go-redis/redis"
)

var RediClient *RedisClient

type RedisClient struct {
	Client *redis.Client
}

func NewCache() *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6377",
		})
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println("Failed to ping Redis server:", err)
		return nil
	}
	fmt.Println("Redis server is up and running:", pong)

	return &RedisClient{Client: client}
}

func (*RedisClient) Close() error {
	return nil
}

func (*RedisClient) Ping() error {
	return nil
}

func (c *RedisClient) AddUserToCache(req *proto.User, expiration time.Duration) error {
	userJSON, err := json.Marshal(req)
	if err != nil {
		return err
	}
	_, err = c.Client.Set(req.Email, userJSON, expiration).Result()
	return err
}

func (c *RedisClient) GetUserFromCache(email string) (*proto.User, error) {

	userJSON, err := c.Client.Get(email).Bytes()
	if err != nil {
		return nil, err
	}
	var user proto.User
	err = json.Unmarshal(userJSON, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
