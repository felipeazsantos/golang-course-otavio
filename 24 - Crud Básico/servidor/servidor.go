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

// CriaruUsuario insere um usuário no banco de dados
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

// BuscarUsuarios retorna todos os usuários do banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("select * from usuarios")

	if erro != nil {
		w.Write([]byte("Erro ao buscar os usuários no banco de dados!"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao buscar ao escanear o usuário"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter usuários para JSON"))
		return
	}
}

//BuscarUsuario trás um usuário específico do banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

}