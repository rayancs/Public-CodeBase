package models

import "github.com/google/uuid"

type CategoryModel struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name" binding:"required"`
	Color string    `json:"color" binding:"required"`
}
type CategoryRequestModel struct {
	Name  string    `json:"name" binding:"required"`
	Color string    `json:"color" binding:"required"`
}