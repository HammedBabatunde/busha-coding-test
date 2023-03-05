package models

import (
	"time"

	"github.com/emekarr/coding-test-busha/utils"
	"github.com/emekarr/coding-test-busha/validator"
)

type Comment struct {
	ID        string `gorm:"primarykey" validate:"required,uuid" json:"id" `
	CreatedAt string `validate:"required,datetime" json:"createdAt"`
	UpdatedAt string `validate:"required,datetime" json:"updatedAt"`

	Name        string `validate:"required,alpha" json:"name"`
	Comment     string `validate:"required,alpha,lt=501" json:"comment"`
	CommenterIP string `validate:"required,ip" json:"commenterIP"`
}

func (c *Comment) InitFields() {
	if c.ID != "" {
		c.ID = utils.GenerateID()
	}
	if c.CreatedAt != "" {
		currentTime := time.Now().UTC().String()
		c.CreatedAt = currentTime
		c.UpdatedAt = currentTime
	}
}

func (c *Comment) Validate() error {
	return validator.Validator.Validate(c)
}