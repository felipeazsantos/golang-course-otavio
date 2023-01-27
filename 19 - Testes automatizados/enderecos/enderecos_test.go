package enderecos

import (

	"testing"
)

// TESTE DE UNIDADE

func TestTipoDeEndereco(t *testing.T) {
	enderecoParaTeste := "Rua Paulista"
	TipoDeEnderecoEsperado := "Avenida"
	TipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if TipoDeEnderecoRecebido != TipoDeEnderecoEsperado {
		t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e Recebeu %s",
			TipoDeEnderecoEsperado,
			TipoDeEnderecoRecebido)
	}
}