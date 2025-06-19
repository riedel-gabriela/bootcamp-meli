package loader

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/riedel-gabriela/desafio-goweb/internal"
)

// NewLoaderTicketCSV creates a new ticket loader from a CSV file
func NewLoaderTicketCSV(filePath string) *LoaderTicketCSV {
	return &LoaderTicketCSV{
		filePath: filePath,
	}
}

// LoaderTicketCSV represents a ticket loader from a CSV file
type LoaderTicketCSV struct {
	filePath string
}

// Load loads the tickets from the CSV file
func (t *LoaderTicketCSV) Load() (l map[int]internal.TicketAttributes, err error) {
	// open the file
	f, err := os.Open(t.filePath)
	if err != nil {
		err = fmt.Errorf("error opening file: %v", err)
		return
	}
	defer f.Close()

	// read the file
	r := csv.NewReader(f)

	// read the records
	l = make(map[int]internal.TicketAttributes)
	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			err = fmt.Errorf("error reading record: %v", err)
			return nil, err
		}

		id := record[0]
		price, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing price for record %v: %v", record, err)
		}
		ticket := internal.TicketAttributes{
			Name:    string(record[1]),
			Email:   string(record[2]),
			Country: string(record[3]),
			Hour:    string(record[4]),
			Price:   price,
		}

		// add the ticket to the map
		intID, err := strconv.Atoi(id)
		if err != nil {
			return nil, fmt.Errorf("error parsing id for record %v: %v", record, err)
		}
		l[intID] = ticket
	}

	return
}
