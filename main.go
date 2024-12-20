package main

import (
	"database/sql"
	"formative-14/controllers"
	"formative-14/database"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}

		DB, err = sql.Open("postgres", databaseURL)
		defer DB.Close()
		err = DB.Ping()
		if err != nil {
			panic(err)
		}
	}

	database.DBMigrate(DB)

	router := gin.Default()
	router.GET("/persons", controllers.GetAllPerson)
	router.POST("/persons", controllers.InsertPerson)
	router.PUT("/persons/:id", controllers.UpdatePerson)
	router.DELETE("/persons/:id", controllers.DeletePerson)

	router.Run(":8080")
}
