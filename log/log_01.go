package main

import (
	"go.uber.org/zap"
	"time"
)

func main() {
	normalLog()
	performanceLog()
}
func normalLog() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", "url",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", "url")
}
func performanceLog() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", "url"),
		zap.Int("attempt", 3),
		zap.Time("time", time.Now()),
		zap.Duration("backoff", time.Second),
	)
}
