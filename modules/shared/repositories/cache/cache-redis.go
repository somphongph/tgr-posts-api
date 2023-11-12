package cache

import (
	"os"
	"strconv"
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

func InitCache() *RedisStore {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASS"), // no password set
		DB:       0,                       // use default DB
	})

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
	// Set time in second
	intVar, err := strconv.Atoi(os.Getenv("REDIS_SHORT_CACHE"))
	if err != nil {
		return err
	}
	dur := time.Duration(intVar) * time.Second
	err = c.Set(key, value, dur).Err()

	return err
}

func (c *RedisStore) SetLongCache(key string, value interface{}) error {
	// Set time in second
	intVar, err := strconv.Atoi(os.Getenv("REDIS_LONG_CACHE"))
	if err != nil {
		return err
	}

	dur := time.Duration(intVar) * time.Second
	err = c.Set(key, value, dur).Err()

	return err
}
