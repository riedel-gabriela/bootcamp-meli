package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	tickets.LoadDataTickets("tickets.csv")
	total, _ := tickets.GetTotalTickets("Brazil")
	fmt.Println("total:", total)

	totalPeriodo, _ := tickets.GetCountByPeriod("início da manhã")
	fmt.Println("total:", totalPeriodo)

	totalPercent, _ := tickets.AverageDestination("Brazil")
	fmt.Println("total:", totalPercent, "%")
}
