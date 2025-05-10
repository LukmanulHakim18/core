package redis

import (
	"context"
	"fmt"
	"strings"
	"time"

	redis "github.com/go-redis/redis/v8"
	apmgoredis "go.elastic.co/apm/module/apmgoredisv8/v2"
)

type EventKey string

const (
	EventSet     = "set"
	EventExpired = "expired"
	EventDel     = "del"
)

var (
	RedisNil  = redis.Nil
	ErrorLock = fmt.Errorf("Unable to process request because it is locked")
)

type ClientRedis struct {
	client *redis.Client
}

func NewRedis(address string, port int, password string, db int) ClientRedis {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", address, port),
		Password: password,
		DB:       db,
	})
	client.AddHook(apmgoredis.NewHook())
	client.AddHook(newMetricHook())
	var res = ClientRedis{
		client: client,
	}
	return res
}

func (r *ClientRedis) Client() *redis.Client {
	return r.client
}

func (r *ClientRedis) Ping() error {
	return r.client.Ping(context.Background()).Err()
}

// func for set Redis as key space event
// redis cant provide event to listener service.
func (r *ClientRedis) SetConfigKEA() error {
	return r.client.Do(context.Background(), "CONFIG", "SET", "notify-keyspace-events", "KEA").Err()
}

// deprecated
func (r *ClientRedis) ListenEventHappenOnKey(ctx context.Context, key string) *redis.PubSub {
	return r.client.PSubscribe(ctx, fmt.Sprintf("__keyspace*:%s", key))
}

// listening key and returning event.
// if db not set it wil listen to all database in redis
func (r *ClientRedis) ListenKey(ctx context.Context, db, key string) *redis.PubSub {
	if db == "" {
		db = "*"
	}
	return r.client.PSubscribe(ctx, fmt.Sprintf("__keyspace%s:%s", db, key))
}

func (r *ClientRedis) ListenEvent(ctx context.Context, db, event string) *redis.PubSub {
	if db == "" {
		db = "*"
	}
	return r.client.PSubscribe(ctx, fmt.Sprintf("__keyevent%s:%s", db, event))
}

// make pattern key like this CONTEXT_PROCESS:LOCK:DATA_KEY
func (r *ClientRedis) Lock(ctx context.Context, lockKey string, unlockUntil time.Duration) error {
	return r.client.SetNX(ctx, lockKey, true, unlockUntil).Err()
}

// unlock
func (r *ClientRedis) Unlock(ctx context.Context, lockKey string) error {
	return r.client.Del(ctx, lockKey).Err()
}

// Build key with : separator
func BuildKey(key ...string) string {
	return strings.Join(key, ":")
}
