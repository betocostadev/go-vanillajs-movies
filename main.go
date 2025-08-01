// run with 'air'
package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"betocosta.com/reelingit/data"
	"betocosta.com/reelingit/logger"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Initialize logger
	logInstance := initializeLogger()

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Connect to DB
	dbConnStr := os.Getenv("DATABASE_URL")
	if dbConnStr == "" {
		log.Fatal("DATABASE URL Not Found")
	}

	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to the DB: %v", err)
	}
	// With defer it is not closing now, it will close later
	defer db.Close()

	// Initialize Data Repository for Movies
	movieRepo, err := data.NewMovieRepository(db, logInstance)
	if err != nil {
		log.Fatalf("Failed to initialize movierepository")
	}

	// It doest have complains in this way because we were not exporting storage
	// we were using storage instead of Storage in our movieHandle file
	// TODO - Decide on how to cann the movie handler
	// movieHandler := handlers.MovieHandler{
	// 	storage: movieRepo, logger: logInstance,
	// }
	// like new MovieHandler(movieRepo, logInstance)

	// now with Storage in movie_handlers.go we made it public and can access like below
	// movieHandler := handlers.MovieHandler{}
	// movieHandler.Storage = movieRepo
	// movieHandler.Logger = logInstance

	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)
	// http.HandleFunc("/api/movies/random", movieHandler.GetRandomMovies)

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
