package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	logger := log.New(os.Stdout, "//", log.LstdFlags)
	serv := server.NewServer(logger)
	logger.Println("Starting the server ...")
	err := serv.HTTP.ListenAndServe()
	if err != nil {
		logger.Fatalf("starting the server error: %v", err)
	}
}
