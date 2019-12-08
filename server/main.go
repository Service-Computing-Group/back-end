package main

import (
	"os"

	"github.com/spf13/pflag"
)

const (
	PORT string = "8080"
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
	server := service.NewServer()
	server.Run(":" + port)
}
