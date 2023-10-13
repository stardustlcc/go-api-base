package redis

import (
	"go-api-base/pkg/global"
	"time"

	"github.com/go-redis/redis/v7"
)

var _ Repo = (*redisRepo)(nil)

type Repo interface {
	Set(key, value string, ttl time.Duration) error
	Get(key string) (string, error)
	Close() error
}

type redisRepo struct {
	rdb *redis.Client
}

func NewRedis() (Repo, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:         global.Conf.Redis.Host,
		Password:     global.Conf.Redis.PassWord,
		DB:           global.Conf.Redis.DB,
		PoolSize:     global.Conf.Redis.PoolSize,
		MaxRetries:   global.Conf.Redis.MaxRetries,
		MinIdleConns: global.Conf.Redis.MinIdleConns,
	})

	if err := rdb.Ping().Err(); err != nil {
		return nil, err
	}

	return &redisRepo{rdb: rdb}, nil
}

func (r *redisRepo) Set(key string, val string, ttl time.Duration) error {
	return r.rdb.Set(key, val, ttl).Err()
}

func (r *redisRepo) Get(key string) (string, error) {
	return r.rdb.Get(key).Result()
}

func (r *redisRepo) Close() error {
	return r.rdb.Close()
}
