package config

import (
	"JurrassicParkAPI/app/repository/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"os"
)

// Database modal
type Database struct {
	*sqlx.DB
}

// NewDatabaseConnection connects to the database that has been outlined in the .env using the keyword DB_DSN
func NewDatabaseConnection() *Database {
	host := os.Getenv("HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("USER")
	pwd := os.Getenv("PWD")
	dbName := os.Getenv("DB_NAME")
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user,
		pwd,
		host,
		port,
		dbName,
	)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal("Error attempting to connect to database: ", err)
	}
	log.Info("Successfully connected to database")

	if err != nil {
		log.Fatal("Could not auto Migrate, Error: ", err)
	}

	createTables(db)

	return &Database{
		DB: db,
	}
}

func createTables(db *sqlx.DB) {
	_, err := db.Exec(sql.CreateCrateTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(sql.CreateDinosaurTable)
	if err != nil {
		log.Fatal(err)
	}

}
