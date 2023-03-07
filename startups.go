package main

import (
	db "github.com/emekarr/coding-test-busha/db"
	"github.com/emekarr/coding-test-busha/logger"
	"github.com/emekarr/coding-test-busha/migrations"
	redisRepo "github.com/emekarr/coding-test-busha/repository/redis"
)

func StartServices() {
	// initialise logger
	logger.InitializeLogger()
	// connect to database
	db.ConnectToDB()
	// set up redis repo
	redisRepo.SetUpRedisRepo()
	// run migrations
	migrations.RunMigrations()
}

func CleanUp() {
	logger.Info("all services cleaned up")
}
