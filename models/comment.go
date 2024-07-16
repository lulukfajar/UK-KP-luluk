package models

import "time"

type Comment struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	PhotoID   uint      `gorm:"not null"`
	Message   string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
