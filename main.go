package db

import (
	"context"
	"log"

	"github.com/khafizullokh02/warehouse_management/api"
	db "github.com/khafizullokh02/warehouse_management/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://root:secret@localhost:5432/warehouse_management?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	store := db.NewStore(context.Background(), dbSource)
	server := api.NewServer(store)

	err := server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
