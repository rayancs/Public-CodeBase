package service

import (
	"moody/moody-api/src/database"
	"moody/moody-api/src/models"
)

func (app *database.App) RegisterUserMood() *models.ResponseModel {
		app.InsertUsersMood()
}
