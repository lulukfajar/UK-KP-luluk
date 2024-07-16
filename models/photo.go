package models

import "time"

type Photo struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Caption   string
	PhotoURL  string `gorm:"not null"`
	UserID    uint
	User      User
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
