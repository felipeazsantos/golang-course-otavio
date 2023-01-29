package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type usuario struct {
	ID uint32 `json:"id"`
	Nome string `json:"nome"`
	Email string `json:"email"`
}

//CriaruUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição."))
		return
	}

	var u usuario

	if erro = json.Unmarshal(corpoRequisicao, &u); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct."))
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		w.Write([]byte("Falha ao conectar com o banco de dados."))
		return
	}
	defer db.Close()

	// INSERT INTO usuarios (nome, email) values ('pedro', 'pedro@gmail.com') -> No MySQL

	statement, erro := db.Prepare("INSERT INTO usuarios (nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement."))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(u.Nome, u.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar comando de inserção do usuário no banco de dados."))
		return
	}

	IdInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o id inserido!"))
		return
	}

	// Status Code
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", IdInserido)))

}
