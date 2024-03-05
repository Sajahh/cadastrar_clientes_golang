package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"meu_projeto_omie_gennera/pkg/models"
	"net/http"
)

func CadastrarClienteOmie(cliente models.ClienteOmie, appKey string, appSecret string, idPerson string) (string, error) {
	url := "https://app.omie.com.br/api/v1/geral/clientes/"
	payload := models.ClientePayload{
		Call:      "IncluirCliente",
		AppKey:    appKey,
		AppSecret: appSecret,
		Param:     []models.ClienteOmie{cliente},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("erro ao serializar o payload: %w", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", fmt.Errorf("erro ao fazer a requisição: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return string(body), fmt.Errorf("erro com ID %s - status da resposta: %d, corpo: %s", idPerson, resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("erro ao ler a resposta para ID %s: %w", idPerson, err)
	}

	return string(body), nil
}
