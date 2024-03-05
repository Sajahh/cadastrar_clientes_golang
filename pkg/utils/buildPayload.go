package utils

import (
	"fmt"
	"meu_projeto_omie_gennera/pkg/models" // Assegure-se de que este é o caminho correto
)

func PrepararClienteParaOmie(cliente *models.ClienteOmie) error {

	if len(cliente.Estado) != 2 {
		sigla, ok := EstadoParaSigla(cliente.Estado)
		if !ok {
			return fmt.Errorf("estado não reconhecido: %s", cliente.Estado)
		}
		cliente.Estado = sigla
	}

	cliente.Endereco = AbreviarEndereco(cliente.Endereco)

	return nil
}
