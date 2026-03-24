package seeder

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/jackc/pgx/v5"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func SeedData(conn *pgx.Conn) {
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

func generateRandom1KBString() string {
	b := make([]byte, 1024)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func LoadToDatabase(conn *pgx.Conn) {
	ctx := context.Background()
	startTime := time.Now()

	// 1. Seed Users First (Required for Foreign Key)
	fmt.Println("Seeding 100,000 Users...")
	userBatchSize := 10000
	for i := 0; i < 100000; i += userBatchSize {
		userRows := make([][]interface{}, userBatchSize)
		for j := 0; j < userBatchSize; j++ {
			id := i + j + 1
			userRows[j] = []interface{}{
				fmt.Sprintf("User %d", id),
				fmt.Sprintf("user%d@example.com", id),
			}
		}
		_, err := conn.CopyFrom(
			ctx,
			pgx.Identifier{"users"},
			[]string{"name", "email"},
			pgx.CopyFromRows(userRows),
		)
		if err != nil {
			log.Fatalf("User seeding failed: %v", err)
		}
	}

	// 2. Seed 1 Million Orders
	fmt.Println("Seeding 1,000,000 Orders (Defeating DB compression)...")

	var payloads [50]string
	for i := 0; i < 50; i++ {
		payloads[i] = generateRandom1KBString()
	}

	totalOrders := 1000000
	batchSize := 10000
	batches := totalOrders / batchSize

	for b := 0; b < batches; b++ {
		orders := make([][]interface{}, batchSize)
		for i := 0; i < batchSize; i++ {
			// This now maps to IDs 1 through 100,000 which we just created
			userID := (i % 100000) + 1
			randomPayload := payloads[rand.Intn(len(payloads))]
			orders[i] = []interface{}{userID, 99.99, "COMPLETED", randomPayload}
		}

		_, err := conn.CopyFrom(
			ctx,
			pgx.Identifier{"orders"},
			[]string{"user_id", "amount", "status", "description"},
			pgx.CopyFromRows(orders),
		)
		if err != nil {
			log.Fatalf("Batch %d failed: %v\n", b, err)
		}

		if (b+1)%10 == 0 {
			fmt.Printf("Inserted %d / %d orders...\n", (b+1)*batchSize, totalOrders)
		}
	}

	fmt.Printf("Seeding Complete in %s!\n", time.Since(startTime))
}

func CreateTables(conn *pgx.Conn) {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS orders (
			id SERIAL PRIMARY KEY,
			user_id INT NOT NULL REFERENCES users(id),
			amount NUMERIC(10, 2) NOT NULL,
			status VARCHAR(50) NOT NULL,
			description TEXT
		);`,
		`CREATE TABLE IF NOT EXISTS items (
			id SERIAL PRIMARY KEY,
			order_id INT NOT NULL REFERENCES orders(id),
			product_name VARCHAR(255) NOT NULL,
			quantity INT NOT NULL,
			price NUMERIC(10, 2) NOT NULL
		);`,
	}

	for _, query := range queries {
		_, err := conn.Exec(context.Background(), query)
		if err != nil {
			log.Fatalf("Failed to execute query: %v\nQuery: %s\n", err, query)
		}
	}

	fmt.Println("Tables created successfully!")
}
