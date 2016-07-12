package util

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

// RedisSettings identifies uniquely one redis connection
type RedisSettings struct {
	Host     string
	Password string
}

var pool *redis.Pool

func newPool(server string, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, authErr := c.Do("AUTH", password); authErr != nil {
					c.Close()
					return nil, authErr
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

// GetConnection creates and returns a new redis connection pool based on the given settings
func GetConnection(settings RedisSettings) redis.Conn {
	if pool == nil {
		pool = newPool(settings.Host, settings.Password)
	}
	return pool.Get()
}