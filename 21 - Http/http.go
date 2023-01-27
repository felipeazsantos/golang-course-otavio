package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá mundo!"))
}

func main() {

	// HTTṔ É UM PROTOCOLO DE COMUNICAÇÃO - BASE DE COMUNICAÇÃO WEB

	// CLIENTE (FAZ REQUISIÇÃO) - SERVIDOR (PROCESSA REQUISIÇÃO E ENVIA RESPOSTA)

	// Request - Response

	// Rotas
	// URI -> identiicador do recurso
	// métodos -> GET, POST, PUT, DELETE

	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Carregar páginas de usuários"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
