-- Primeira Parte

-- Responda às seguintes perguntas:

-- O que é um JOIN em um banco de dados e para que ele é usado?
-- É usado para obter dados de várias tabelas relacionadas. Consiste na combinação de dados de uma tabela com dados de outra tabela, com base em uma ou mais condições em comum.

-- Explicar dois tipos de JOIN.
-- Inner Join é usado para obter dados relacionados de duas ou mais tabelas.
-- Left Join é usado para obter dados da tabela da esquerda mais dados relacionados da tabela da direita.

-- Para que serve o GROUP BY?
-- Agrupa os resultados de acordo com as colunas especificadas.
-- Gera um único registro para cada grupo de linhas que compartilham as colunas especificadas.
-- Reduz o número de linhas na consulta.
-- Geralmente usado em conjunto com funções de agregação, para obter dados resumidos e agrupados pelas colunas necessárias.

-- Para que é usado o HAVING?
-- A cláusula HAVING é usada para incluir condições em algumas funções SQL.
-- Isso afeta os resultados obtidos pelo Group By.

-- Segunda Parte

-- Propõe-se realizar as seguintes consultas ao banco de dados movies_db.sql trabalhado na primeira aula. Importe o arquivo movies_db.sql do PHPMyAdmin ou do MySQL Workbench e resolva as seguintes consultas:

-- Exibir o título e o nome do gênero de todas as séries.
SELECT s.title AS series_title, g.name AS genre_name FROM series s
JOIN genres g ON s.genre_id = g.id;
-- Mostre o título dos episódios, o nome e o sobrenome dos atores que trabalham em cada episódio.
SELECT e.title AS episode_title, a.first_name AS actor_first_name, a.last_name AS actor_last_name
FROM episodes e
JOIN actor_episode ae ON e.id = ae.episode_id
JOIN actors a ON ae.actor_id = a.id;
-- Mostre o título de todas as séries e o número total de temporadas de cada série.
SELECT s.title AS series_title, COUNT(t.id) AS total_seasons
FROM series s
JOIN seasons t ON s.id = t.serie_id
GROUP BY s.title;
-- Mostre o nome de todos os gêneros e o número total de filmes de cada gênero, desde que seja maior ou igual a 3.
SELECT g.name AS genre_name, COUNT(m.id) AS total_movies
FROM genres g
JOIN movies m ON g.id = m.genre_id
GROUP BY g.name
HAVING COUNT(m.id) >= 3;
-- Mostre apenas o nome e o sobrenome dos atores que trabalharam em todos os filmes de Guerra nas Estrelas e não os repita.
SELECT DISTINCT a.first_name, a.last_name
FROM actors a
JOIN actor_movie am ON a.id = am.actor_id
JOIN movies m ON am.movie_id = m.id
WHERE m.title LIKE 'La Guerra de las galaxias%';
