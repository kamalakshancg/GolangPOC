# Seeder Go App

A simple Go application for seeding a PostgreSQL database with sample data: 100k users, 500k orders, and 1M items.

## Prerequisites

- Go 1.19 or later
- PostgreSQL database running locally
- A database named `poc_db` with the following tables:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    amount DECIMAL(10,2) NOT NULL,
    status VARCHAR(50) NOT NULL,
    description TEXT
);

CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    order_id INTEGER REFERENCES orders(id),
    product_name VARCHAR(255) NOT NULL,
    quantity INTEGER NOT NULL,
    price DECIMAL(10,2) NOT NULL
);
```

## Setup

1. Ensure PostgreSQL is running on `localhost:5432`
2. Create the database `poc_db`
3. Create the tables as shown above
4. Copy `.env.example` to `.env` and update the database connection details if needed

## Environment Variables

Create a `.env` file in the root directory with the following variables:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=poc_db
```

## Running the App

```bash
go run main.go
```

## Troubleshooting

- **Connection failed**: Ensure PostgreSQL is running and the connection string is correct.
- **Table does not exist**: Create the required tables in the `poc_db` database.
- **Permission denied**: Check that the PostgreSQL user has the necessary permissions to insert data.

## Dependencies

- [pgx](https://github.com/jackc/pgx) - PostgreSQL driver for Go
- [godotenv](https://github.com/joho/godotenv) - Load environment variables from .env file