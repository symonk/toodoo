package main

import (
	"github.com/symonk/toodoo/cmd/server"
	"github.com/symonk/toodoo/internal/config"
	"github.com/symonk/toodoo/internal/db"
	"github.com/symonk/toodoo/internal/logging"
)

// main is the core entry point to the backend API.
// it is responsible for instantiating core services
// and handling graceful exit etc.
func main() {
	config.Init()
	logging.Init()
	// TODO: Read from config etc;
	db.Init("postgresql://postgres:postgres@localhost:5432/toodoo?sslmode=disable")
	server.Init()

}
