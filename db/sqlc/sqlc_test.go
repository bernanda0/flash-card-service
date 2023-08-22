package sqlc

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	DB_DRIVER            = "postgres"
	DB_CONNECTION_STRING = "postgresql://bernanda:bernanda@localhost:5432/sr-db-test?sslmode=disable"
)

var q_test *Queries

func TestMain(m *testing.M) {
	db, err1 := sql.Open(DB_DRIVER, DB_CONNECTION_STRING)
	if err1 != nil {
		fmt.Println("Error creating DB connection", err1)
	}

	err2 := db.Ping()
	if err2 != nil {
		fmt.Println("Error connecting to DB ", err2)
	}

	q_test = New(db)

	os.Exit(m.Run())
}
