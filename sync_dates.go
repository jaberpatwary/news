package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=root dbname=news port=5432 sslmode=disable TimeZone=Asia/Dhaka"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Updating existing news records with updated_at = created_at...")

	// RAW SQL update for speed and reliability in migration
	err = db.Exec("UPDATE articles SET updated_at = created_at WHERE updated_at IS NULL OR updated_at = '0001-01-01 00:00:00+00'").Error
	if err != nil {
		log.Fatal("Failed to update records:", err)
	}

	fmt.Println("Success! All previous news records are now synchronized.")
}
