package main

import (
	"database/sql"
	"log"

	"github.com/HtetOoNaing/simple-bank-backend-master-class-golang-postgres-kubernetes-gRPC/api"
	db "github.com/HtetOoNaing/simple-bank-backend-master-class-golang-postgres-kubernetes-gRPC/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}