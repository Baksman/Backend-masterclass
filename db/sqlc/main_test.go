package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"

	"github.com/baksman/backend_masterclass/util"
)

// const (
// 	dbDriver = "postgres"
// 	dbSource = "postgresql://ibrahim:ibrahim@localhost:5432/simple_bank?sslmode=disable"
// )

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	// var err error
	if err != nil {
		log.Fatalf("error loading config %v", err.Error())
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatalf("cannot connect to db %v", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())

}
