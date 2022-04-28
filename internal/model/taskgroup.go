package model

import "gorm.io/gorm"

// TaskGroup 任务分组
type TaskGroup struct {
	gorm.Model
	Name   string
	UserID uint `json:"user_id" gorm:"index"`
}
