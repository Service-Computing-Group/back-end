package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Service-Computing-Group/back-end/service"

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

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	pPort := pflag.StringP("port", "p", PORT, "PORT for httpd listening")
	pflag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}
	server := service.NewServer()
	server.Run(":" + port)
}
