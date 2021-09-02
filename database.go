package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Returns a connection handle for an in-memory sqlite database.
func GetDatabaseConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return &gorm.DB{}, err
	}

	prepareTables(db)

	return db, nil
}

// Create some database tables if they don't already exist. In a production application,
// we'd either not do this or populate demo date with a migration or something.
func prepareTables(db *gorm.DB) {

	db.AutoMigrate(&Employee{})
	employees := []Employee{
		{
			Name: "Bilbo Baggins",
			Role: "Burglar",
		},
		{
			Name: "Frodo Baggins",
			Role: "Ringbearer",
		},
	}

	db.Create(&employees)
}
