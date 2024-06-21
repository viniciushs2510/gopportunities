package config

import (
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	db, err = InitializeSQLite()
	if err != nil {
		return err
	}

	return nil
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}

func GetDB() *gorm.DB {
	return db
}
