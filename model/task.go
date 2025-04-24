package model

import (
	"taskMangementService/enums"
	"time"
)


type Task struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    Title       string    `json:"title" gorm:"not null"`
    Description string    `json:"description" gorm:"not null"`
    Status      enums.TaskStatus    `json:"status" gorm:"not null"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}