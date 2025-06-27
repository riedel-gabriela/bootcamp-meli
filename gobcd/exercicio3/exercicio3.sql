-- Mostrar todos os registros da tabela de filmes.
SELECT * FROM movies;
-- Exibir o primeiro nome, o sobrenome e a classificação de todos os atores.
SELECT first_name, last_name, rating FROM actors;
-- Exibir o título de todas as séries e usar aliases para que o nome da tabela e o campo estejam em inglês.
SELECT title AS series_title FROM series;
-- Exibir o primeiro nome e o sobrenome dos atores cuja classificação seja superior a 7,5.
SELECT first_name, last_name FROM actors WHERE rating > 7.5;
-- Exibir o título do filme, a classificação e os prêmios dos filmes com classificação superior a 7,5 e com mais de dois prêmios.
SELECT title, rating, awards FROM movies WHERE rating > 7.5 AND awards > 2;
-- Exibir o título do filme e a classificação ordenados por classificação em ordem crescente.
SELECT title, rating FROM movies ORDER BY rating ASC;
-- Exibir os títulos dos filmes e a classificação ordenados por classificação.
SELECT title, rating FROM movies ORDER BY rating;
-- Exibir os títulos dos três primeiros filmes no banco de dados.
SELECT title FROM movies LIMIT 3;
-- Exibir os 5 filmes mais bem classificados.
SELECT title, rating FROM movies ORDER BY rating DESC LIMIT 5;
-- Listar os 10 primeiros atores.
SELECT first_name, last_name FROM actors LIMIT 10;
-- Exibir o título e a classificação de todos os filmes cujo título é Toy Story.
SELECT title, rating FROM movies WHERE title = 'Toy Story';
-- Exibir todos os atores cujos nomes começam com Sam.
SELECT first_name, last_name FROM actors WHERE first_name LIKE 'Sam%';
-- Exibir o título dos filmes lançados entre 2004 e 2008.
SELECT title FROM movies WHERE release_year BETWEEN 2004 AND 2008;
-- Exibir o título dos filmes com classificação superior a 3, com mais de 1 prêmio e com data de lançamento entre 1988 e 2009.
SELECT title FROM movies WHERE rating > 3 AND awards > 1 AND release_year BETWEEN 1988 AND 2009;
-- Ordenar os resultados por classificação.
SELECT title, rating FROM movies WHERE rating > 3 AND awards > 1 AND release_year BETWEEN 1988 AND 2009 ORDER BY rating DESC;