package model

import "gorm.io/gorm"

// Point 任务点数
type Point struct {
	gorm.Model
	Num       int    `gorm:"column:num"`
	TaskID    uint   `gorm:"column:task_id;index:idx_task_id;index:idx_user_id_task_id,priority:2"`
	TaskTitle string `gorm:"column:task_title"`
	UserID    uint   `gorm:"column:user_id;index:idx_user_id_task_id,priority:1"`
}
