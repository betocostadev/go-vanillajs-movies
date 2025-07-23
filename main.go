// run with 'air'
package main

import (
	"log"
	"net/http"
	"os"

	"betocosta.com/reelingit/handlers"
	"betocosta.com/reelingit/logger"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Initialize logger
	logInstance := initializeLogger()

	movieHandler := handlers.NewMovieHandler(logInstance)

	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)
	http.HandleFunc("/api/movies/random", movieHandler.GetRandomMovies)

	// Handler for static files (frontend)
	http.Handle("/", http.FileServer(http.Dir("public")))
	// fmt.Println("Serving files")

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port
	}
	addr := ":" + port
	logInstance.Info("Server starting on " + addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		logInstance.Error("Server failed to start", err)
		log.Fatalf("Server failed: %v", err)
	}
}

func initializeLogger() *logger.Logger {
	logInstance, err := logger.NewLogger("movie-service.log")
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	return logInstance
}
