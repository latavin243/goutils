package main

import "go.uber.org/zap"

func main() {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	logger, _ := config.Build()
	defer logger.Sync()

	url := "localhost:8080/api"
	sugar := logger.Sugar()
	sugar.Infow("error", "url", url, "attempt", 3)
}
