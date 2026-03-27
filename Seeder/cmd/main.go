package main

import (
	"context"
	"log"

	"github.com/kamalakshancg/GolangPOC/config"
	seeder "github.com/kamalakshancg/GolangPOC/internal"
)

func main() {
	conn, err := config.GetDBConnection(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	seeder.CreateTables(conn)

	seeder.LoadUsersAndOrders(conn)

	seeder.CreateDBIndexAndOptimize(conn)
}
