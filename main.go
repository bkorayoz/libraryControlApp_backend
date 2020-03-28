package main

import (
	"libraryControlApp_backend/gateways/postgresql"
	"libraryControlApp_backend/server"
	"libraryControlApp_backend/util"
)

func main() {
	var cfg util.Config
	cfg.GetConfigData()

	postgresql.Init("localhost", cfg.Database.Name, cfg.Database.Username, cfg.Database.Password)
	defer postgresql.Close()

	go server.StartServer()
}
