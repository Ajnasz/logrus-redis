package logrusredis

import (
	"github.com/Sirupsen/logrus"
	"github.com/go-redis/redis"
	"time"
)

// LogrusRedis delivers logs to a Redis List
type LogrusRedis struct {
	client      *redis.Client
	key         string
	rangeCursor int64
	len         int64
	Expire      time.Duration
	formatter   *logrus.TextFormatter
}

// Fire adds logrus entry into redis list
func (r *LogrusRedis) Fire(entry *logrus.Entry) error {

	body, err := r.formatter.Format(entry)

	if err != nil {
		return err
	}

	err = r.client.RPush(r.key, body).Err()

	if err != nil {
		return err
	}

	err = r.client.Expire(r.key, time.Duration(1)*time.Hour).Err()

	return err
}

// Levels returns the available logging levels.
func (r *LogrusRedis) Levels() []logrus.Level {
	return logrus.AllLevels
}

// NewLogrusRedis creates LogrusRedis instance
func NewLogrusRedis(client *redis.Client, key string) *LogrusRedis {
	return &LogrusRedis{
		client:    client,
		key:       key,
		Expire:    time.Duration(1) * time.Hour,
		formatter: &logrus.TextFormatter{DisableColors: true},
	}
}
