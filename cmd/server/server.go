package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/symonk/toodoo/internal/logging"
)

// Init initializes the http server with the gin router
// as the handler
func Init() {
	logging.Logger.Info("Toodoo backend API server running.")
	err := godotenv.Load()
	if err != nil {
		logging.Logger.Error("unable to load environment configurations.")
		os.Exit(1)
	}
	logging.Logger.Info("successfully loaded environment config.")
	ctx, stopper := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stopper()
	router := NewRouter()
	server := &http.Server{
		// TODO: implement config etc.
		Addr:    ":9999",
		Handler: router,
	}
	// Run the server in a go routine and listen for any signals pushed onto the
	// notify channel for exiting.
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logging.Logger.Error(fmt.Sprintf("webserver error %s", err.Error()))
		}

	}()
	<-ctx.Done()

	// Give the server a period of time to gracefully shutdown.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		// TODO: also slog/fatal
		panic(err)
	}
	logging.Logger.Info("Server gracefully shutdown.")
}
