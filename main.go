package main

import (
	"flag"
	"os"
	"path/filepath"
	"prod/conf"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	config *conf.Config
)

func init() {

	writeToFile := flag.Bool("f", false, "write logs to file")
	flag.Parse()

	// Set the config file name
	fileName := flag.String("c", "config", "config file name")

	absPath, err := filepath.Abs("./conf")
	if err != nil {
		panic("config file not found in " + filepath.Join(absPath))
	}

	viper.SetConfigName(*fileName)
	viper.AddConfigPath(absPath)

	if err = viper.ReadInConfig(); err != nil {
		logrus.Fatalf("could not read config file: %v", err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		logrus.Fatalf("config file invalid: %v", err)
	}

	// set logger
	level, err := logrus.ParseLevel(config.Log.Level)
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
	logrus.Info(config)
}