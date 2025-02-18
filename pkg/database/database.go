package database

import(
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func Conect(connStr string) (*sql.DB, error) {
	log.Printf("Attempting to connect to the database with connection string: %s", connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("Error opening database connection: %v", err)
		return nil, err
	}

	log.Println("Pinging the database to check connectivity...")
	err = db.Ping()
	if err != nil {
		log.Printf("Error pinging the database: %v", err)
		return nil, err
	}

	log.Println("Database connection established successfully")
	return db, nil
}