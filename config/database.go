package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	connStr := "postgresql://pgsql_yjyr_user:1DuPO6cwhs0oXIBWoHs79kYpVk76ZTHN@dpg-d03n52buibrs73aekdhg-a.singapore-postgres.render.com/pgsql_yjyr"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Error pinging DB:", err)
	}

	fmt.Println("Connected to the database")

	// Create table if not exists
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT,
		status TEXT DEFAULT 'pending',
		user_name TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err = DB.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Error creating tasks table:", err)
	}

	fmt.Println("Table created or already exists")
}
