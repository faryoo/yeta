package cache_test

import (
	"github.com/faryoo/yeta/cache"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	opts := &cache.RedisOpts{
		Host: "127.0.0.1:6379",
	}
	redis := cache.NewRedis(opts)
	// redis.SetConn(redis.conn)
	var err error
	timeoutDuration := 1 * time.Second

	if err = redis.Set("username", "silenceper", timeoutDuration); err != nil {
		t.Error("set Error", err)
	}

	if !redis.IsExist("username") {
		t.Error("IsExist Error")
	}

	name := redis.Get("username").(string)
	if name != "silenceper" {
		t.Error("get Error")
	}

	if err = redis.Delete("username"); err != nil {
		t.Errorf("delete Error , err=%v", err)
	}
}
