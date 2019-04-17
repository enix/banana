package services

import (
	"encoding/json"
	"errors"
	"os"

	"enix.io/banana/src/logger"
	"enix.io/banana/src/models"
	"github.com/go-redis/redis"
)

// Db : Use this API to interact with redis
var Db *redis.Client

// DbGet : Convenience function to avoir JSON unmarshalling
func DbGet(key string, out interface{}) error {
	result := Db.Get(key)

	err := result.Err()
	if err != nil {
		return err
	}

	bytes, err := result.Bytes()
	if err != nil {
		return err
	}

	json.Unmarshal(bytes, &out)
	return err
}

// DbSet : Convenience function to avoir JSON marshalling
func DbSet(key string, value interface{}) error {
	str, err := json.Marshal(value)
	if err != nil {
		return err
	}

	result := Db.Set(key, []byte(str), 0)
	return result.Err()
}

// OpenDatabaseConnection : Connect to redis databae
// Calls such as DbGet and DbSet will crash if called before this
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

	var agents []models.Agent
	err = DbGet("agents", agents)
	if err == redis.Nil {
		DbSet("agents", make([]models.Agent, 0))
	}

	logger.Log("etablished connection with redis database")
	return nil
}
