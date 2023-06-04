package cache

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
)

type Cache interface {
	Set(key string, value interface{}, ttl time.Duration) error
	Get(key string) (interface{}, error)
	Delete(key string) error
}

type MemoryCache struct {
	client *redis.Client
}

func NewMemoryCache() (*MemoryCache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &MemoryCache{
		client: client,
	}, nil

}
func (c *MemoryCache) Set(key string, value interface{}, ttl time.Duration) error {

	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = c.client.Set(key, data, ttl).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *MemoryCache) Get(key string) (interface{}, error) {
	data, err := c.client.Get(key).Result()
	if err != nil {
		return nil, err
	}

	var value interface{}
	err = json.Unmarshal([]byte(data), &value)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (c *MemoryCache) Delete(key string) error {
	err := c.client.Del(key).Err()
	if err != nil {
		return err
	}

	return nil
}
