package server

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/symonk/toodoo/internal/logging"
)

// Init initializes the http server with the gin router
// as the handler
func Init() {
	logging.Logger.Info("Toodoo backend API server running.")
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
			// TODO: Implment logging slog and fatal here
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
	fmt.Println("Server gracefully shutdown.")
}
