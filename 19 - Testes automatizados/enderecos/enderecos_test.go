package enderecos

import (

	"testing"
)

// TESTE DE UNIDADE

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado string
}

func TestTipoDeEndereco(t *testing.T) {

	cenarios := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça das rosas", "Tipo inválido"},
		{"Estrada dos alvarengas", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"", "Tipo inválido"},
	}

	//if TipoDeEnderecoRecebido != TipoDeEnderecoEsperado {
	//	t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e Recebeu %s",
	//		TipoDeEnderecoEsperado,
	//		TipoDeEnderecoRecebido)
	//}

	for _, cenario := range cenarios {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)

		if (retornoRecebido != cenario.retornoEsperado) {
			t.Errorf("Valor recebido diferente do valor esperado! Recebido %s | Esperado %s",
				retornoRecebido,
				cenario.retornoEsperado)
		}
	}
}