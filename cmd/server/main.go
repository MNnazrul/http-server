package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/codecrafters-io/http-server-starter-go/internal/server"
)

func main() {
	directory := getDirectory()
	srv := server.NewServer(server.DefaultPort, directory)
	
	if err := srv.Start(); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		os.Exit(1)
	}
}

func getDirectory() string {
	dirPtr := flag.String("directory", "", "Path to directory")
	flag.Parse()
	return *dirPtr
}

