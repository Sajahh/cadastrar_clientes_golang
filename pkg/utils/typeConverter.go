package utils

import (
	"meu_projeto_omie_gennera/pkg/models"
	"strconv"
)

func ConverterPessoaParaClienteOmie(pessoa models.PessoaDados) models.ClienteOmie {
	return models.ClienteOmie{
		CodigoClienteIntegracao: strconv.Itoa(pessoa.IDPerson), // Convertendo int para string
		CodigoClienteOmie:       "",                            // Este campo pode precisar ser gerado ou vindo de outra fonte, já que não parece ter um equivalente direto
		Email:                   pessoa.Email,
		RazaoSocial:             pessoa.Name, // Assumindo que Name é a razão social
		CnpjCpf:                 pessoa.CPF,
		Endereco:                pessoa.Street,
		EnderecoNumero:          pessoa.StreetNumber,
		Complemento:             "", // Sem equivalente direto, deixado como vazio
		Bairro:                  "", // Sem equivalente direto, deixado como vazio
		Estado:                  pessoa.State,
		Cidade:                  pessoa.City,
		Cep:                     pessoa.Zipcode,
	}
}
