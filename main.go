package main

import (
	"github.com/symonk/toodoo/internal/config"
	"github.com/symonk/toodoo/internal/db"
	"github.com/symonk/toodoo/internal/logging"
	"github.com/symonk/toodoo/internal/server"
)

// main is the core entry point to the backend API.
// it is responsible for instantiating core services
// and handling graceful exit etc.
func main() {
	config.Init()
	logging.Init()
	db.Init("")
	server.Init()

}
