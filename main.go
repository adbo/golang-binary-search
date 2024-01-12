package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"main/data"
	"main/handlers"
	"net/http"
)

func main() {
	initConfig()
	logger := initLogger()

	store := data.NewNumberStore(logger)
	err := store.LoadNumbers("input.txt")
	if err != nil {
		logger.WithError(err).Fatal("Failed to load numbers")
	}

	handler := handlers.NewNumberIndexHandler(store, logger)
	http.HandleFunc("/endpoint/", handler.HandleNumberIndexRequest)

	port := viper.GetString("server.port")
	logger.Infof("Starting server on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logger.WithError(err).Fatal("Failed to start server")
	}
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}

func initLogger() *logrus.Logger {
	logger := logrus.New()
	level, err := logrus.ParseLevel(viper.GetString("logging.level"))
	if err != nil {
		logger.Warn("Logging level not set correctly, defaulting to 'info'")
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)
	return logger
}
