package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

var _host = "127.0.0.1"
var _port = 6379
var _auth = ""

var _max_idle = 20
var _max_active = 50
var _idle_timeout = 60

var Pool *redis.Pool = nil

func init() {
	var server = fmt.Sprintf("%s:%d", _host, _port)
	Pool = NewPool(server, _auth, _max_idle, _max_idle, _idle_timeout)
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

/*
 expire: N秒后锁失效，允许其他客户端竞争
*/
func Lock(conn redis.Conn, key string, expire int) bool {
	var now int64 = time.Now().Unix()

	r1, err := conn.Do("SETNX", key, now+int64(expire))
	if err != nil {
		return false
	}

	v1, err := redis.Int(r1, err)
	if err != nil {
		return false
	}

	if v1 == 1 {
		return true
	}

	/*此时key存在，查看对应的值*/
	r, err := conn.Do("GET", key)
	if err != nil {
		return false
	}

	v2, err := redis.Int64(r, err)
	if err != nil {
		return false
	}

	if now < v2 {
		/*值未过期，放弃锁*/
		return false
	} else {
		/*获取旧值，设置新值*/
		r, err := conn.Do("GETSET", key, now+int64(expire))
		if err != nil {
			return false
		}

		v3, err := redis.Int64(r, err)
		if err != nil {
			return false
		}

		if now >= v3 {
			return true
		}
		/*
		   else情况：表示其他redis客户端抢先一步设置成功，此时放弃锁
		   return false
		*/
	}

	return false
}

/*
 释放锁
*/
func Unlock(conn redis.Conn, key string) bool {
	var _, err = conn.Do("DEL", key)
	if err != nil {
		return false
	}

	return true
}

/*
 尝试加锁
*/
func TryLock(conn redis.Conn, key string, expire int, timeout int) bool {
	var b = Lock(conn, key, expire)
	if b {
		return b
	}

	if timeout <= 0 {
		return false
	}

	var ticker = time.NewTicker(time.Duration(timeout) * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-time.After(100 * time.Millisecond):
			if Lock(conn, key, expire) {
				/*成功lock后返回，否则一直持续到超时*/
				return true
			}
		}

		select {
		case <-ticker.C:
			return false
		default:
			//DO NOTHING
		}
	}
}

var WithDo = func(cmd string, args ...interface{}) (interface{}, error) {
	var conn = Pool.Get()
	if conn.Err() != nil {
		return nil, conn.Err()
	}
	defer func() { var _ = conn.Close() }()
	return conn.Do(cmd, args)
}
