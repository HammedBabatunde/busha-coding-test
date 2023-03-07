package redis

import (
	"errors"
	"os"

	"github.com/emekarr/coding-test-busha/logger"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

var (
	Client *redis.Client
)

func ConnectRedis() {
	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		logger.Error(errors.New("could not start redis"), zap.Error(err))
		panic("failed to start redis")
	}

	logger.Info("connected to redis")
	Client = redis.NewClient(opt)
}
