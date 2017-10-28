package dataface

import (
	"log"
	"strconv"

	"github.com/go-redis/redis"
)

//RedisDB is the database structure type for Redis, which
//complies with the database{} interface
type RedisDB struct {
	Address string `json:"address"`
	client  *redis.Client
}

//NewRedisDB creates a new redis client
func NewRedisDB(host string, port int, password string) (*RedisDB, error) {
	var red RedisDB
	red.client = redis.NewClient(&redis.Options{
		Addr:     host + ":" + strconv.Itoa(port),
		Password: password, // no password
		DB:       0,        // use url DB
	})
	t, err := red.client.Ping().Result()

	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(t)
	return &red, nil
}

//Put adds the URLdata json to the key string in redis
func (r RedisDB) Put(key string, value []byte) error {
	return r.client.Set(key, value, 0).Err()
}

//Get uses the key to return the URL translation
func (r RedisDB) Get(key string) ([]byte, error) {
	return r.client.Get(key).Bytes()
}

// Close the session
func (r RedisDB) Close() error {
	return r.client.Close()
}
