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

const (
	dbSource = "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable"
)

var testStore Store
func TestMain(m *testing.M) {
	var err error
	testDB, err = pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("can't connect to db:", err)
	}

	testQueries = New(testDB) // Now testDB correctly implements DBTX
	os.Exit(m.Run())
}
