package service

import (
	"errors"

	"moody/moody-api/src/database"
	"moody/moody-api/src/models"
)

func CreateNewUser(params models.UserRequestModel) (bool, error) {
	// first check if the old user exists
	// verify the data
	res := ValidateEmailStr(params.Email)
	if !res.Valid {
		return false, errors.New(res.Errors)
	}
	user, err := database.FindUserByEmail(params.Email)
	if err != nil {
		return false, err
	}
	// if the user exists return false and the error message
	if user != nil {
		return true, nil

	}
	// create a new user
	_, err = database.InsertUserModel(params)
	if err != nil {
		return false, err
	}
	return true, nil
}
