package redis

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

/*Distributed Locks Manager*/

// Lock 加锁，expire: N秒后锁失效，允许其他客户端竞争
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

// Unlock 释放锁
func Unlock(conn redis.Conn, key string) bool {
	var _, err = conn.Do("DEL", key)
	if err != nil {
		return false
	}

	return true
}

// TryLock 尝试加锁
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
