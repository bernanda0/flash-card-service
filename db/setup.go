package db

import (
	"br/simple-service/db/sqlc"
	"database/sql"
	"log"

	_ "github.com/lib/pq" // Import the PostgreSQL driver package
)

func Instantiate(l *log.Logger) (*sql.DB, *sqlc.Queries) {
	connStr := "user=bernanda password=bernanda dbname=sr-db sslmode=disable"
	db, err1 := sql.Open("postgres", connStr)
	if err1 != nil {
		l.Println("Error connecting to DB ", err1)
		return nil, nil
	}

	err2 := db.Ping()
	if err2 != nil {
		l.Println("Error connecting to DB ", err2)
		return nil, nil
	}

	l.Println("🛢️  DB Connected")
	return db, sqlc.New(db)
}
