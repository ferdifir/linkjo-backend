package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"entgo.io/ent/dialect/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Import driver PostgreSQL

	"linkjo/ent" // Ganti dengan nama proyek
)

var Client *ent.Client

func ConnectDB() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Format DSN untuk PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Membuat driver database menggunakan ent/dialect/sql
	drv, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to open database driver:", err)
	}

	// Membuat client Ent menggunakan driver yang telah dibuat
	client := ent.NewClient(ent.Driver(drv))

	// Tes koneksi database
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal("Failed to create schema:", err)
	}

	Client = client
	fmt.Println("Database connected successfully with Ent!")
}
