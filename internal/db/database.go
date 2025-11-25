package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	if err := DB.AutoMigrate(&User{}, &Team{}, &PRequest{}, &Assignment{}); err != nil {
		log.Fatal("Failed to apply migrations: ", err)
	} else {
		log.Println("Database migration succeeded")
	}

	fmt.Println("Connection to the database has been established successfully")
}
