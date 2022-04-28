package model

import "gorm.io/gorm"

// Task 任务表
type Task struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	// created, finished, canceled
	Status string `json:"status" gorm:"index:idx_group_id_status,priority:2"`
	// none，day，workday，weekend，weekday, month, year
	RepeatType string `json:"repeat_type"`
	// 完成任务能得到的积分
	Point   int `json:"point"`
	GroupID int `json:"group_id" gorm:"index:idx_group_id_status,priority:1"`
}
