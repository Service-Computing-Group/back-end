package main

import (
	"os"

	"github.com/Service-Computing-Group/back-end/database"

	"github.com/Service-Computing-Group/back-end/service"

	"github.com/spf13/pflag"
)

const (
	PORT string = "8081"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	pPort := pflag.StringP("port", "p", PORT, "PORT for httpd listening")
	pflag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}
	database.OpenDB("./database/test2.db")
	server := service.NewServer()
	server.Run(":" + port)

	defer database.CloseDB()
}
