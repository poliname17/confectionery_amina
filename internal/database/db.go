package database

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(dsn string) *gorm.DB {
	var db *gorm.DB
	var err error

	for i := 0; i < 10; i++ { // пробуем 10 раз
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("✅ Successfully connected to the database")
			return db
		}
		log.Println("Waiting for database...", err)
		time.Sleep(2 * time.Second) // ждём 2 секунды перед следующей попыткой
	}

	log.Fatal("❌ Failed to connect to database:", err)
	return nil
}
