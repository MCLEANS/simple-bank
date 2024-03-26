package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

/*
*
Contains a DBTX member which can either be a connection or a transaction
*
*/
var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable"
)

/*
*
Entry point of all Go unit tests inside a particular package
*
*/
func TestMain(m *testing.M) {

	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Cannot connect to Db", err)
	}

	/**
	Use the connection to create out testQueries Object
	**/
	testQueries = New(conn)

	/**
	This will run the tests and return a error code to tell us whether the test passed or failed
	**/
	os.Exit(m.Run())
}
