package models

import "github.com/google/uuid"

type MoodModel struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"  binding:"required"`
	Description string    `json:"description"  binding:"required"`
	CategoryID  uuid.UUID `json:"categoryid"  binding:"required"`
	Emoji       string    `json:"emoji" binding:"required"`
}
type MoodRequestModel struct {
	Name        string    `json:"name"  binding:"required"`
	Description string    `json:"description"  binding:"required"`
	CategoryID  uuid.UUID `json:"categoryid"  binding:"required"`
	Emoji       string    `json:"emoji" binding:"required"`
}
