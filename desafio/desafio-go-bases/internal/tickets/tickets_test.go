package tickets

import (
	"os"
	"strings"
	"testing"
)

func createTestCSV(t *testing.T, filename string, content string) {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test CSV file %s: %v", filename, err)
	}
}

func TestLoadDataTickets(t *testing.T) {
	t.Run("Valid CSV should load tickets", func(t *testing.T) {
		testCSVContent := `1,Alice,alice@email.com,New York,08:00,150
                            2,Bob,bob@email.com,London,14:30,200
                            3,Charlie,charlie@email.com,New York,11:59,160
                            `
		csvFileName := "test_valid.csv"
		createTestCSV(t, csvFileName, testCSVContent)
		defer os.Remove(csvFileName)
		err := LoadDataTickets(csvFileName)
		if err != nil {
			t.Errorf("LoadDataTickets failed unexpectedly for valid CSV: %v", err)
		}
		if len(tickets) != 3 {
			t.Errorf("Expected 3 tickets to be loaded, got %d", len(tickets))
		}
		if _, found := tickets[1]; !found {
			t.Errorf("Expected ticket with ID 1 to be loaded")
		}
		if tickets[1].Nome != "Alice" {
			t.Errorf("Expected ticket ID 1 to have Name 'Alice', got '%s'", tickets[1].Nome)
		}
	})
	t.Run("Invalid CSV should handle errors and skip records", func(t *testing.T) {
		testCSVContent := `10,InvalidName,,Dest,09:00,50
                            12,BadPrice,email@test.com,Dest,10:00,ABC
                            13,DuplicateID,dup@email.com,DupDest,11:00,100
                            13,DuplicateID2,dup2@email.com,DupDest2,12:00,120
                            14,ValidOne,valid@test.com,Dest,13:00,200
                            `
		csvFileName := "test_invalid.csv"
		createTestCSV(t, csvFileName, testCSVContent)
		defer os.Remove(csvFileName)

		tickets = make(map[int]*Ticket)
		err := LoadDataTickets(csvFileName)
		if err != nil {
			t.Errorf("LoadDataTickets returned an unexpected error for invalid CSV: %v", err)
		}
		if len(tickets) != 2 {
			t.Errorf("Expected 2 valid tickets to be loaded, got %d. Map: %+v", len(tickets), tickets)
		}
		if _, found := tickets[13]; !found {
			t.Errorf("Expected ticket with ID 13 to be loaded")
		}
		if _, found := tickets[14]; !found {
			t.Errorf("Expected ticket with ID 14 to be loaded")
		}
	})
	t.Run("Non-existent CSV file should return error", func(t *testing.T) {
		nonExistentFile := "non_existent.csv"
		err := LoadDataTickets(nonExistentFile)
		if err == nil {
			t.Error("Expected error for non-existent file, but got nil")
		}
		if !strings.Contains(err.Error(), "no such file or directory") &&
			!strings.Contains(err.Error(), "The system cannot find the file specified") {
			t.Errorf("Unexpected error message for non-existent file: %v", err)
		}
	})
}

func TestGetTotalTickets(t *testing.T) {
	tickets = map[int]*Ticket{
		1: {ID: 1, Destino: "Brazil"},
		2: {ID: 2, Destino: "Brazil"},
		3: {ID: 3, Destino: "Argentina"},
	}
	t.Run("Should return correct count for existing destination", func(t *testing.T) {
		count, err := GetTotalTickets("Brazil")
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if count != 2 {
			t.Errorf("Expected 2 tickets for Brazil, got %d", count)
		}
	})
	t.Run("Should return error for non-existent destination", func(t *testing.T) {
		count, err := GetTotalTickets("Chile")
		if err == nil {
			t.Error("Expected error for non-existent destination, got nil")
		}
		if count != 0 {
			t.Errorf("Expected 0 tickets for Chile, got %d", count)
		}
	})
}

func TestGetCountByPeriod(t *testing.T) {
	tickets = map[int]*Ticket{
		1: {ID: 1, Hora: "05:30"},
		2: {ID: 2, Hora: "08:00"},
		3: {ID: 3, Hora: "14:00"},
		4: {ID: 4, Hora: "21:00"},
		5: {ID: 5, Hora: "23:59"},
	}
	t.Run("Should count tickets for 'início da manhã'", func(t *testing.T) {
		count, err := GetCountByPeriod(inicio)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if count != 1 {
			t.Errorf("Expected 1 ticket for início da manhã, got %d", count)
		}
	})
	t.Run("Should count tickets for 'noite'", func(t *testing.T) {
		count, err := GetCountByPeriod(noite)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if count != 5 {
			t.Errorf("Expected 5 tickets for noite, got %d", count)
		}
	})
	t.Run("Should return error for invalid period", func(t *testing.T) {
		_, err := GetCountByPeriod("madrugada")
		if err == nil {
			t.Error("Expected error for invalid period, got nil")
		}
	})
}

func TestAverageDestination(t *testing.T) {
	tickets = map[int]*Ticket{
		1: {ID: 1, Destino: "Brazil"},
		2: {ID: 2, Destino: "Brazil"},
		3: {ID: 3, Destino: "Argentina"},
		4: {ID: 4, Destino: "Brazil"},
	}
	t.Run("Should return correct average for destination", func(t *testing.T) {
		avg, err := AverageDestination("Brazil")
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if avg != 75 {
			t.Errorf("Expected 75%% for Brazil, got %d", avg)
		}
	})
	t.Run("Should return 0 for destination not present", func(t *testing.T) {
		avg, err := AverageDestination("Chile")
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if avg != 0 {
			t.Errorf("Expected 0%% for Chile, got %d", avg)
		}
	})
	t.Run("Should return error if no tickets loaded", func(t *testing.T) {
		tickets = map[int]*Ticket{}
		_, err := AverageDestination("Brazil")
		if err == nil {
			t.Error("Expected error when no tickets loaded, got nil")
		}
	})
}
