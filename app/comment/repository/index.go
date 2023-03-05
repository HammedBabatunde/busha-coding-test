package repository

import (
	"github.com/emekarr/coding-test-busha/app/comment/models"
	"github.com/emekarr/coding-test-busha/db"
	gormRepo "github.com/emekarr/coding-test-busha/repository/gorm"
)

var CommentRepository = gormRepo.GormRepository[models.Comment]{Gorm: db.GormDB}
