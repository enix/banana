package services

import (
	"encoding/json"
	"errors"

	"enix.io/banana/src/logger"
	"github.com/go-redis/redis"
)

// Db : Use this API to interact with redis
var Db *redis.Client

// DbSet : Convenience function to avoid manual JSON encoding
func DbSet(key string, value interface{}) error {
	str, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return Db.Set(key, str, 0).Err()
}

// DbGet : Convenience function to avoid manual JSON decoding
func DbGet(key string) (out interface{}) {
	str := Db.Get(key).Val()
	json.Unmarshal([]byte(str), out)
	return
}

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
