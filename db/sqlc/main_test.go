package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *pgxpool.Pool
var testStore Store

const (
	dbSource = "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	// Initialize database connection
	testDB, err = pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("can't connect to db:", err)
	}

	// Initialize the queries and store
	testQueries = New(testDB)
	testStore = NewStore(testDB)

	// Run the tests
	code := m.Run()

	// Clean up
	if testDB != nil {
		testDB.Close()
	}

	os.Exit(code)
}
