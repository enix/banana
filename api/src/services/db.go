package services

import (
	"encoding/json"
	"errors"
	"os"

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

	if str == "[]" {
		// yes go does not handle that... https://stackoverflow.com/a/33183170
		out = make([]interface{}, 0)
		return
	}

	json.Unmarshal([]byte(str), out)
	return
}

// OpenDatabaseConnection : Connect to redis databae
//													Calls such as DbGet and DbSet will crash if called before this
func OpenDatabaseConnection() error {
	Db = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWD"),
		DB:       0,
	})

	pong, err := Db.Ping().Result()
	if err != nil || pong != "PONG" {
		return errors.New("failed to connect to redis database")
	}

	agents := DbGet("agents")
	if agents == nil {
		DbSet("agents", make([]interface{}, 0))
	}

	logger.Log("etablished connection with redis database")
	return nil
}
