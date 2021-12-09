package database

import (
	"strconv"

	"github.com/go-redis/redis/v7"
	"github.com/pkg/errors"
)

type RedisControl struct {
	Client *redis.Client
}

var PemgrDB, StateDB RedisControl

func InitializePemgr(addr string) {
	PemgrDB.Client = redis.NewClient(&redis.Options{
		Addr:   addr,
		Password: "",
		DB:     0,
	})
	_, err := PemgrDB.Client.Ping().Result()
	if err != nil {
		panic(err)
	}

	PemgrDB.Client.FlushAll()
}

func InitializeSonic() {
	StateDB.Client = redis.NewClient(&redis.Options{
		Network: "unix",
		Addr:   "/var/run/redis/redis.sock",
		Password: "",
		DB:     6,
	})
	_, err := StateDB.Client.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func (r *RedisControl) HGet(key string, field string) (string, error) {
	return r.Client.HGet(key, field).Result()
}

func (r *RedisControl) HGetInt64(key string, field string) (int64, error) {
	val, err := r.Client.HGet(key, field).Result()
	if err != nil {
		return 0, errors.WithStack(err)
	}

	valInt, err := strconv.ParseInt(val, 10, 64)
	return valInt, errors.WithStack(err)
}

func (r *RedisControl) HGetAll(key string) (map[string]string, error) {
	return r.Client.HGetAll(key).Result()
}

func (r *RedisControl) HSet(key string, field string, data interface{}) (error) {
	return r.Client.HSet(key, field, data).Err()
}

func (r *RedisControl) Get(key string) (string, error) {
	return r.Client.Get(key).Result()
}

func (r *RedisControl) GetInt32(key string) (int32, error) {
	val, err := r.Client.Get(key).Result()
	if err != nil {
		return 0, errors.WithStack(err)
	}

	valInt, err := strconv.ParseInt(val, 10, 32)
	return int32(valInt), errors.WithStack(err)
}

func (r *RedisControl) GetInt64(key string) (int64, error) {
	val, err := r.Client.Get(key).Result()
	if err != nil {
		return 0, errors.WithStack(err)
	}

	valInt, err := strconv.ParseInt(val, 10, 64)
	return valInt, errors.WithStack(err)
}

func (r *RedisControl) GetFloat32(key string) (float32, error) {
	val, err := r.Client.Get(key).Result()
	if err != nil {
		return 0, errors.WithStack(err)
	}

	valFloat, err := strconv.ParseFloat(val, 32)
	return float32(valFloat), errors.WithStack(err)
}

func (r *RedisControl) GetFloat64(key string) (float64, error) {
	val, err := r.Client.Get(key).Result()
	if err != nil {
		return 0, errors.WithStack(err)
	}

	valFloat, err := strconv.ParseFloat(val, 64)
	return valFloat, errors.WithStack(err)
}

func (r *RedisControl) Set(key string, val string) error {
	return errors.WithStack(r.Client.Set(key, val, 0).Err())
}

func (r *RedisControl) Del(key string) error {
	return errors.WithStack(r.Client.Del(key).Err())
}

func (r *RedisControl) Keys(pattern string) ([]string, error) {
	return r.Client.Keys(pattern).Result()
}
