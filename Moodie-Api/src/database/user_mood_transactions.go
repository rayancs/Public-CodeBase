package database

import (
	"moody/moody-api/src/models"

	"github.com/google/uuid"
)

// CREATES A USER TABLE
func CreateUserMoodMigration() (bool, error) {
	err := ExecuteAuto(CreateUserMoodQuery())
	if err != nil {
		return false, err
	}
	return true, nil
}
func (db *App) InsertUsersMood(params *models.UserMoodRequest) (bool, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return false, err
	}
	parsedUserID, err := IsValidUUID(params.UserID)
	if err != nil {
		return false, err
	}
	parsedMoodID, err := IsValidUUID(params.MoodID)
	if err != nil {
		return false, err
	}
	// ... rest of the code
	err = db.ExecuteArgs(InsertUserMoodQuery(), id, parsedUserID, params.Note, parsedMoodID)
	if err != nil {
		return false, err
	}
	return true, nil

}
