/*
Exercício 1 - Teste de ping
Vamos criar um aplicativo da Web com o package net/http nativo do Go, que tem um endpoint /ping que, quando colado, responde com um texto que diz "pong".

O endpoint deve ser um método GET
A resposta do pong deve ser enviada como texto, NÃO como JSON.
*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, err := w.Write([]byte("pong"))

	if err != nil {
		log.Printf("Erro ao escrever resposta para %s: %v\n", r.URL.Path, err)
	}
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	fmt.Println("Servidor rodando na porta 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
