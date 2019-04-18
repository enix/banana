package services

import (
	"encoding/json"
	"errors"
	"os"
	"reflect"

	"enix.io/banana/src/logger"
	"github.com/go-redis/redis"
)

// Db : Use this API to interact with redis
var Db *redis.Client

// DbGet : Convenience function to avoid JSON unmarshalling
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

	json.Unmarshal(bytes, out)
	return err
}

// DbSet : Convenience function to avoid JSON marshalling
func DbSet(key string, value interface{}) error {
	str, err := json.Marshal(value)
	if err != nil {
		return err
	}

	result := Db.Set(key, []byte(str), 0)
	return result.Err()
}

// DbZAdd : Add given value to sorted set with the given score
func DbZAdd(key string, score float64, value interface{}) error {
	str, err := json.Marshal(value)
	if err != nil {
		return err
	}

	result := Db.ZAdd(key, redis.Z{
		Score:  score,
		Member: str,
	})

	return result.Err()
}

// DbZRange : Get given keys
func DbZRange(key string, from, to int64, sample interface{}) ([]interface{}, error) {
	elems, err := Db.ZRange(key, from, to).Result()
	if err != nil {
		return nil, err
	}

	out := make([]interface{}, 0)
	newElem := reflect.New(reflect.TypeOf(sample)).Interface()
	for _, elem := range elems {
		err := json.Unmarshal([]byte(elem), &newElem)
		if err != nil {
			return nil, err
		}
		out = append(out, newElem)
	}

	return out, nil
}

// DbMGet : Get given keys
func DbMGet(keys []string, sample interface{}) ([]interface{}, error) {
	elems, err := Db.MGet(keys...).Result()
	if err != nil {
		return nil, err
	}

	out := make([]interface{}, 0)
	newElem := reflect.New(reflect.TypeOf(sample)).Interface()
	for _, elem := range elems {
		err := json.Unmarshal([]byte(elem.(string)), &newElem)
		if err != nil {
			return nil, err
		}
		out = append(out, newElem)
	}

	return out, nil
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

	logger.Log("etablished connection with redis database")
	return nil
}
