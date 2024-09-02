package database

import (
	"moody/moody-api/src/models"
	"github.com/google/uuid"
)

func InsertCategory(params models.CategoryRequestModel) (bool, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return false, err
	}
	db, err := Database()
	if err != nil {
		return false, err
	}
	defer db.Close()
	// Insert into database
	executableStatement, err := db.Prepare(catDbInsertQuery())
	if err != nil {
		return false, err
	}
	defer executableStatement.Close()
	_, err = executableStatement.Exec(id, params.Name, params.Color)
	if err != nil {
		return false, err
	}
	return true, nil
}
