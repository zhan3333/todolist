package service

import (
	"gorm.io/gorm"

	"go-framework/internal/model"
)

type Task struct {
	DB *gorm.DB
}

func NewTask() *Task {
	return &Task{}
}

func (t *Task) Create(m *model.Task) error {
	return t.DB.Create(m).Error
}
