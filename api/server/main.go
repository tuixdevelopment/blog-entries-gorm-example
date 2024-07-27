package main

import "stellar_backend/internal/server"

func main() {
	server.Server.Start(":8080")
}
