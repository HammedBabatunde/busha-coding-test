package redis

import (
	"context"
	"errors"
	"time"

	redisDB "github.com/emekarr/coding-test-busha/db/redis"
	"github.com/emekarr/coding-test-busha/logger"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

var (
	RedisRepo RedisRepository
)

type RedisRepository struct {
	Clinet *redis.Client
}

func generateContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 15*time.Second)
}

func (redisRepo *RedisRepository) CreateEntry(key string, payload interface{}, ttl time.Duration) (bool, error) {
	c, cancel := generateContext()

	defer func() {
		cancel()
	}()

	_, err := redisRepo.Clinet.Set(c, key, payload, ttl).Result()

	if err != nil {
		logger.Error(errors.New("could not write to redis. set"), zap.Error(err))
		return false, err
	}

	return true, nil
}

func (redisRepo *RedisRepository) FindOne(key string) (*string, error) {
	c, cancel := generateContext()

	defer func() {
		cancel()
	}()

	result, err := redisRepo.Clinet.Get(c, key).Result()

	if err != nil {
		if err.Error() == "redis: nil" {
			return nil, nil
		}
		logger.Error(errors.New("could not read from redis. get"), zap.Error(err))
		return nil, err
	}
	return &result, nil
}

func (redisRepo *RedisRepository) DeleteOne(key string) (bool, error) {
	c, cancel := generateContext()

	defer func() {
		cancel()
	}()

	result, err := redisRepo.Clinet.Del(c, key).Result()

	if err != nil {
		logger.Error(errors.New("could not delete from redis. del"), zap.Error(err))
		return false, err
	}
	if int(result) != 1 {
		return false, nil
	}
	return true, nil
}

func (redisRepo *RedisRepository) CreateInSet(key string, score float64, member interface{}) (bool, error) {
	c, cancel := generateContext()

	defer func() {
		cancel()
	}()

	result := redisRepo.Clinet.ZAdd(c, key, &redis.Z{
		Score: score, Member: member,
	})
	if result.Err() != nil {
		logger.Error(errors.New("could not add to set in redis. zadd"), zap.Error(result.Err()))
		return false, result.Err()
	}

	return result != nil, nil
}

func (redisRepo *RedisRepository) FindSet(key string) (*[]string, error) {
	c, cancel := generateContext()

	defer func() {
		cancel()
	}()

	result := redisRepo.Clinet.ZRange(c, key, 0, -1)
	if result.Err() != nil {
		logger.Error(errors.New("could not read from set in redis. zrange"), zap.Error(result.Err()))
		return nil, result.Err()
	}
	if result == nil {
		return nil, nil
	}
	val := result.Val()
	return &val, nil
}

func SetUpRedisRepo() {
	RedisRepo = RedisRepository{Clinet: redisDB.Client}
	logger.Info("redis repository initialisation complete")
}
