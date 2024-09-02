// all user queries
package database

import (
	"database/sql"
	"moody/moody-api/src/models"
	"github.com/google/uuid"
)

func InsertUserModel(params models.UserRequestModel) (bool, error) {
	// Insert user model into database
	// questionalbe functionality for the db layer , as many orms produce guuid naturaly , we are making it in the querry
	id, err := uuid.NewUUID()
	if err != nil {
		// Insert user model into database
		return false, err
	}
	// the transaction
	db, err := Database()
	if err != nil {
		return false, err
	}
	defer db.Close()
	// the query
	// find the user first

	statement, err := db.Prepare(userInsertQuery())
	if err != nil {
		return false, err
	}
	_, err = statement.Exec(id, params.Email)
	if err != nil {
		return false, err
	}
	return true, nil

}
func FindUserByEmail(email string) (*models.User, error) {
	// the transaction
	db, err := Database()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	// the query
	// this is how sql get works in go  make a the query using db.QueryRow , provide query and params
	// create a user object that binds to the db ,
	// use db scan to create bindings
	row := db.QueryRow(findUserByIdQuery(), email)
	// create the object that is supposed to provide the binding
	var userModel models.User
	// bind the row to the object
	// if you switch the values aka email then id , the following would happen
	/*
		|id|email|...|
		|1 |a   |...|
		0,0 bind to params 1
		0,1 bind to params 2
		thus the schema  should be kept in mind
		row.Scan (v1,v2)
	*/
	err = row.Scan(&userModel.ID, &userModel.Email)
	// check for now rows first , if ur checking for general errors then sq.NoRows is also included in the error , and the function will return an err , that will terminate the
	// signup process , check for now rows helps to return a null user that helps to continue the signup process and keeps the func going
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &userModel, nil

}
