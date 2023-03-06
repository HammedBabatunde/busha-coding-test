package repository

import (
	"sync"

	"github.com/emekarr/coding-test-busha/app/comment/models"
	"github.com/emekarr/coding-test-busha/db"
	gormRepo "github.com/emekarr/coding-test-busha/repository/gorm"
)

var once = &sync.Once{}

var CommentRepository gormRepo.GormRepository[models.Comment]

func GetCommentRepository() gormRepo.GormRepository[models.Comment] {
	once.Do(func() {
		CommentRepository = gormRepo.GormRepository[models.Comment]{Gorm: db.GormDB}
	})
	return CommentRepository
}
