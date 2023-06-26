package redislock

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type RedisLock struct {
	client     *redis.Redis
	key        string
	expiration int
}

func NewLock(client *redis.Redis, key string, expiration int) *RedisLock {
	return &RedisLock{
		client:     client,
		key:        key,
		expiration: expiration,
	}
}

func (r *RedisLock) Lock(id string) (bool, error) {
	return r.client.SetnxEx(r.key, id, r.expiration)
	// return r.client.SetNX(context.TODO(), r.key, id, r.expiration).Result()
}

const unLockScript = `
if (redis.call("get", KEYS[1] == KEYS[2]) then
	redis.call("del", KEYS[1]
    return true
end
return false
`

func (r *RedisLock) UnLock(id string) error {
	_, err := r.client.EvalCtx(context.TODO(), unLockScript, []string{r.key, id})
	if err != nil && err != redis.Nil {
		return err
	}
	return nil
}
