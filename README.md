# Redis hook foor Logrus

Saves logrus events into [https://redis.io](Redis) [https://redis.io/topics/data-types-intro#redis-lists](List).

## Example


```go
package main

import (
	"github.com/Ajnasz/logrus-redis"
	"github.com/Sirupsen/logrus"
	"github.com/go-redis/redis"
	"log"
)

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{})
	hook := logrusredis.NewLogrusRedis(redisClient, "redis_logs")
	logrus.AddHook(hook)
}

func main() {
	logrus.Info("Main started")
	logrus.WithFields(logrus.Fields{
		"foo": "bar",
		"baz": "qux",
	}).Warn("Hopp")

	res := redisClient.LRange("redis_logs", 0, 100)

	err := res.Err()

	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.Result())
}
```
