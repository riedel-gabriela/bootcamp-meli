package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

/*
Exercício 2 - Manipulação do Body

Vamos criar um endpoint chamado /greetings. Com uma pequena estrutura com nome e sobrenome que, quando colada, deve responder no texto "Hello + nome + sobrenome".

O endpoint deve ser um método POST
O package JSON deve ser usado para resolver o exercício.
A resposta deve seguir a seguinte estrutura: "Hello Andrea Rivas".
A estrutura deve ter a seguinte aparência:

{
    "firstName": "Andrea",
    "lastName": "Rivas".
}
*/

type GreetingRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func greetingsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// inicia uma variavel com a estrutura do request
	var greeting GreetingRequest
	err := json.NewDecoder(r.Body).Decode(&greeting)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response := fmt.Sprintf("Hello %s %s", greeting.FirstName, greeting.LastName)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, err = w.Write([]byte(response))

	if err != nil {
		log.Printf("Error writing response for %s: %v\n", r.URL.Path, err)
	}
}

func main() {
	http.HandleFunc("/greetings", greetingsHandler)
	fmt.Println("Server running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
