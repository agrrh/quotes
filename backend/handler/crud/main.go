// Package crud - CRUD operations
package crud

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// Client - Client to access CRUD operations
type Client struct {
	RedisClient redis.Client
	Context     context.Context
}

// CreateClient - Create CRUD Client
func CreateClient(addr string, password string, db int) Client {
	var ctx = context.Background()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	client := Client{
		RedisClient: *redisClient,
		Context:     ctx,
	}

	return client
}

// Create - Client.Create() operation
func (c Client) Create(key string, value string) (bool, error) {
	err := c.RedisClient.Set(c.Context, key, value, 0).Err()
	if err != nil {
		return false, err
	}

	return true, nil
}

// Read - Client.Read() operation
func (c Client) Read(key string) (string, error) {
	val, err := c.RedisClient.Get(c.Context, key).Result()

	if err != nil {
		return "", err
	}

	return val, nil
}

// Update - Client.Update() operation
func (c Client) Update(key string, value string) (bool, error) {
	return c.Create(key, value)
}

// Delete - Client.Delete() operation
func (c Client) Delete(key string) (bool, error) {
	err := c.RedisClient.Del(c.Context, key).Err()
	if err != nil {
		return false, err
	}

	return true, nil
}

// List - Client.List() operation
func (c Client) List(pattern string) ([]string, error) {
	val, err := c.RedisClient.Keys(c.Context, pattern).Result()

	if err != nil {
		return make([]string, 0), err
	}

	return val, nil
}
