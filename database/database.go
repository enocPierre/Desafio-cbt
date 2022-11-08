package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

func StartDB() {
	str := "host=localhost port=25432 user=admin dbname=transactions sslmode=disable password=123456"

	//database, err := gorm.Open((postgres.Open)(str), &gorm.Config{})
	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})

	if err != nil {
		fmt.Println("Could not connect to the Postgres Database")
		log.Fatal("Error: ", err)
	}

	db = database
	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
	//migrations.RunMigrations(db)
}

func GetDatabase() *gorm.DB {
	return db
}
