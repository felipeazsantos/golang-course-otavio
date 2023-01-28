package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome string
	Email string
}

func main() {

	// Joga todos os templates HTML pra dentro da variavel
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func (w http.ResponseWriter, r *http.Request) {
		u := usuario{"Pedro", "pedro@gmail.com"}

		// renderiza um arquivo html que está dentro da variável templates
		// params: responsewriter, nome do template, dados para passar ao template
		templates.ExecuteTemplate(w, "home.html", u)
	})

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))

}
