package utils

import (
	"meu_projeto_omie_gennera/pkg/models"
	"strconv"
)

func ConverterPessoaParaClienteOmie(pessoa models.PessoaDados) models.ClienteOmie {
	return models.ClienteOmie{
		CodigoClienteIntegracao: pessoa.CPF,
		CodigoClienteOmie:       strconv.Itoa(pessoa.IDPerson),
		Email:                   pessoa.Email,
		RazaoSocial:             pessoa.Name,
		CnpjCpf:                 pessoa.CPF,
		Endereco:                pessoa.Street,
		EnderecoNumero:          pessoa.StreetNumber,
		Complemento:             "",
		Bairro:                  pessoa.Neighborhood,
		Estado:                  pessoa.State,
		Cidade:                  pessoa.City,
		Cep:                     pessoa.Zipcode,
	}
}
