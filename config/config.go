package config

import (
	"gorm.io/gorm"
)

// package var
var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	return nil
}

func GetLogger(p string) *Logger {
	logger := NewLogger(p)
	return logger
}
