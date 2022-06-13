INSERT INTO films (title, released_at, fee, rating, studio)
VALUES ('Карты, деньги, два ствола', 1998, '1035000', 'PG-18', 'HandMade Films'),
       ('Побег из Шоушенка', 1994, '28418000', 'PG-13', 'Castle Rock Entertainment'),
       ('1+1', 2011, '1725813', 'PG-10', 'Gaumont');

INSERT INTO actors (full_name, birthdate)
VALUES ('Джейсон Флеминг', '1966-09-25'),
       ('Тим Роббинс', '1958-10-16'),
       ('Франсуа Клюзе', '1955-09-21'),
       ('Омар Си', '1978-01-20');

INSERT INTO actors_films (actor_id, film_id)
VALUES (1, 1),
       (2, 2),
       (3, 3),
       (4, 3);

INSERT INTO directors (full_name, birthdate)
VALUES ('Гай Ричи', '1968-10-10'),
       ('Фрэнк Дарабонт', '1959-01-28'),
       ('Оливье Накаш', '1973-04-15');

INSERT INTO directors_films (director_id, film_id)
VALUES (1,1),
       (2, 2),
       (3, 3);