-- Você está solicitando o banco de dados movies_db.sql:

-- Adicionar um filme à tabela de movies.
INSERT INTO movies (title, release_year, rating, awards, genre_id)
VALUES ('Inception', 2010, 8.8, 4, NULL);
-- Adicione um gênero à tabela de genres.
INSERT INTO genres (name, ranking, active)
VALUES ('Science Fiction', 1, TRUE);
-- Associe o gênero criado no ponto 2 ao filme no ponto 1. gênero.
UPDATE movies
SET genre_id = (SELECT id FROM genres WHERE name = 'Science Fiction')
WHERE title = 'Inception';
-- Modifique a tabela de atores para que pelo menos um ator tenha como favorito o filme adicionado no ponto 1.
UPDATE actors
SET favorite_movie_id = (SELECT id FROM movies WHERE title = 'Inception')
WHERE first_name = 'Leonardo' AND last_name = 'DiCaprio';
-- Crie uma cópia temporária da tabela de movies.
CREATE TEMPORARY TABLE temp_movies AS
SELECT * FROM movies;
-- Remova dessa tabela temporária todos os filmes que ganharam menos de 5 awards.
DELETE FROM temp_movies
WHERE awards < 5;
-- Obtenha a lista de todos os gêneros que têm pelo menos um movies.
SELECT DISTINCT g.name
FROM genres g
JOIN movies m ON g.id = m.genre_id
WHERE m.id IS NOT NULL;
-- Obtenha a lista de atores cujo filme favorito ganhou mais de 3 awards.
SELECT a.first_name, a.last_name
FROM actors a
JOIN movies m ON a.favorite_movie_id = m.id
WHERE m.awards > 3;
-- Crie um índice sobre o nome na tabela de movies.
CREATE INDEX idx_movies_title ON movies(title);
-- Verifique se o índice foi criado corretamente.
-- Para verificar se o índice foi criado corretamente, você pode usar o comando:
SELECT * FROM pg_indexes WHERE tablename = 'movies' AND indexname = 'idx_movies_title';
-- No banco de dados de movies, você notaria uma melhora significativa com a criação de índices? Analise e justifique sua resposta.
-- Sim, a criação de índices no banco de dados "movies" pode melhorar significativamente o desempenho das consultas, especialmente em tabelas grandes.
-- Em qual outra tabela você criaria um índice e por quê? Justifique sua resposta.
-- Eu criaria um índice na tabela "actors" sobre os campos "first_name" e "last_name", pois isso aceleraria as consultas que filtram atores por nome.