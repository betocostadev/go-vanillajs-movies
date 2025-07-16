// run with 'air'
package main

import (
	"log"
	"net/http"

	"betocosta.com/reelingit/handlers"
	"betocosta.com/reelingit/logger"
)

func main() {
	// Initialize logger
	logInstance := initializeLogger()

	// Handler for static files (frontend)
	http.Handle("/", http.FileServer(http.Dir("public")))

	movieHandler := handlers.NewMovieHandler(logInstance)
	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)

	// Start server
	const addr = ":8080"
	logInstance.Info("Server starting on " + addr)
	// fmt.Println("Serving files")
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
