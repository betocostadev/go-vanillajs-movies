package handlers

import (
	"encoding/json"
	"net/http"

	"betocosta.com/reelingit/data"
	"betocosta.com/reelingit/logger"
)

type MovieHandler struct {
	Storage data.MovieStorage
	Logger  *logger.Logger
}

// NewMovieHandler creates a new MovieHandler instance
func NewMovieHandler(storage data.MovieStorage, log *logger.Logger) *MovieHandler {
	return &MovieHandler{
		Storage: storage,
		Logger:  log,
	}
}

// Utility functions
func (h *MovieHandler) writeJSONResponse(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		h.Logger.Error("JSON encoding error", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return err
	}
	return nil
}

func (h *MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.Storage.GetTopMovies()
	if err != nil {
		http.Error(w, "Internal error fetching top Movies", 500)
		h.Logger.Error("Get Top Movies Error", err)
	}
	h.writeJSONResponse(w, movies)

	if h.writeJSONResponse(w, movies) == nil {
		h.Logger.Info("Successfully served top movies")
	}
}

func (h *MovieHandler) GetRandomMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.Storage.GetRandomMovies()
	if err != nil {
		http.Error(w, "Internal error fetching random Movies", 500)
		h.Logger.Error("Get Random Movies Error", err)
	}
	h.writeJSONResponse(w, movies)

	if h.writeJSONResponse(w, movies) == nil {
		h.Logger.Info("Successfully served random movies")
	}
}

// writer to send data back to the client
// reader to get the data
// wrap the handler before the function name
// hardcoded for when we were not using the DB, just fake data
/* func (h *MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {
	movies := []models.Movie{
		{
			ID:          1,
			TMDB_ID:     101,
			Title:       "The Hacker",
			ReleaseYear: 2022,
			Genres:      []models.Genre{{ID: 1, Name: "Thriller"}},
			Keywords:    []string{"hacking", "cybercrime"},
			Casting:     []models.Actor{{ID: 1, FirstName: "Jane", LastName: "Doe"}},
		},
		{
			ID:          2,
			TMDB_ID:     102,
			Title:       "Space Dreams",
			ReleaseYear: 2020,
			Genres:      []models.Genre{{ID: 2, Name: "Sci-Fi"}},
			Keywords:    []string{"space", "exploration"},
			Casting:     []models.Actor{{ID: 2, FirstName: "John", LastName: "Star"}},
		},
		{
			ID:          3,
			TMDB_ID:     103,
			Title:       "The Lost City",
			ReleaseYear: 2019,
			Genres:      []models.Genre{{ID: 3, Name: "Adventure"}},
			Keywords:    []string{"jungle", "treasure"},
			Casting:     []models.Actor{{ID: 3, FirstName: "Lara", LastName: "Hunt"}},
		},
	}

	if h.writeJSONResponse(w, movies) == nil {
		h.Logger.Info("Successfully served top movies")
	}
} */
