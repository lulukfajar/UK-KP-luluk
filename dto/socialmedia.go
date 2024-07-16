package dto

import (
	"time"
)

type SocialMediaCreateDTO struct {
	Name           string `json:"name" binding:"required"`
	SocialMediaURL string `json:"social_media_url" binding:"required"`
	UserID         uint   `json:"user_id" binding:"required"`
}

type SocialMediaUpdateDTO struct {
	Name           string `json:"name" binding:"required"`
	SocialMediaURL string `json:"social_media_url" binding:"required"`
}

type SocialMediaResponseDTO struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaURL string    `json:"social_media_url"`
	UserID         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
