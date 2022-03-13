package main

import (
	"context"
	"log"
	"messaging-backend/handler"
	"messaging-backend/repository"
	"messaging-backend/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	log.Print("Starting Server...")

	// Initializing Dependency
	dataSource, err := NewDataSource()
	if err != nil {
		log.Fatalf("Unable to initialize database %v\n", err)
	}

	userRepository := repository.NewPGUserRepository(&dataSource.DB)
	userService := service.NewUserService(userRepository)

	router := gin.Default()

	config := &handler.Config{
		R:           router,
		UserService: userService,
	}

	handler.NewHandler(config)

	// Create a http server
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Garceful server shutdown
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v\n", err)
		}
	}()

	log.Printf("Listening on port %v\n", server.Addr)

	// Wait for kill signal of channel
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// This blocks waits until a signal is passed into the quit channel
	<-quit

	// The context is used to inform the server it has 2 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Closing database
	if err := dataSource.close(); err != nil {
		log.Fatalf("a problem occurred while closing database: %v\n", err)
	}

	// Shutdown server
	log.Println("Shutting down server...")
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}
}
