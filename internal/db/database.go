package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	if err := db.AutoMigrate(&User{}, &Team{}, &PR{}, &Assignment{}); err != nil {
		log.Fatal("Failed to apply migrations: ", err)
	} else {
		log.Println("Database migration succeeded")
	}

	fmt.Println("Connection to the database has been established successfully")
}
