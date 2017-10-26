package redis

import (
	"errors"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

var _host = "127.0.0.1"
var _port = 6379
var _auth = ""

var _max_idle = 20
var _max_active = 50
var _idle_timeout = 60

var Pools map[string]*redis.Pool = nil

func init() {
	Pools = make(map[string]*redis.Pool)
	var server = fmt.Sprintf("%s:%d", _host, _port)
	Pools[""] = NewPool(server, _auth, _max_idle, _max_idle, _idle_timeout)
}

var NewPool = func(server string, auth string, max_idle, max_active, idle_timeout int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     max_idle,
		MaxActive:   max_active,
		IdleTimeout: time.Duration(idle_timeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			var c, err_c = redis.Dial("tcp", server)
			if err_c != nil {
				return nil, err_c
			}

			if auth != "" {
				var _, err_auth = c.Do("AUTH", auth)
				if err_auth != nil {
					var _ = c.Close()
					return nil, err_auth
				}
			}

			return c, err_c
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}

			var _, err = c.Do("PING")
			return err
		},
	}
}

var WithDo = func(pn string, cmd string, args ...interface{}) (interface{}, error) {
	var pool, ok = Pools[pn]
	if !ok {
		return nil, errors.New(fmt.Sprintf("pool<%s> not found", pn))
	}

	var conn = pool.Get()
	if conn.Err() != nil {
		return nil, conn.Err()
	}
	defer func() { var _ = conn.Close() }()

	return conn.Do(cmd, args)
}
