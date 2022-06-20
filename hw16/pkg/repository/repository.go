package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kevin-glare/hardcode-dev-go/hw16/pkg/model"
)

type Film interface {
	AllFilm(ctx context.Context, studioID int) ([]model.Film, error)
	AddFilms(ctx context.Context, films []model.Film) error
	RemoveFilm(ctx context.Context, id int) error
	UpdateFilm(ctx context.Context, input model.Film) error
}

type Repository struct {
	Film
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		Film: NewFilmRepo(db),
	}
}
