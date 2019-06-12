package services

import (
	"encoding/json"
	"errors"
	"os"
	"reflect"

	"github.com/go-redis/redis"
	"k8s.io/klog"
)

// Db : Can be used to interact with redis
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

// DbZRevRange : Get given keys
func DbZRevRange(key string, from, to int64, sample interface{}) ([]interface{}, error) {
	elems, err := Db.ZRevRange(key, from, to).Result()
	if err != nil {
		return nil, err
	}

	out := make([]interface{}, 0)
	for _, elem := range elems {
		newElem := reflect.New(reflect.TypeOf(sample)).Interface()
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
	for _, elem := range elems {
		newElem := reflect.New(reflect.TypeOf(sample)).Interface()
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
	if err != nil {
		return err
	}
	if pong != "PONG" {
		return errors.New("failed to connect to redis database")
	}

	klog.Info("etablished connection with redis database")
	return nil
}
