package api

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"meu_projeto_omie_gennera/pkg/models"
	"net/http"
)

func CadastrarClienteOmie(cliente models.ClienteOmie, appKey string, appSecret string) string {

	url := "https://app.omie.com.br/api/v1/geral/clientes/"
	payload := models.ClientePayload{
		Call:      "IncluirCliente",
		AppKey:    appKey,
		AppSecret: appSecret,
		Param:     []models.ClienteOmie{cliente},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Erro ao serializar o payload: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatalf("Erro ao fazer a requisição: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Erro ao ler a resposta: %v", err)
	}

	bodyStr := string(body)
	return bodyStr
}
