package models

import "github.com/google/uuid"

type UserMoodModel struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
	Note   uuid.UUID `json:"note"`
	MoodID uuid.UUID `json:"mood_id"`
}
type UserMoodRequest struct {
	UserID string    `json:"user_id"`
	Note   uuid.UUID `json:"note"`
	MoodID string    `json:"mood_id"`
}
