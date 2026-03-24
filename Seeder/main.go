package main

import (
	"context"
	"log"
)

func main() {
	conn, err := GetDBConnection(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	CreateTables(conn)

	LoadToDatabase(conn)
}
