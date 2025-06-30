-- Exercício 1
-- Com a base de dados "movies", propõe-se criar uma tabela temporária chamada "TWD" e armazenar nela os episódios de todas as temporadas de "The Walking Dead".
-- Executa uma consulta à tabela de tempo para ver os episódios da primeira temporada.
CREATE TEMPORARY TABLE TWD AS
SELECT * FROM episodes WHERE series_id = (SELECT id FROM series WHERE title = 'The Walking Dead');

-- Exercício 2
-- No banco de dados "movies", selecione uma tabela para criar um índice e, em seguida, verifique a criação do índice.
-- Analise por que você criaria um índice na tabela indicada e com quais critérios você escolheria o(s) campo(s).

-- Criei um índice na tabela "movies" sobre os campos "id" e "title".
-- Isso pode melhorar o desempenho das consultas que filtram ou ordenam por esses campos, especialmente em tabelas grandes.
CREATE INDEX idx_movies_title ON movies(id, title);