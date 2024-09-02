package database

import (
	"fmt"
	"log"
	"moody/moody-api/src/internals"
	"moody/moody-api/src/literals"
)

// !!! always use prepared statements in the database

// all of the databse migrations are here
func MigrateAll() {
	// migrate all the tables
	literals.RedLog("Migrations Innit , Migration Name : Alpha ")
	migration_alpha()
}
func migration_alpha() bool {
	// database di
	db, err := Database()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// user table migration !
	// prepare the statement
	statement, err := db.Prepare(userTable())
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	// execute the migration
	// migrations are only done while an update is pushed , more logs are used to get the developers
	// updtodate and for tracking the errors that may happen during migrations
	literals.BlueLog("UserTable Execution Innit")
	_, err = statement.Exec()
	if err != nil {
		log.Fatal(err)
		return false
	}
	literals.RedLog("USER MIGRATED")
	// cats migration start !!
	literals.BlueLog("CAT MIGRATION INIT")
	statement1, err := db.Prepare(catDbMigrartion())
	if err != nil {
		log.Fatal(err)
		literals.RedLog("Fail At Cat1 -" + err.Error())
	}
	_, err = statement1.Exec()
	defer statement1.Close()
	if err != nil {
		log.Fatal(err)
		literals.RedLog("Fail At Cat2 -" + err.Error())
	}
	literals.RedLog("Cat Migrated")

	// mood migration !!
	literals.BlueLog("MOOD MIGRATION INIT")
	statement2, err := db.Prepare(moodDbMigrationQuery())

	if err != nil {
		log.Fatal(err)
		literals.RedLog("Fail At Mood1 -" + err.Error())
	}
	_, err = statement2.Exec()
	if err != nil {
		log.Fatal(err)
		literals.RedLog("Fail At Mood2 -" + err.Error())
	}
	literals.RedLog("MOOD MIGRATED")
	defer statement2.Close()
	// this is atr the end of the migrations , if all is a success then this is logged in the end
	literals.GreenLog("All Migrations Ended")
	return true
}
func BetterMigration() {
	// this is a better way to do migrations , it is more safe and more efficient
	migrationCounter := 0
	isUserMood, err := CreateUserMoodMigration()
	if err != nil {
		internals.RedLog(err.Error())
	}
	if isUserMood {
		migrationCounter++
		internals.GreenLog("success , migrations done =" + fmt.Sprintf("%d", migrationCounter))
	}
}
