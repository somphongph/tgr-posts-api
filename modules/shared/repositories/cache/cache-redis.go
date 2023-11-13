package cache

import (
	"tgr-posts-api/configs"
	"time"

	"github.com/go-redis/redis"
)

type Cached interface {
	GetCache(string) (string, error)
	SetCache(string, interface{}, int) error
	SetShortCache(string, interface{}) error
	SetLongCache(string, interface{}) error
}

type RedisStore struct {
	*redis.Client
}

var (
	shortCache = 0
	longCache  = 0
)

func InitCache(cfg *configs.Redis) *RedisStore {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Host,
		Password: cfg.Pass, // no password set
		DB:       0,        // use default DB
	})

	shortCache = cfg.ShortCache
	longCache = cfg.LongCache

	return &RedisStore{rdb}
}

func (c *RedisStore) GetCache(key string) (string, error) {
	val, err := c.Get(key).Result()

	return val, err
}

func (c *RedisStore) SetCache(key string, value interface{}, duration int) error {
	// Set time in second
	dur := time.Duration(duration) * time.Second
	err := c.Set(key, value, dur).Err()

	return err
}

func (c *RedisStore) SetShortCache(key string, value interface{}) error {
	dur := time.Duration(shortCache) * time.Second
	err := c.Set(key, value, dur).Err()

	return err
}

func (c *RedisStore) SetLongCache(key string, value interface{}) error {
	dur := time.Duration(longCache) * time.Second
	err := c.Set(key, value, dur).Err()

	return err
}
