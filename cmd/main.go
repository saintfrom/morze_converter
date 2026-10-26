package main

import (
	"log"
	"os"

	"morze_converter/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "server: ", log.LstdFlags)

	srv := server.NewServer(logger)
	if err := srv.Start(); err != nil {
		logger.Fatal(err)
	}
}
