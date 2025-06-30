-- Com base no mesmo, apresenta as consultas SQL para resolver os seguintes requisitos:

-- Enumera os dados dos autores.
SELECT * FROM AUTOR;
-- Indica o nome e a idade dos alunos
SELECT Nombre, Apellido, Edad FROM ESTUDIANTE;
-- Que alunos pertencem ao curso de informática?
SELECT Nombre, Apellido FROM ESTUDIANTE WHERE Curso = 'Informática';
-- Quais são os autores de nacionalidade francesa ou italiana?
SELECT Nombre, Apellido FROM AUTOR Where Nacionalidad IN ('Francesa', 'Italiana');
-- Quais os livros que não são da área da Internet?
SELECT Titulo FROM LIBRO WHERE Area != 'Internet';
-- Enumera os livros publicados pela Salamandra.
SELECT Titulo FROM LIBRO WHERE Editorial = 'Salamandra';
-- Enumera os nomes dos alunos cuja idade é superior à média.
SELECT Nombre, Apellido FROM ESTUDIANTE WHERE Edad > (SELECT AVG(Edad) FROM ESTUDIANTE);
-- Enumera os nomes dos alunos cujo apelido começa com a letra G.
SELECT Apellido FROM ESTUDIANTE WHERE Apellido LIKE 'G%';
-- Faz uma lista dos autores do livro "O Universo: Guia de Viagem". (Apenas os nomes devem ser indicados).
SELECT a.Nombre, a.Apellido
FROM AUTOR a
JOIN LIBROAUTOR la ON a.idAutor = la.idAutor
JOIN LIBRO l ON la.idLibro = l.idLibro
WHERE l.Titulo = 'O Universo: Guia de Viagem';
-- Que livros emprestaste ao leitor "Filippo Galli"?
SELECT l.Titulo
FROM LIBRO l
JOIN PRESTAMO p ON l.idLibro = p.idLibro
JOIN ESTUDIANTE e ON p.idLector = e.idLector
WHERE e.Nombre = 'Filippo' AND e.Apellido = 'Galli';
-- Indica o nome do aluno mais novo.
SELECT Nombre, Apellido FROM ESTUDIANTE ORDER BY Edad ASC LIMIT 1;
-- Enumera os nomes dos alunos a quem foram emprestados livros da Base de Dados.
SELECT DISTINCT e.Nombre, e.Apellido
FROM ESTUDIANTE e
JOIN PRESTAMO p ON e.idLector = p.idLector
JOIN LIBRO l ON p.idLibro = l.idLibro;
-- Enumera os livros que pertencem à autora J.K. Rowling.
SELECT l.Titulo
FROM LIBRO l
JOIN LIBROAUTOR la ON l.idLibro = la.idLibro
JOIN AUTOR a ON la.idAutor = a.idAutor
WHERE a.Nombre = 'J.K.' AND a.Apellido = 'Rowling';
-- Enumera os títulos dos livros que deviam ser devolvidos em 16/07/2021.
SELECT l.Titulo
FROM LIBRO l
JOIN PRESTAMO p ON l.idLibro = p.idLibro
WHERE p.FechaDevolucion = '2021-07-16';