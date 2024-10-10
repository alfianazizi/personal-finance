package main

import (
	"github.com/alfianazizi/personal-finance/config"
	"github.com/alfianazizi/personal-finance/database"
	"github.com/alfianazizi/personal-finance/server"
)

func main() {
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)
	server.NewFiberServer(conf, db).Start()
}
