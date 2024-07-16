package models

import "time"

type SocialMedia struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `gorm:"type:varchar(100);not null"`
	SocialMediaURL string `gorm:"type:text;not null"`
	UserID         uint
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}
