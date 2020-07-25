package apiserver

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

// APIServer struct.
type APIServer struct {
	config *Config
}

// NewServer method for init new server instance.
func NewServer(config *Config) *APIServer {
	return &APIServer{
		config: config,
	}
}

// Start method for start new server.
func (s *APIServer) Start() {
	// Define waiting time.
	var wait time.Duration

	// Initialize a new router.
	router := mux.NewRouter()

	// Add static files, if prefix and path was defined in config.
	if s.config.Static.Prefix != "" && s.config.Static.Path != "" {
		router.PathPrefix(s.config.Static.Prefix).Handler(
			http.StripPrefix(
				s.config.Static.Prefix,
				http.FileServer(
					http.Dir(s.config.Static.Path),
				),
			),
		)
	}

	// Register API routes.
	Routes(router)

	// Configure API server.
	server := &http.Server{
		Handler: router,
		Addr:    s.config.Server.Host + ":" + s.config.Server.Port,

		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C).
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	server.Shutdown(ctx)

	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("API server is shutting down...")
	os.Exit(0)
}
