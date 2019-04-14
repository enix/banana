package services

import (
	"errors"

	"enix.io/banana/src/logger"
	"github.com/go-redis/redis"
)

// Db : Use this API to interact with redis
var Db *redis.Client

// OpenDatabaseConnection : Connect to redis databae
//													Calls such as DbGet and DbSet will crash if called before this
func OpenDatabaseConnection() error {
	Db = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	pong, err := Db.Ping().Result()
	if err != nil || pong != "PONG" {
		return errors.New("failed to connect to redis database")
	}

	logger.Log("etablished connection with redis database")
	return nil
}
