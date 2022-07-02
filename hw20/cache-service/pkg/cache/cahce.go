package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/kevin-glare/hardcode-dev-go/hw20/common/pkg/model"
	"time"
)

type Cache struct {
	redis *redis.Client
}

var cacheDuration = 24 * 7 * time.Hour

func NewCache(addr string) *Cache {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	return &Cache{redis: redisClient}
}

func (c *Cache) Set(input model.Link) error {
	key := linkKey(input.ShortUrl)

	val, err := json.Marshal(input)
	if err != nil {
		return err
	}

	return c.redis.Set(key, val, cacheDuration).Err()
}

func (c *Cache) Get(shortLink string) (*model.Link, error) {
	var link model.Link
	cmd := c.redis.Get(linkKey(shortLink))

	err := json.Unmarshal([]byte(cmd.Val()), &link)
	if err != nil {
		return nil, err
	}

	if link.Url == "" {
		return nil, errors.New("Link not found")
	}

	return &link, nil
}

func linkKey(key string) string {
	return fmt.Sprintf("links/%s", key)
}
