package gorm

import (
	"errors"

	"github.com/emekarr/coding-test-busha/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type GormRepository[T interface{}] struct {
	Gorm *gorm.DB
}

func (gr *GormRepository[T]) errFilter(err error, filter error) error {
	if errors.Is(err, filter) {
		return nil
	}
	return err
}

func (gr *GormRepository[T]) CreateOne(payload *T) (*T, error) {
	err := gr.Gorm.Create(payload).Error
	if err != nil {
		logger.Error(errors.New("db error - could not insert record"), zap.Error(err))
		return nil, err
	}
	return payload, nil
}

// Parameter filter can be of type string or struct pointer
func (gr *GormRepository[T]) FindMany(filter interface{}) (*[]T, error) {
	payload := []T{}
	result := gr.Gorm.Where(filter).Find(&payload)
	if err := gr.errFilter(result.Error, gorm.ErrRecordNotFound); err != nil {
		logger.Error(errors.New("db error - find many search failed"), zap.Error(result.Error))
		return nil, err
	}
	return &payload, nil
}

func (gr *GormRepository[T]) CountDocs(filter interface{}) (*int64, error) {
	var count int64
	result := gr.Gorm.Where(filter).Count(&count)
	if err := gr.errFilter(result.Error, gorm.ErrRecordNotFound); err != nil {
		logger.Error(errors.New("db error - find many search failed"), zap.Error(result.Error))
		return nil, err
	}
	return &count, nil
}
