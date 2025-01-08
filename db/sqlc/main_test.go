package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:postgres@localhost:5432/go_project_2?sslmode=disable"
)

var testQueries * Queries
var testDB * pgxpool.Pool

func TestMain(m *testing.M) {
	var err error
	testDB, err = pgxpool.New(context.Background(), dbSource)

	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	defer testDB.Close()

	testQueries = New(testDB)

	os.Exit(m.Run())
}