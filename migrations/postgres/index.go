package postgres

import (
	"github.com/emekarr/coding-test-busha/app/comment/models"
	"github.com/emekarr/coding-test-busha/db"
)

func Migrate() {
	db.Migrate(&models.Comment{})
}
