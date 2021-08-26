// @Title  
// @Description  
// @Author  Wangwengang  2021/8/24 下午8:21
// @Update  Wangwengang  2021/8/24 下午8:21
package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/wwengg/arsenal/config"
)

var (
	Redis *RedisV8
)

type RedisV8 struct {
	client redis.UniversalClient
}

func Setup(){
	var err error
	Redis,err = NewRedis()
	if err != nil {
		panic(err)
	}
}


func NewRedis()(*RedisV8, error){
	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:              config.ConfigHub.Redis.Addr,
		DB:                 config.ConfigHub.Redis.Db,
		Password: 			config.ConfigHub.Redis.Password,
		MasterName:         config.ConfigHub.Redis.MasterName,
	})
	r := &RedisV8{
		client: client,
	}
	err := r.connet()
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (r *RedisV8) String() string{
	return "redis"
}

func (r *RedisV8) connet() error{
	var err error
	_, err = r.client.Ping(context.Background()).Result()
	return err
}


func (r *RedisV8) Get(key string) (string, error){
	return r.client.Get(context.Background(),key).Result()
}


// Set value with key and expire time
func (r *RedisV8) Set(key string, val interface{}, expire int) error{
	return r.client.Set(context.Background(),key, val, time.Duration(expire)*time.Second).Err()
}


// Del delete key in redis
func (r *RedisV8) Del(key string) error {
	return r.client.Del(context.Background(),key).Err()
}



// HashGet from key
func (r *RedisV8) HashGet(hk, key string) (string, error) {
	return r.client.HGet(context.Background(),hk, key).Result()
}

func (r *RedisV8)HashSet(hk,key string,val interface{}) error{
	return r.client.HSet(context.Background(),hk,key,val).Err()
}

// HashDel delete key in specify redis's hashtable
func (r *RedisV8) HashDel(hk, key string) error {
	return r.client.HDel(context.Background(),hk, key).Err()
}

// Increase
func (r *RedisV8) Increase(key string) error {
	return r.client.Incr(context.Background(),key).Err()
}

func (r *RedisV8) Decrease(key string) error {
	return r.client.Decr(context.Background(),key).Err()
}


// Set ttl
func (r *RedisV8) Expire(key string, dur time.Duration) error {
	return r.client.Expire(context.Background(),key, dur).Err()
}

// GetClient 暴露原生client
func (r *RedisV8) GetClient() redis.UniversalClient {
	return r.client
}