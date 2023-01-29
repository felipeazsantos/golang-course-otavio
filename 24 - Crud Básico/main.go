package main

import (
	"crud/servidor"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// CRUD - CREATE, RETRIVE, UPDATE, DELETE

	// CREATE - POST
	// RETRIVE - GET
	// UPDATE - PUT
	// DELETE - DELETE

	router := mux.NewRouter()
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", servidor.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.BuscarUsuario).Methods(http.MethodGet)

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
