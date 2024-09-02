package models

import "github.com/google/uuid"

// the user model
type User struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email" binding:"required"`
}
type UserRequestModel struct {
	Email string `json:"email" binding:"required"`
}
