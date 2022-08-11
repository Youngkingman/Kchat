package dbutil

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Youngkingman/Kchat/kchat/pkg/setting"
	redis "github.com/go-redis/redis"
)

type RedisCli struct {
	Cli *redis.Client
}

var redisOnce sync.Once

//ErrKeyNotExists not exists
var ErrKeyNotExists = redis.Nil

func NewRedisCli(config *setting.RedisSettingS) (*RedisCli, error) {
	var client *redis.Client
	redisOnce.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr:       fmt.Sprintf("%s", config.Host),
			Password:   config.Auth,
			DB:         config.Db,
			MaxRetries: 2,
			PoolSize:   config.PoolSize,
			// 不想配了
			ReadTimeout:  time.Millisecond * time.Duration(500),
			WriteTimeout: time.Millisecond * time.Duration(500),
			IdleTimeout:  time.Second * time.Duration(60),
		})
	})
	_, err := client.Ping().Result()
	_client := &RedisCli{client}
	return _client, err
}

//Get get redis value
func (rc *RedisCli) Get(key string) (string, error) {
	return rc.Cli.Get(key).Result()
}

//Set set redis key value
func (rc *RedisCli) Set(key string, val string, expiration time.Duration) error {
	err := rc.Cli.Set(key, val, expiration).Err()
	if err != nil {
		log.Printf("redis set key: %s val: %s fail: %s", key, val, err)
	}
	return err
}

//Del delete the key
func (rc *RedisCli) Del(key string) error {
	return rc.Cli.Del(key).Err()
}

//TTL change the Ttl
func (rc *RedisCli) TTL(key string) (time.Duration, error) {
	r := rc.Cli.TTL(key)
	return r.Val(), r.Err()
}

//Client return the raw redis client
func (rc *RedisCli) Client() *redis.Client {
	return rc.Cli
}

//RPush RPush
func (rc *RedisCli) RPush(key string, value string) error {

	return rc.Cli.RPush(key, value).Err()
}

//LPush LPush
func (rc *RedisCli) LPush(key string, value string) error {
	return rc.Cli.LPush(key, value).Err()
}

//RPop RPop
func (rc *RedisCli) RPop(key string) (string, error) {
	r := rc.Cli.RPop(key)
	return r.Val(), r.Err()
}

//LPop LPop
func (rc *RedisCli) LPop(key string) (string, error) {
	r := rc.Cli.LPop(key)
	return r.Val(), r.Err()
}

//LLen LLen
func (rc *RedisCli) LLen(key string) (int64, error) {
	r := rc.Cli.LLen(key)
	return r.Val(), r.Err()
}

//LRem LRem
func (rc *RedisCli) LRem(key string, count int64, value string) error {
	return rc.Cli.LRem(key, count, value).Err()
}

//Expire set redis key expire
func (rc *RedisCli) Expire(key string, expiration time.Duration) error {
	err := rc.Cli.Expire(key, expiration).Err()
	return err
}
