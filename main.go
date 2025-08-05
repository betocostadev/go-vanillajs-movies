// run with 'air'
package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"betocosta.com/reelingit/data"
	"betocosta.com/reelingit/handlers"
	"betocosta.com/reelingit/logger"
	"betocosta.com/reelingit/routes"
	"betocosta.com/reelingit/utils"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func initializeLogger() *logger.Logger {
	logInstance, err := logger.NewLogger(utils.LogFileName)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	return logInstance
}

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
		log.Fatalf("Failed to initialize movie repository: %v", err)
	}

	// It doest have complains in this way because we were not exporting storage
	// we were using storage instead of Storage in our movieHandle file
	// We could use this way like a factory
	// movieHandler := handlers.NewMovieHandler(movieRepo, logInstance)
	// or
	// movieHandler := handlers.MovieHandler{
	// 	storage: movieRepo, logger: logInstance,
	// }

	// now with Storage in movie_handlers.go we made it public and can access like below
	movieHandler := handlers.MovieHandler{}
	movieHandler.Storage = movieRepo
	movieHandler.Logger = logInstance

	// Initialize handlers
	// Order is important
	http.HandleFunc(routes.MoviesTopRoute, movieHandler.GetTopMovies)
	http.HandleFunc(routes.MoviesRandomRoute, movieHandler.GetRandomMovies)
	http.HandleFunc(routes.MoviesSearchRoute, movieHandler.SearchMovies)
	http.HandleFunc(routes.MoviesRoute, movieHandler.GetMovie)
	http.HandleFunc(routes.GenresRoute, movieHandler.GetGenres)
	http.HandleFunc(routes.AccountRegisterRoute, movieHandler.GetGenres)
	http.HandleFunc(routes.AccountAuthenticateRoute, movieHandler.GetGenres)

	// Handler for static files (frontend)
	http.Handle("/", http.FileServer(http.Dir("public")))
	// fmt.Println("Serving files")

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = utils.DefaultPort // default port
	}
	addr := ":" + port
	logInstance.Info("Server starting on " + addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		logInstance.Error("Server failed to start", err)
		log.Fatalf("Server failed: %v", err)
	}

}
