package main

import (
	"database/sql"
	"moody/moody-api/src/database"
	"moody/moody-api/src/endpoints"
	"moody/moody-api/src/internals"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type App struct {
	DB *sql.DB
}

func main() {
	// load go file
	err := godotenv.Load()
	if err != nil {
		os.Exit(-1)
	}
	// do not create any database operations before the database isn initiated
	Database := App{}
	Database.DbInit(internals.DB_MAX_OPEN_CONNS(), internals.DB_MAX_IDLE_CONNS(), internals.DB_SET_MAX_CONN_LIFETIME())
	// the new clearner migration method
	database.BetterMigration()
	router := gin.Default()
	router.Use(cors.Default())
	endpoints.InjectEndpoints(router)
	router.Run()
}
