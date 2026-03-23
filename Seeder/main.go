package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {
	conn, err := GetDBConnection(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	fmt.Println("Seeding 100k Users...")
	users := make([][]interface{}, 100000)
	for i := 0; i < 100000; i++ {
		users[i] = []interface{}{fmt.Sprintf("User %d", i+1), fmt.Sprintf("user%d@example.com", i+1)}
	}
	_, _ = conn.CopyFrom(context.Background(), pgx.Identifier{"users"}, []string{"name", "email"}, pgx.CopyFromRows(users))

	fmt.Println("Seeding 500k Orders & 1M Items...")
	for b := 0; b < 50; b++ { // 50 batches of 10k orders
		orders := make([][]interface{}, 10000)
		for i := 0; i < 10000; i++ {
			userID := (i % 100000) + 1
			orders[i] = []interface{}{userID, 99.99, "COMPLETED", "Wide description text for hydration testing"}
		}
		_, _ = conn.CopyFrom(context.Background(), pgx.Identifier{"orders"}, []string{"user_id", "amount", "status", "description"}, pgx.CopyFromRows(orders))
	}

	// Seed Items (2 per order)
	for b := 0; b < 100; b++ {
		items := make([][]interface{}, 10000)
		for i := 0; i < 10000; i++ {
			orderID := (b * 5000) + (i / 2) + 1
			items[i] = []interface{}{orderID, "Product XYZ", 2, 49.99}
		}
		_, _ = conn.CopyFrom(context.Background(), pgx.Identifier{"items"}, []string{"order_id", "product_name", "quantity", "price"}, pgx.CopyFromRows(items))
	}
	fmt.Println("Seeding Complete!")
}
