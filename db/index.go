package db

import (
	dbGorm "github.com/emekarr/coding-test-busha/db/gorm"
	"gorm.io/gorm"
)

var GormDB *gorm.DB

func ConnectToDB() {
	GormDB = dbGorm.ConnectToProgres()
}

func Migrate(payload ...interface{}) error {
	if GormDB == nil {
		ConnectToDB()
	}
	return dbGorm.PostgresMigrate(GormDB, payload...)
}
