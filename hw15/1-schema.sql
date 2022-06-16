-- CREATE DATABASE wiki_films;

-- удаление рейтингов (если есть)
DROP TYPE IF EXISTS ratings CASCADE;

-- удаление таблиц (если есть)
DROP TABLE IF EXISTS actors_films;
DROP TABLE IF EXISTS directors_films;
DROP TABLE IF EXISTS films;
DROP TABLE IF EXISTS actors;
DROP TABLE IF EXISTS directors;

-- enum для рейтингов
CREATE TYPE ratings AS ENUM ('PG-10', 'PG-13', 'PG-18');

-- фильмы (название, год выхода, актёры, режиссёры, сборы, рейтинг) и студии (название).
CREATE TABLE films (
    id SERIAL PRIMARY KEY ,
    title VARCHAR NOT NULL DEFAULT '',
    released_at  INTEGER NOT NULL CHECK ( released_at > 1800 ),
    fee INTEGER NOT NULL DEFAULT 0,
    rating ratings NOT NULL,
    studio VARCHAR NOT NULL,
    UNIQUE (title, released_at)
);
CREATE INDEX CONCURRENTLY IF NOT EXISTS films_title_idx ON films USING btree (lower(title));
CREATE INDEX CONCURRENTLY IF NOT EXISTS films_studio_idx ON films USING btree (lower(studio));

-- актёры (имя, дата рождения);
CREATE TABLE actors (
    id SERIAL PRIM`ARY KEY,
    full_name VARCHAR NOT NULL DEFAULT '',
    birthdate DATE` NOT NULL
);
CREATE INDEX CONCURRENTLY IF NOT EXISTS actors_full_name_idx ON actors USING btree (lower(full_name));

-- режиссёры (имя, дата рождения);
CREATE TABLE directors (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR NOT NULL DEFAULT '',
    birthdate DATE NOT NULL
);
CREATE INDEX CONCURRENTLY IF NOT EXISTS directors_full_name_idx ON directors USING btree (lower(full_name));

-- соединительная таблица актеров и фильмов
CREATE TABLE actors_films (
    id BIGSERIAL PRIMARY KEY,
    actor_id BIGINT NOT NULL REFERENCES actors(id) ON DELETE CASCADE ON UPDATE CASCADE DEFAULT 0,
    film_id INTEGER NOT NULL REFERENCES films(id) ON DELETE CASCADE ON UPDATE CASCADE DEFAULT 0,
    UNIQUE(actor_id, film_id)
);
CREATE INDEX CONCURRENTLY IF NOT EXISTS actors_films_actor_id_idx ON actors_films(actor_id);
CREATE INDEX CONCURRENTLY IF NOT EXISTS actors_films_film_id_idx ON actors_films(film_id);

-- соединительная таблица режиссеров и фильмов
CREATE TABLE directors_films (
    id BIGSERIAL PRIMARY KEY,
    director_id BIGINT NOT NULL REFERENCES directors(id) ON DELETE CASCADE ON UPDATE CASCADE DEFAULT 0,
    film_id INTEGER NOT NULL REFERENCES films(id) ON DELETE CASCADE ON UPDATE CASCADE DEFAULT 0,
    UNIQUE(director_id, film_id)
);
CREATE INDEX CONCURRENTLY IF NOT EXISTS directors_films_director_id_idx ON directors_films(director_id);
CREATE INDEX CONCURRENTLY IF NOT EXISTS directors_films_film_id_idx ON directors_films(film_id);
