package database

import (
	"database/sql"
	"moody/moody-api/src/literals"
	"moody/moody-api/src/models"
	"github.com/google/uuid"
)

func InsertSingleMood(params models.MoodRequestModel) (bool, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return false, err
	}

	// database instance
	db, err := Database()
	if err != nil {
		return false, err
	}
	defer db.Close()
	preparedStmt, err := db.Prepare(moodInsertQuery())
	if err != nil {
		return false, err
	}
	_, err = preparedStmt.Exec(
		id,
		params.Name,
		params.Description,
		params.CategoryID,
		params.Emoji)
	if err != nil {
		return false, err
	}
	defer preparedStmt.Close()
	return true, nil

}

// to used if  you want rollback
func InsertSingleMoodWithRollBack(params models.MoodRequestModel) (bool, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return false, err
	}

	// database instance
	db, err := Database()
	if err != nil {
		return false, err
	}
	defer db.Close()
	tx, err := TxBegin(db)
	if err != nil {
		return false, err
	}

	preparedStmt, err := tx.Prepare(moodInsertQuery())
	if err != nil {
		// if any error occurs during the statement it will
		if rbErr := TxRollback(tx); rbErr != nil {
			return false, rbErr
		}
		literals.RedLog("TRANS FAILED , ROLLBACK")
		return false, err
	}
	_, err = preparedStmt.Exec(
		id,
		params.Name,
		params.Description,
		params.CategoryID,
		params.Emoji)
	if err != nil {
		return false, err
	}
	defer preparedStmt.Close()
	// COMMIT THE TRANSACTION
	if err := TxCommit(tx); err != nil {
		literals.RedLog("Error Dring the Commit e-EDFE73C33C" + err.Error())
	}
	return true, nil

}

// this is ment to be used in a loop as it is provided with db instance , rather than a new instance every time the function is called
/*
	***detailed explination***
	db or tx instance
	*db/tx
	*tx begin()
	|-------------------------------------
	// if err := func(*db/*tx ,params)err{
	// execute.query(...params) 			loop this section
	// };
	|-------------------------------------
	err !=nil {
	*tx.rollback()
	}
	tx.commit()
	!!! note this requires the context to be called in the funtion
*/
func InsertSingleMoodWithTx(tx *sql.Tx, params models.MoodRequestModel) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	preparedStmt, err := tx.Prepare(moodInsertQuery())
	if err != nil {
		// if any error occurs during the statement it will
		return err
	}
	_, err = preparedStmt.Exec(
		id,
		params.Name,
		params.Description,
		params.CategoryID,
		params.Emoji)
	if err != nil {
		return err
	}
	defer preparedStmt.Close()
	return nil
}
