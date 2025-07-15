package models

type Movie struct {
	ID          int
	TMDB_ID     int
	Title       string
	Tagline     *string // pointer - and can be nil
	ReleaseYear int
	Genres      []Genre // Collection
	Overview    *string
	Score       *float32
	Popularity  *float32
	Keywords    []string
	Language    *string
	PosterURL   *string
	TrailerURL  *string
	Casting     []Actor
}
