package models

import (
	"time"

	"github.com/emekarr/coding-test-busha/utils"
	"github.com/emekarr/coding-test-busha/validator"
)

type Comment struct {
	ID        string    `gorm:"primarykey" validate:"required,uuid" json:"id" `
	CreatedAt time.Time `validate:"required" json:"createdAt"`
	UpdatedAt time.Time `validate:"required" json:"updatedAt"`

	Name        string `gorm:"index" validate:"required,ascii" json:"name,omitempty"`
	Comment     string `validate:"required,lt=501" json:"comment"`
	CommenterIP string `validate:"required,ip" json:"commenterIP"`
}

func (c *Comment) InitFields() {
	if c.ID == "" {
		c.ID = utils.GenerateID()
	}
	if c.CreatedAt.IsZero() {
		currentTime := time.Now().UTC()
		c.CreatedAt = currentTime
		c.UpdatedAt = currentTime
	}
}

func (c *Comment) Validate() *[]string {
	return validator.Validator.Validate(c)
}
