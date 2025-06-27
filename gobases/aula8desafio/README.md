# desafio-go-bases

## Objetivo
O objetivo deste guia prático é consolidar e aprofundar os conceitos vistos no Go Bases. Para isso, apresentaremos um desafio integrador que nos permitirá revisar os tópicos que estudamos.

## Forma de trabalho
Uma pequena companhia aérea tem um sistema de reserva de passagens para diferentes países, que retorna um arquivo com as informações das passagens reservadas nas últimas 24 horas.

A companhia aérea precisa de um programa para extrair informações das vendas do dia para analisar as tendências de compra.

O arquivo em questão é um arquivo de valores separados por vírgula (csv), em que os campos são compostos por: id, nome, e-mail, país de destino, horário do voo e preço.

Você está pronto? 

## Desafio
Atenção! Os exemplos abaixo são apenas para orientação, o desafio pode ser resolvido de várias maneiras.

### Exercício 1: ok

Uma função que calcula quantas pessoas viajam para um determinado país:

```func GetTotalTickets(destination string) (int, error) {}```

(exemplo 1)
Dica: o VS Code nos permite pesquisar uma palavra em um arquivo com Ctrl + F ou ⌘ + F.


### Exercício 2: ok
Uma ou mais funções que calculam quantas pessoas viajam no 
- início da manhã (0 → 6)
- manhã (7 → 12)
- tarde (13 → 19)
- noite (20 → 23)

```func GetCountByPeriod(time string) (int, error) {}```

Dica: em Go, para manipular caracteres, temos o pacote strings. Se você quiser manipular datas, o pacote time

### Exercício 3:
Calcule a porcentagem de pessoas que viajam para um determinado país em um determinado dia, em relação ao restante:

```func AverageDestination(destination string, total int) (float64, error) {}```

### Exercício 4:
Crie testes de unidade para cada um dos requisitos acima, com um mínimo de 2 casos por requisito:

importar "testing".

```func TestGetTotalTickets(t *testing.T) {}```
Dica: podemos testar vários casos no mesmo teste.