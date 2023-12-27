package db

import (
	"context"
	"database/sql"
	"log"

	"github.com/khafizullokh02/warehouse_management/api"
	db "github.com/khafizullokh02/warehouse_management/db/sqlc"
	"github.com/khafizullokh02/warehouse_management/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(context.Background(), config.DBSource)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
