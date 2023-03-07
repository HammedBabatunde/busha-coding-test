package db

import (
	dbGorm "github.com/emekarr/coding-test-busha/db/gorm"
	dbRedis "github.com/emekarr/coding-test-busha/db/redis"
	"gorm.io/gorm"
)

var GormDB *gorm.DB

func ConnectToDB() {
	GormDB = dbGorm.ConnectToProgres()
	dbRedis.ConnectRedis()
}

func Migrate(payload ...interface{}) error {
	if GormDB == nil {
		ConnectToDB()
	}
	return dbGorm.PostgresMigrate(GormDB, payload...)
}
