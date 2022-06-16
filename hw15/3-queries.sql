-- выборка фильмов с названием студии;
SELECT * FROM films;

-- выборка фильмов для некоторого актёра;
SELECT films.* FROM films
JOIN actors_films af ON films.id = af.film_id
JOIN actors a ON a.id = af.actor_id
WHERE a.full_name = 'Омар Си';

-- подсчёт фильмов для некоторого режиссёра;
SELECT COUNT(films.*) FROM films
JOIN directors_films df ON films.id = df.film_id
JOIN directors d ON d.id = df.director_id
WHERE d.full_name = 'Гай Ричи';

-- выборка фильмов для нескольких режиссёров из списка (подзапрос);
SELECT COUNT(DISTINCT(films.*)) FROM films
JOIN directors_films df ON films.id = df.film_id
JOIN directors d ON d.id = df.director_id
WHERE d.full_name = 'Гай Ричи' OR d.birthdate > '1968-01-01';

-- подсчёт количества фильмов для актёра;
SELECT COUNT(films.*) FROM films
JOIN actors_films af ON films.id = af.film_id
JOIN actors a ON a.id = af.actor_id
WHERE a.full_name = 'Омар Си';

-- выборка актёров и режиссёров, участвовавших более чем в 2 фильмах;
SELECT full_name, COUNT(*)
FROM (
    SELECT a.full_name as full_name FROM films
    JOIN actors_films af ON films.id = af.film_id
    JOIN actors a ON a.id = af.actor_id
    UNION ALL
    SELECT d.full_name as full_name FROM films
    JOIN directors_films df ON films.id = df.film_id
    JOIN directors d ON d.id = df.director_id
    ) all_records
GROUP BY full_name
HAVING COUNT(*) > 2;

-- подсчёт количества фильмов со сборами больше 1000;
SELECT COUNT(films.*) FROM films
WHERE fee > 1000;

-- подсчитать количество режиссёров, фильмы которых собрали больше 1000;
SELECT films.title, count(d.*) from films
JOIN directors_films df ON films.id = df.film_id
JOIN directors d ON d.id = df.director_id
WHERE films.fee > 1000
GROUP BY films.title;

-- выборка различных фамилий актёров;
SELECT DISTINCT(split_part(full_name, ' ', 2)) FROM actors;

-- подсчёт количества фильмов, имеющих дубли по названию.
SELECT title, count(films.*) FROM films
GROUP BY title
HAVING count(films.*) > 1;
