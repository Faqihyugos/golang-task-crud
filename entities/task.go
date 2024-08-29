package entities

import "time"

type Task struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	TaskName        string    `json:"task_name"`
	TaskDescription string    `json:"task_description"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
