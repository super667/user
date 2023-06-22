package redislock

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisLock struct {
	client     *redis.Client
	key        string
	expiration time.Duration
}

func NewLock(client *redis.Client, key string, expiration time.Duration) *RedisLock {
	return &RedisLock{
		client:     client,
		key:        key,
		expiration: expiration,
	}
}

func (r *RedisLock) Lock(id string) (bool, error) {
	return r.client.SetNX(context.TODO(), r.key, id, r.expiration).Result()
}

const unLockScript = `
if (redis.call("get", KEYS[1] == KEYS[2]) then
	redis.call("del", KEYS[1]
    return true
end
return false
`

func (r *RedisLock) UnLock(id string) error {
	_, err := r.client.Eval(context.TODO(), unLockScript, []string{r.key, id}).Result()
	if err != nil && err != redis.Nil {
		return err
	}
	return nil
}
