package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"meu_projeto_omie_gennera/pkg/models"
	"net/http"
)

func RetornarPessoaEspecificaGennera(listaIDPerson []string, tokenGennera string) []models.PessoaDados {
	var pessoasDados []models.PessoaDados

	for _, idPerson := range listaIDPerson {

		url := fmt.Sprintf("https://api2.gennera.com.br/institutions/696/persons/%s", idPerson)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Printf("Erro ao criar a requisição para o ID %s: %v", idPerson, err)
			continue
		}
		req.Header.Add("x-access-token", tokenGennera)

		client := &http.Client{}
		resp, err := client.Do((req))
		if err != nil {
			log.Printf("Erro ao fazer a requisição para o ID %s: %v", idPerson, err)
			continue
		}

		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			log.Printf("Erro ao fazer a requisição para o ID %s: %v", idPerson, err)
			continue
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Erro ao ler o corpo da resposta para o ID %s: %v", idPerson, err)
			continue
		}
		var pessoa models.PessoaDados
		if err := json.Unmarshal(bodyBytes, &pessoa); err != nil {
			log.Printf("Erro ao fazer o Unmarshal da resposta para o ID %s: %v", idPerson, err)
			continue
		}

		pessoasDados = append(pessoasDados, pessoa)
	}
	return pessoasDados
}
