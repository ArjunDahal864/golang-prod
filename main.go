package main

import (
	"flag"
	"golang-prod/conf"
	"golang-prod/db"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	config *conf.Config
)

func loadConfig(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env.example")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
}

func init() {

	writeToFile := flag.Bool("f", false, "write logs to file")
	flag.Parse()

	loadConfig(".")
	
	// set logger
	level, err := logrus.ParseLevel(config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     true,
	})
	logrus.SetReportCaller(true)
	if *writeToFile {
		f, err := os.OpenFile("log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
		if err != nil {
			panic(err)
		}
		logrus.SetOutput(f)
	}
}

func main() {
	dbConnection := db.New(*config)
	dbConnection.ConnectToSQliteDatabase()
	logrus.Info("connected to database")
}
