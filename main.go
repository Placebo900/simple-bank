package main

import (
	"database/sql"
	"log"

	api "github.com/Placebo900/simple-bank/api"
	db "github.com/Placebo900/simple-bank/db/sqlc"
	"github.com/Placebo900/simple-bank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.ParseToConfig(".")
	if err != nil {
		log.Fatal("cannot read config:", err)
	}

	testDB, err := sql.Open(config.DriverName, config.DataSourceName)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.SQLStore{DB: testDB, Queries: db.New(testDB)}

	server := api.NewServer(&store)
	log.Fatal(server.Start(config.Address))
}
