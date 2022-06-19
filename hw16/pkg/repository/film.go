package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kevin-glare/hardcode-dev-go/hw16/pkg/model"
)

type FilmRepo struct {
	db *pgxpool.Pool
}

func NewFilmRepo(db *pgxpool.Pool) *FilmRepo {
	return &FilmRepo{
		db: db,
	}
}

func(r *FilmRepo) AllFilm(ctx context.Context, studioID int) ([]model.Film, error) {
	films := make([]model.Film, 0)
	query := "SELECT id, title, released_at, fee, rating, studio_id FROM films"

	if studioID != 0 {
		query = fmt.Sprintf("%s where studio_id = %v", query, studioID)
	}

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return films, err
	}
	defer rows.Close()

	for rows.Next() {
		var film model.Film
		err = rows.Scan(&film.ID, &film.Title, &film.ReleasedAt, &film.Fee, &film.Rating, &film.StudioID)
		if err != nil {
			return films, err
		}

		films = append(films, film)
	}

	return films, rows.Err()
}

func (r *FilmRepo) AddFilms(ctx context.Context, films []model.Film) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	var batch = &pgx.Batch{}
	for _, film := range films {
		batch.Queue(
			`INSERT INTO films(title, released_at, fee, rating, studio_id) VALUES ($1, $2, $3, $4, $5)`,
			film.Title, film.ReleasedAt, film.Fee, film.Rating, film.StudioID,
		)
	}

	res := tx.SendBatch(ctx, batch)
	err = res.Close()
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *FilmRepo) RemoveFilm(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, "DELETE FROM films where id=$1", id)
	return err
}

func (r *FilmRepo) UpdateFilm(ctx context.Context, input model.Film) error {
	_, err := r.db.Exec(ctx, "update films set title=$1, released_at=$2, fee=$3, rating=$4, studio_id=$5 where id=$6",
		input.Title, input.ReleasedAt, input.Fee, input.Rating, input.StudioID, input.ID)
	return err
}