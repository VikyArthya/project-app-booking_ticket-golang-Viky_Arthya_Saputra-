package util

import (
	"go.uber.org/zap"
)

// Global logger instance
var Logger *zap.Logger

// Initialize zap logger
func InitLogger() {
	var err error
	Logger, err = zap.NewProduction() // menggunakan log level production (bisa ganti ke Development)
	if err != nil {
		panic("unable to initialize zap logger")
	}
}
