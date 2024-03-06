package main

import (
	"fmt"
	"log"
	"meu_projeto_omie_gennera/pkg/api"
	"meu_projeto_omie_gennera/pkg/utils"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Erro ao carregar o arquivo .env")
	}

	tokenGennera := os.Getenv("GENNERA_TOKEN")
	if tokenGennera == "" {
		log.Fatal("GENNERA_TOKEN não definido no .env")
	}
	appKey := os.Getenv("APP_KEY")
	if appKey == "" {
		log.Fatal("APP_KEY não definido no .env")
	}
	appSecret := os.Getenv("APP_SECRET")
	if appSecret == "" {
		log.Fatal("APP_SECRET não definido no .env")
	}
	fmt.Println("Iniciando o aplicativo...")

	idPersonsStr := os.Getenv("ID_PERSON")
	idPersons := strings.Split(idPersonsStr, ",")

	var idsComErro []string

	for _, idStr := range idPersons {
		pessoasDados, err := api.RetornarPessoaEspecificaGennera([]string{idStr}, tokenGennera)
		if err != nil {
			log.Printf("Erro ao obter dados do Gennera para ID %s: %v", idStr, err)
			idsComErro = append(idsComErro, idStr)
			continue
		}

		for _, pessoa := range pessoasDados {
			// Adicionando um delay entre cada requisição pode ajudar a evitar atingir limites de taxa
			time.Sleep(1 * time.Second)
			clienteOmie := utils.ConverterPessoaParaClienteOmie(pessoa)
			err := utils.PrepararClienteParaOmie(&clienteOmie)
			if err != nil {
				log.Printf("Erro ao preparar cliente para Omie para ID %s: %v", idStr, err)
				continue
			}
			bodyResponse, err := api.CadastrarClienteOmie(clienteOmie, appKey, appSecret, idStr)
			if err != nil || bodyResponse == "" {
				log.Printf("Erro ao cadastrar cliente na Omie para ID %s: %v", idStr, err)
				idsComErro = append(idsComErro, idStr)
			} else {
				log.Printf("Cliente cadastrado com sucesso na Omie para o ID %s", idStr)
			}
			// Adiciona outro delay após a requisição, se necessário.
			time.Sleep(1 * time.Second)
		}
	}

	if len(idsComErro) > 0 {
		utils.SalvarIDsComErro(idsComErro)
	}

	fmt.Println("Programa finalizado.")
}
