package mytool

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func Get(key string) (value string) {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Print(err)
	}
	defer c.Close()

	v, err := redis.String(c.Do("GET", key))
	if err != nil {
		fmt.Print(err)
	}
	return v
}

func Set(key string, value string) (ok string) {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Print(err)
	}
	defer c.Close()
	//过期
	v, err := redis.String(c.Do("SET", key, value, "EX", "5"))
	if err != nil {
		fmt.Print(err)
	}
	return v
}
