package gorm

import (
	"fmt"
	"os"

	"github.com/emekarr/coding-test-busha/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToProgres() *gorm.DB {
	dsn := os.Getenv("POSTGRESSQL_CONNECTION_STRING")
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// check if db exists
	stmt := fmt.Sprintf("SELECT * FROM pg_database WHERE datname = '%s';", os.Getenv("POSTGRESSQL_DB_NAME"))
	rs := db.Raw(stmt)
	if rs.Error != nil {
		logger.Error(rs.Error)
		return nil
	}
	var rec = make(map[string]interface{})
	if rs.Find(rec); len(rec) == 0 {
		stmt := fmt.Sprintf("CREATE DATABASE %s;", os.Getenv("POSTGRESSQL_DB_NAME"))
		if rs := db.Exec(stmt); rs.Error != nil {
			logger.Error(rs.Error)
			return nil
		}

		// close db connection
		sql, err := db.DB()
		defer func() {
			sql.Close()
		}()
		if err != nil {
			logger.Error(rs.Error)
			return nil
		}
	}
	return db
}
