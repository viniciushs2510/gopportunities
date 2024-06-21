package config

import (
	"os"

	"github.com/viniciushs2510/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	// Check if SQLite is available
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Infof("SQLite file does not exist, creating it.")

		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}

		file.Close()
	}

	// Create db and auto migrate schema
	db, error := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if error != nil {
		logger.Errorf("Error opening SQLite: %v", error)
		return nil, error
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Error migrating schema: %v", err)
		return nil, err
	}

	return db, nil
}
