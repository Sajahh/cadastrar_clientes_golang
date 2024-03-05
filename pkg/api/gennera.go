package api

import (
	"encoding/json"
	"fmt"
	"io"
	"meu_projeto_omie_gennera/pkg/models"
	"net/http"
)

func RetornarPessoaEspecificaGennera(listaIDPerson []string, tokenGennera string) ([]models.PessoaDados, error) {
	var pessoasDados []models.PessoaDados

	for _, idPerson := range listaIDPerson {
		url := fmt.Sprintf("https://api2.gennera.com.br/institutions/696/persons/%s", idPerson)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, fmt.Errorf("erro ao criar a requisição para o ID %s: %v", idPerson, err)
		}
		req.Header.Add("x-access-token", tokenGennera)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("erro ao fazer a requisição para o ID %s: %v", idPerson, err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			responseBody, _ := io.ReadAll(resp.Body)
			return nil, fmt.Errorf("erro ao fazer a requisição para o ID %s: status %d, resposta: %s", idPerson, resp.StatusCode, string(responseBody))
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("erro ao ler o corpo da resposta para o ID %s: %v", idPerson, err)
		}

		var pessoa models.PessoaDados
		if err := json.Unmarshal(bodyBytes, &pessoa); err != nil {
			return nil, fmt.Errorf("erro ao fazer o Unmarshal da resposta para o ID %s: %v", idPerson, err)
		}

		pessoasDados = append(pessoasDados, pessoa)
	}
	return pessoasDados, nil
}
