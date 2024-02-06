package main

import (
	"fileService/internal/config"
	"fmt"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	// Init AWS storage

	// Init Redis for storing file metadata and cache

	// Init Kafka for message broking

	// Init server
}