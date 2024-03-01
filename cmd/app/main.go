package main

import (
	"fmt"
	"log"
	"meu_projeto_omie_gennera/pkg/api"
	"meu_projeto_omie_gennera/pkg/utils"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
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

	for _, id := range idPersons {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()

			idStr := strconv.Itoa(id)
			pessoasDados := api.RetornarPessoaEspecificaGennera([]string{idStr}, tokenGennera)
			if err != nil {
				log.Printf("Erro ao obter dados do Gennera para ID %d: %v", id, err)
				return
			}

			// Aqui, assumimos que pessoasDados retorna uma slice, mas você ajustará conforme sua lógica
			for _, pessoa := range pessoasDados {
				// Aqui você precisa converter pessoa (tipo PessoaDados) para ClienteOmie, se necessário
				clienteOmie := utils.ConverterPessoaParaClienteOmie(pessoa)
				bodyResponse := api.CadastrarClienteOmie(clienteOmie, appKey, appSecret)
				if bodyResponse == "" {
					log.Printf("Erro ao cadastrar cliente na Omie para ID %d: %v", id, err)
				}
			}
		}(id)
	}
	wg.Wait()
	fmt.Println("Processamento concluído para todos os IDs.")
}
