package cache

import (
	"belajar-golang-clean-architecture/model/domain"
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisCache struct {
	Client *redis.Client
	Expired int
}

func NewRedisCache(client *redis.Client, expired int) ProductCache {
	return &redisCache{
		Client: client,
		Expired: expired,
	}
}

func (cache *redisCache) Set(ctx context.Context, key string, value interface{}) error {

	json, err := json.Marshal(value)
	if err != nil {
		return err
	}

	cache.Client.Set(ctx, key, json, time.Duration(cache.Expired)*time.Minute)
	return nil	
}

func (cache *redisCache) Get(ctx context.Context, key string) ([]domain.Product, error) {

	val, err := cache.Client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var products []domain.Product
	err = json.Unmarshal([]byte(val), &products)
	if err != nil {
		return nil, err
	}
	return products, nil
}