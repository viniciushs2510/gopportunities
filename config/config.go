package config

import (
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	looger *Logger
)

func Init() error {
	return nil
}

func GetLogger(p string) *Logger {
	looger = NewLogger(p)
	return looger
}
