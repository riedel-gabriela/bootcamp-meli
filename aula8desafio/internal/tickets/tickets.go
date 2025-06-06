package tickets

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type Ticket struct {
	ID      int
	Nome    string
	Email   string
	Destino string
	Hora    string
	Preco   float64
}

var tickets map[int]*Ticket

type CSVDataError struct {
	Line    int
	Message string
}

func (e *CSVDataError) Error() string {
	return fmt.Sprintf("Erro na linha %d do CSV: %s", e.Line, e.Message)
}

func LoadDataTickets(path string) error {
	file, err := os.Open(path)
	if err != nil {
		fmt.Errorf("Erro ao abrir o arquivo CSV '%s': %v", path, err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	lineNumber := 0
	tickets = make(map[int]*Ticket)
	for {
		record, err := csvReader.Read()
		lineNumber++

		if err == io.EOF {
			break
		}
		if err != nil {
			panic("Aviso: Erro fatal de leitura na linha %d: %v. Interrompendo.\n")
		}

		if len(record) != 6 {
			fmt.Printf("Aviso: %s. Pulando registro.\n",
				(&CSVDataError{Line: lineNumber, Message: fmt.Sprintf("Número de colunas incorreto. Esperado 6, encontrado %d", len(record))}).Error())
			continue
		}

		preco, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			fmt.Printf("Aviso: Erro ao converter preço '%s' para int. Pulando registro: %v\n", record[0], err)
			continue
		}
		id, err := strconv.Atoi(record[0])
		if err != nil {
			fmt.Printf("Aviso: Erro ao converter ID '%s' para int. Pulando registro: %v\n", record[0], err)
			continue
		}

		ticket := Ticket{
			ID:      id,
			Nome:    record[1],
			Email:   record[2],
			Destino: record[3],
			Hora:    record[4],
			Preco:   preco,
		}
		tickets[id] = &ticket
	}
	return nil
}

// ejemplo 1
func GetTotalTickets(destination string) (int, error) {
	total := 0
	for _, ticket := range tickets {
		if ticket.Destino == destination {
			total++
		}
	}
	if total == 0 {
		return 0, errors.New("cannot find destination")
	}
	return total, nil
}

// ejemplo 2
// - início da manhã (0 → 6)
// - manhã (7 → 12)
// - tarde (13 → 19)
// - noite (20 → 23)

const (
	inicio = "início da manhã"
	manha  = "manhã"
	tarde  = "tarde"
	noite  = "noite"
)

func GetCountByPeriod(period string) (int, error) {
	total := 0
	cutOffHour := 0
	cutOffMin := 59
	switch period {
	case inicio:
		cutOffHour = 6
	case manha:
		cutOffHour = 12
	case tarde:
		cutOffHour = 19
	case noite:
		cutOffHour = 23
	default:
		return 0, errors.New("period not found")
	}

	for _, ticket := range tickets {
		t, err := time.Parse("15:04", ticket.Hora)
		if err != nil {
			fmt.Println("Erro ao analisar hora")
			continue
		}

		if t.Hour() < cutOffHour || (t.Hour() == cutOffHour && t.Minute() == cutOffMin) {
			total++
		}
	}
	return total, nil
}

// ejemplo 3
func AverageDestination(destination string) (int, error) {
	totalDestination := 0
	total := len(tickets)
	if total == 0 {
		return 0, errors.New("no tickets loaded")
	}
	for _, ticket := range tickets {
		if ticket.Destino == destination {
			totalDestination++
		}
	}
	return (totalDestination * 100) / total, nil
}
