package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

/*
O mesmo estudo do exercício anterior solicita uma funcionalidade para poder registrar novos dados de clientes. Os dados necessários são:

File
Name
ID
Phone number
Adress

Atividade 1: Antes de registrar um cliente, é necessário verificar se o cliente já existe. Para fazer isso, você precisa ler os dados de uma matriz. Caso ele se repita, você precisa tratar o erro adequadamente, como vimos até agora. Esse erro deve:
1.- gerar um panic;
2.- console iniciar a mensagem: “Error: client already exists”, e continuar com a execução do programa normalmente.

Atividade 2: Depois de tentar verificar se o cliente a ser registrado já existe, desenvolva uma função para validar se todos os dados a serem registrados para um cliente contêm um valor diferente de zero. Essa função deve retornar pelo menos dois valores. Um deles deverá ser do tipo erro, caso um valor zero seja inserido como parâmetro (lembre-se dos valores zero de cada tipo de dado, por exemplo: 0, “”, nil).

Atividade 3: Antes de encerrar a execução, mesmo que ocorram panics, as seguintes mensagens devem ser impressas no console: “End of execution” e “Several errors were detected at runtime”. Use o defer para atender a esse requisito..

Requisitos gerais:

Use a recover para recuperar o valor de qualquer pânico que possa ocorrer.
Lembre-se de realizar as validações necessárias para cada retorno que possa conter um valor de erro.
Gere um erro, personalizando-o de acordo com sua preferência usando uma das funções Go (execute também a validação relevante para o caso de erro retornado).

*/

// --- Client Structure ---
type Client struct {
	File        string
	Name        string
	ID          string
	PhoneNumber string
	Address     string
}

type ClientAlreadyExistsError struct {
	ClientID string
}

func (e *ClientAlreadyExistsError) Error() string {
	return fmt.Sprintf("Error: client with ID '%s' already exists.", e.ClientID)
}

type DataValidationError struct {
	FieldName string
	Message   string
}

func (e *DataValidationError) Error() string {
	return fmt.Sprintf("Validation error for field '%s': %s", e.FieldName, e.Message)
}

var existingClients = []Client{
	{File: "doc1", Name: "Alice", ID: "A001", PhoneNumber: "11987654321", Address: "Street 1"},
	{File: "doc2", Name: "Bob", ID: "B002", PhoneNumber: "21912345678", Address: "Avenue 2"},
}

func validateExistingClient(c Client) error {
	for _, client := range existingClients {
		if client.ID == c.ID {
			return &ClientAlreadyExistsError{ClientID: client.ID}
		}
	}
	return nil
}

func validateClientData(c Client) error {
	if c.File == "" {
		return &DataValidationError{FieldName: "File", Message: "cannot be empty"}
	}
	if c.Name == "" {
		return &DataValidationError{FieldName: "Name", Message: "cannot be empty"}
	}
	if c.ID == "" {
		return &DataValidationError{FieldName: "ID", Message: "cannot be empty"}
	}
	if c.PhoneNumber == "" {
		return &DataValidationError{FieldName: "Phone Number", Message: "cannot be empty"}
	}
	if c.Address == "" {
		return &DataValidationError{FieldName: "Address", Message: "cannot be empty"}
	}
	return nil
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("\n-------------------------------------------------")
			fmt.Printf("Recovered from panic: %v\n", r)

			if customErr, ok := r.(error); ok {
				var clientExistsErr *ClientAlreadyExistsError
				if errors.As(customErr, &clientExistsErr) {
					fmt.Printf("Error: client already exists (%s). Continuing program normally.\n", clientExistsErr.ClientID)
				} else {
					fmt.Printf("An unexpected error occurred during panic: %v\n", customErr)
				}
			} else {
				fmt.Printf("An unexpected non-error panic occurred: %v\n", r)
			}
			fmt.Println("-------------------------------------------------")
		}
		fmt.Println("End of execution")
		fmt.Println("Several errors were detected at runtime (refer to logs/output)")
	}()

	fmt.Println("Starting client processing from file...")

	path := "example.txt"

	leitura, err := os.Open(path)
	if err != nil {
		panic("error: The indicated file was not found or is damaged.")
	}
	defer leitura.Close()

	scanner := bufio.NewScanner(leitura)
	lineNumber := 0 // para melhorar o log de erros no console
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		fmt.Printf("Processing line %d: '%s'\n", lineNumber, line)

		line_slice := strings.Split(line, ",")
		if len(line_slice) != 5 { // checando se a linha tem o número necessário de dados
			fmt.Printf("Skipping line %d: Incorrect number of fields. Expected 5, got %d.\n", lineNumber, len(line_slice))
			continue
		}

		client := Client{
			File:        line_slice[0],
			Name:        line_slice[1],
			ID:          line_slice[2],
			PhoneNumber: line_slice[3],
			Address:     line_slice[4],
		}
		validateErr := validateClientData(client)
		if validateErr != nil {
			fmt.Printf("Client data validation error on line %d (%s): %v\n", lineNumber, client.ID, validateErr)
			continue
		}
		existsErr := validateExistingClient(client)
		if existsErr != nil {
			panic(existsErr)
		}
		existingClients = append(existingClients, client)
		fmt.Printf("Client '%s' (ID: %s) added to in-memory list.\n", client.Name, client.ID)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error encountered while reading file %s: %v\n", path, err)
	}
	fmt.Printf("Final client list: %+v\n", existingClients)
	fmt.Println("Main execution completed.")
}
