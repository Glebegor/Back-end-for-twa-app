package main

import (
	logrus "github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// if err := initConfig(); err != nil {}
	// if err := godotenv.Load(); err != nil {}
	// db, err := repositoryes.ConnectDB(repository.ConfigDB{})

	return
}
