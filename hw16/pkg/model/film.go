package model

type FilmRating string

const (
	FilmRatingPG10 FilmRating = "PG-10"
	FilmRatingPG13 FilmRating = "PG-13"
	FilmRatingPG18 FilmRating = "PG-18"
)

type Film struct {
	ID         int        `json:"id" db:"id"`
	Title      string     `json:"title" db:"title"`
	ReleasedAt int        `json:"released_at" db:"released_at"`
	Fee        int        `json:"fee" db:"fee"`
	Rating     FilmRating `json:"rating" db:"rating"`
	StudioID   int        `json:"studio_id" db:"studio_id"`
}
