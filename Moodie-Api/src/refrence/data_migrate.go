package refrence

import (
	"encoding/csv"
	"io"
	"moody/moody-api/src/database"
	"moody/moody-api/src/literals"
	"moody/moody-api/src/models"
	"os"

	"github.com/google/uuid"
)

// this is to tranfer data from csv to the data base
// remove duplicates
// migrate data
func MigrateData() {
	literals.BlueLog("Migrating Data To Postgress")
	// loadCsvToDiskAndMigrate()
	streamMigrationToPGSQL()
}

// this is for very small data , for production it is recomended to make a stream to plug it in
func loadCsvToDiskAndMigrate() (bool, error) {
	// OPEN THE FILE
	file, err := os.Open("src\\refrence\\data.csv")
	if err != nil {
		literals.RedLog(err.Error())
	}
	defer file.Close()
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		literals.RedLog(err.Error())
	}
	for _, x := range data {
		name := x[0]
		description := x[1]
		catId := x[2]
		emoji := x[3]
		literals.GreenLog(name + "|" + description + "|" + catId + "|" + emoji + "|")

	}
	return true, nil

}

// this function will call a db instance and also provide rollback options
func streamMigrationToPGSQL() (bool, error) {
	db, err := database.Database()
	if err != nil {
		literals.RedLog(err.Error())
		return false, err
	}
	tx, err := database.TxBegin(db)
	if err != nil {
		literals.RedLog(err.Error())
		return false, err
	}
	file, err := os.Open("src\\refrence\\data.csv")
	if err != nil {
		literals.RedLog(err.Error())
	}
	defer file.Close()
	// create a instance of the reader
	reader := csv.NewReader(file)
	for {

		record, err := reader.Read()
		if err != nil {
			// to check if we reach the eof
			if err == io.EOF {
				break
			}
			literals.RedLog(err.Error() + "EOF")
			return false, err
		}
		name := record[0]
		description := record[1]
		categoryId := record[2]
		image := record[3]

		// convert the string (catId)to uuid
		uuid, err := uuid.Parse(categoryId)
		if err != nil {
			literals.RedLog(err.Error() + "-object value-" + categoryId + "-item-name" + name)
			if rbErr := database.TxRollback(tx); rbErr != nil {
				literals.RedLog(err.Error() + "-Roll back failed Criticle State -rberr-1")
				return false, err
			}
		}
		// insert it into the postgress db

		moodReqModel := models.MoodRequestModel{
			Name:        name,
			Description: description,
			CategoryID:  uuid,
			Emoji:       image,
		}
		literals.PurpleLog(">Insert-" + name)
		if err = database.InsertSingleMoodWithTx(tx, moodReqModel); err != nil {
			literals.RedLog(err.Error())
			if rbErr := database.TxRollback(tx); rbErr != nil {
				literals.RedLog(err.Error() + "-Roll back failed Criticle State rberr-2")
				return false, err
			}
			return false, err
			// rollback to the thing mate
		}

	}
	if err := tx.Commit(); err != nil {
		if rbErr := database.TxRollback(tx); rbErr != nil {
			// this is a very criticle error thie means rollback can not be made
			literals.RedLog(err.Error() + "-Roll back failed Criticle State -rberr-3")
			return false, err
		}
		return false, err

	}
	return true, nil
}
