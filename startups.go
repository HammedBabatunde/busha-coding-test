package main

import (
	db "github.com/emekarr/coding-test-busha/db"
	"github.com/emekarr/coding-test-busha/logger"
)

func StartServices() {
	logger.InitializeLogger()

	db.ConnectToDB()
}

func CleanUp() {
	logger.Info("all services cleaned up")
}
