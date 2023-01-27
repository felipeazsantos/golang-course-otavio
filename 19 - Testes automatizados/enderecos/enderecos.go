package enderecos

import "strings"

// TipoDeEndereco valida se o tipo é valido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "rodovia", "estrada"}

	enderecoEmLetrasMinusculas := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetrasMinusculas, " ")[0]

	enderecoDeTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoDeTipoValido = true
		}
	}

	if enderecoDeTipoValido {
		return strings.Title(primeiraPalavraDoEndereco)
	}

	return "Tipo inválido"

}
