package main

import (
	"fmt"
	"log"
	"meu_projeto_omie_gennera/pkg/api"
	"meu_projeto_omie_gennera/pkg/utils"
	"os"
	"strings"
	"sync"

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

	idPersonsStr := os.Getenv("ID_PERSONS")
	idPersons := strings.Split(idPersonsStr, ",")
	var wg sync.WaitGroup

	for _, idStr := range idPersons {
		wg.Add(1)
		go func(idStr string) {
			defer wg.Done()

			pessoasDados := api.RetornarPessoaEspecificaGennera([]string{idStr}, tokenGennera)
			if err != nil {
				log.Printf("Erro ao obter dados do Gennera para ID %s: %v", idStr, err)
				return
			}

			for _, pessoa := range pessoasDados {
				clienteOmie := utils.ConverterPessoaParaClienteOmie(pessoa)
				bodyResponse, err := api.CadastrarClienteOmie(clienteOmie, appKey, appSecret)
				if err != nil || bodyResponse == "" {
					log.Printf("Erro ao cadastrar cliente na Omie para ID %s: %v", idStr, err)
				}
			}
		}(idStr)
	}
	wg.Wait()
	fmt.Println("Processamento concluído para todos os IDs.")
}
