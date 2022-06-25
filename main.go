package main

import (
	"prod/config"
	"prod/db"
	"prod/server"

	"github.com/spf13/viper"
)

var conf = config.Config{}

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("server")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}
}

func main() {
	var gormDB = db.New(conf)
	db.Migrate(gormDB)
	server.Inject(gormDB)
	server := server.NewServer(conf)
	server.Run()
}
