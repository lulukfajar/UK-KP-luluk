package dto

import (
	"time"
)

type CommentCreateDTO struct {
	PhotoID uint   `json:"photo_id" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type CommentUpdateDTO struct {
	PhotoID uint   `json:"photo_id" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type CommentResponseDTO struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	PhotoID   uint      `json:"photo_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
