package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Placebo900/simple-bank/util"
	_ "github.com/lib/pq"
)

var queries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.ParseToConfig("../..")
	if err != nil {
		log.Fatal("cannot read config:", err)
	}
	testDB, err = sql.Open(config.DriverName, config.DataSourceName)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	queries = New(testDB)

	os.Exit(m.Run())
}
