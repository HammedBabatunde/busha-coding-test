package main

import (
	db "github.com/emekarr/coding-test-busha/db"
	"github.com/emekarr/coding-test-busha/logger"
	"github.com/emekarr/coding-test-busha/migrations"
)

func StartServices() {
	// initialise logger
	logger.InitializeLogger()
	// connect to database
	db.ConnectToDB()
	// run migrations
	migrations.RunMigrations()
}

func CleanUp() {
	logger.Info("all services cleaned up")
}
