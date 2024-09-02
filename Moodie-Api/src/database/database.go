package database

import (
	"database/sql"
	"errors"
	"fmt"
	"moody/moody-api/src/internals"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// type App struct {
// 	DB *sql.DB
// }

func (a *App) DbInit(
	maxOpenConns int,
	idleConns int,
	connLifetime time.Duration) error {
	err := godotenv.Load()
	if err != nil {
		internals.ErrorExit("Env File Error", err)
	}
	connStr := os.Getenv("PGSQL")
	if connStr == "" {
		// just do a log and exit the application if it breaks as without presistance the app should not function
		internals.ErrorExit("No Connection String Present", fmt.Errorf("no_conn_str"))
	}
	a.DB, err = sql.Open("postgres", connStr)
	if err != nil {
		internals.ErrorExit("Database Error ", err)
	}
	a.DB.SetMaxOpenConns(maxOpenConns)
	a.DB.SetMaxIdleConns(idleConns)
	a.DB.SetConnMaxLifetime(time.Duration(connLifetime) * time.Minute) // five mins
	return nil
}
func Database() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("could not load env file")
	}
	connStr := os.Getenv("PGSQL")
	if connStr == "" {
		return nil, errors.New("could not get psql connection string")
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// the begin function will be used to execute any transaction
// you can use transaction object to prepare the transatcion , please alwayase use prepared statement
// tx ,_ := TxBegin();
// stmt , _ := tx.Prepare(query())
// stmt.Execute(...params)
func TxBegin(db *sql.DB) (*sql.Tx, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	return tx, nil
}
func TxCommit(tx *sql.Tx) error {
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

// takes in the transaction object and returns the error

func TxRollback(tx *sql.Tx) error {
	if err := tx.Rollback(); err != nil {
		return err
	}
	return nil
}

// this will create a new instance of the database  , not effecient , an old way of doing it
func ExecuteAuto(statement string) error {
	// this function will be used to execute any query
	db, err := Database()
	if err != nil {
		return err
	}
	defer db.Close()
	stmt, err := db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err

	}
	return nil
}
func (a *App) Execute(statement string) error {
	defer a.DB.Close()
	stmt, err := a.DB.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil

}
func (a *App) ExecuteArgs(
	statement string,
	args ...interface{}) error {
	defer a.DB.Close()
	stmt, err := a.DB.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(args...)
	if err != nil {
		return err
	}
	return nil
}
