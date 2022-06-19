package main

import (
	"context"
	"github.com/kevin-glare/hardcode-dev-go/hw16/pkg/model"
	"github.com/kevin-glare/hardcode-dev-go/hw16/pkg/repository"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	db, err := repository.NewPostgresDB(ctx, "postgresql://localhost:5432/wiki_films")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close()

	repositories := repository.NewRepository(db)

	films, err := repositories.AllFilm(ctx, 0)
	if err != nil {
		log.Println(err)
	}
	log.Printf("AllFilm with StudioID=0: %+v\n", films)

	films, err = repositories.AllFilm(ctx, 1)
	if err != nil {
		log.Println(err)
	}
	log.Printf("AllFilm with StudioID=1: %+v\n", films)


	var newFilms = []model.Film{
		model.Film{ Title: "Test#1", ReleasedAt: 2000, Fee: 1_000_000, Rating: model.FilmRatingPG10, StudioID: 1 },
		model.Film{ Title: "Test#2", ReleasedAt: 2010, Fee: 2_000_000, Rating: model.FilmRatingPG13, StudioID: 2 },
	}

	err = repositories.AddFilms(ctx, newFilms)
	if err != nil {
		log.Println(err)
	}
	log.Println("films added")

	err = repositories.RemoveFilm(ctx, 21)
	if err != nil {
		log.Println(err)
	}
	log.Println("film removed")

	film := model.Film{ID: 22, Title: "Test#3", ReleasedAt: 2010, Fee: 2_000_000, Rating: model.FilmRatingPG13, StudioID: 2 }
	err = repositories.UpdateFilm(ctx, film)
	if err != nil {
		log.Println(err)
	}
	log.Println("film updated")

	films, err = repositories.AllFilm(ctx, 0)
	if err != nil {
		log.Println(err)
	}
	log.Printf("AllFilm: %+v\n", films)

	os.Exit(0)
}
