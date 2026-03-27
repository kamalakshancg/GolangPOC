package seeder

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jmoiron/sqlx"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateRandom1KBString() string {
	b := make([]byte, 1024)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func LoadUsersAndOrders(conn *pgx.Conn) {
	ctx := context.Background()
	startTime := time.Now()

	// 1. Start a Transaction
	tx, err := conn.Begin(ctx)
	if err != nil {
		log.Fatal("Could not start transaction:", err)
	}
	defer tx.Rollback(ctx)

	// 2. Clear existing data
	fmt.Println("Cleaning tables and resetting identities...")
	tx.Exec(ctx, "TRUNCATE users, orders, items RESTART IDENTITY CASCADE;")

	// 3. Seed 100,000 Users
	fmt.Println("Seeding 100,000 Users...")
	userBatchSize := 10000
	for i := 0; i < 100000; i += userBatchSize {
		userRows := make([][]interface{}, userBatchSize)
		for j := 0; j < userBatchSize; j++ {
			id := i + j + 1
			userRows[j] = []interface{}{id, fmt.Sprintf("User %d", id), fmt.Sprintf("user%d@example.com", id)}
		}
		tx.CopyFrom(ctx, pgx.Identifier{"users"}, []string{"id", "name", "email"}, pgx.CopyFromRows(userRows))
	}

	// 4. Seed 1,000,000 Orders
	fmt.Println("Seeding 1,000,000 Orders...")
	totalOrders := 1000000
	orderBatchSize := 10000

	// Pre-generate random payloads
	payloads := make([]string, 50)
	for i := 0; i < 50; i++ {
		payloads[i] = generateRandom1KBString()
	}

	// ---> NEW: Possible Statuses
	statuses := []string{"FAILED", "COMPLETED", "RETURN"}

	for b := 0; b < totalOrders/orderBatchSize; b++ {
		orderRows := make([][]interface{}, orderBatchSize)
		for i := 0; i < orderBatchSize; i++ {
			orderID := (b * orderBatchSize) + i + 1
			userID := (orderID % 100000) + 1

			// ---> NEW: Random amount between 100.00 and 1000.00
			randomAmount := 100.0 + rand.Float64()*(1000.0-100.0)

			// ---> NEW: Random status from the slice
			randomStatus := statuses[rand.Intn(len(statuses))]

			orderRows[i] = []interface{}{
				orderID,
				userID,
				randomAmount,
				randomStatus,
				payloads[rand.Intn(50)],
			}
		}
		_, err = tx.CopyFrom(
			ctx,
			pgx.Identifier{"orders"},
			[]string{"id", "user_id", "amount", "status", "description"},
			pgx.CopyFromRows(orderRows),
		)
		if err != nil {
			log.Fatalf("Order batch %d failed: %v", b, err)
		}

		if (b+1)%20 == 0 {
			fmt.Printf("Inserted %d / %d orders...\n", (b+1)*orderBatchSize, totalOrders)
		}
	}

	// 5. Commit
	if err := tx.Commit(ctx); err != nil {
		log.Fatal("Failed to commit users and orders:", err)
	}

	// 6. Sync sequences
	conn.Exec(ctx, "SELECT setval('users_id_seq', (SELECT MAX(id) FROM users))")
	conn.Exec(ctx, "SELECT setval('orders_id_seq', (SELECT MAX(id) FROM orders))")

	fmt.Printf("Users and Orders seeded in %s!\n", time.Since(startTime))

	// Note: Total items changed to 2,000,000 as per your SeedItems call
	SeedItems(ctx, conn, 2000000, totalOrders)
}

func SeedItems(ctx context.Context, conn *pgx.Conn, totalItems int, maxOrderID int) {
	fmt.Printf("Starting random insertion of %d items...\n", totalItems)
	startTime := time.Now()

	batchSize := 10000
	batches := totalItems / batchSize

	for b := 0; b < batches; b++ {
		itemRows := make([][]interface{}, batchSize)
		for i := 0; i < batchSize; i++ {
			// Generate a random order_id between 1 and 1,000,000
			randomOrderID := rand.Intn(maxOrderID) + 1

			itemRows[i] = []interface{}{
				randomOrderID,
				fmt.Sprintf("Product-%d", rand.Intn(500)), // 500 unique products
				rand.Intn(5) + 1, // Quantity 1-5
				19.99,            // Price
			}
		}

		_, err := conn.CopyFrom(
			ctx,
			pgx.Identifier{"items"},
			[]string{"order_id", "product_name", "quantity", "price"},
			pgx.CopyFromRows(itemRows),
		)
		if err != nil {
			log.Fatalf("Item batch %d failed: %v", b, err)
		}

		if (b+1)%50 == 0 {
			fmt.Printf("Inserted %d / %d items...\n", (b+1)*batchSize, totalItems)
		}
	}

	fmt.Printf("Items seeded successfully in %s\n", time.Since(startTime))
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

// OptimizeDatabase creates necessary indexes and updates statistics
func OptimizeDatabase(db *sqlx.DB) error {
	start := time.Now()
	log.Println("⚙️ Starting database optimization (Indexes & Vacuum)...")

	// 1. Define the Index Queries
	// Using IF NOT EXISTS makes this idempotent (safe to run multiple times)
	indexQueries := []string{
		"CREATE INDEX IF NOT EXISTS idx_orders_user_id ON orders(user_id);",
		"CREATE INDEX IF NOT EXISTS idx_items_order_id ON items(order_id);",
		"CREATE INDEX IF NOT EXISTS idx_orders_status_amount ON orders(status, amount);",
		"CREATE INDEX IF NOT EXISTS idx_users_id ON users(id);",
	}

	// 2. Execute Indexes
	for _, query := range indexQueries {
		_, err := db.Exec(query)
		if err != nil {
			return fmt.Errorf("failed to execute index query: %w", err)
		}
	}
	log.Println("✅ Indexes verified/created successfully.")

	// 3. Run VACUUM ANALYZE
	// This forces PostgreSQL to scan the tables and update its internal statistics
	// so it actually uses the indexes we just built.
	log.Println("🧹 Running VACUUM ANALYZE (This may take a few seconds on a 1GB DB)...")
	_, err := db.Exec("VACUUM ANALYZE;")
	if err != nil {
		return fmt.Errorf("failed to run VACUUM ANALYZE: %w", err)
	}

	log.Printf("🚀 Database optimization complete! Took %v\n", time.Since(start))
	return nil
}
