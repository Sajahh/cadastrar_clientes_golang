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

// REVIEW - Testar se ID esta sendo armazenado
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
	var wg sync.WaitGroup

	var idsComErro []string
	var mutex sync.Mutex

	for _, idStr := range idPersons {
		wg.Add(1)
		go func(idStr string) {
			defer wg.Done()

			pessoasDados, err := api.RetornarPessoaEspecificaGennera([]string{idStr}, tokenGennera)
			if err != nil {
				log.Printf("Erro ao obter dados do Gennera para ID %s: %v", idStr, err)
				mutex.Lock()
				idsComErro = append(idsComErro, idStr)
				mutex.Unlock()
				return
			}

			for _, pessoa := range pessoasDados {
				clienteOmie := utils.ConverterPessoaParaClienteOmie(pessoa)
				//REVIEW - Testar abreviacao de endereco
				err := utils.PrepararClienteParaOmie(&clienteOmie)
				if err != nil {
					log.Printf("Erro ao preparar cliente para Omie para ID %s: %v", idStr, err)
				}
				bodyResponse, err := api.CadastrarClienteOmie(clienteOmie, appKey, appSecret, idStr)
				if err != nil || bodyResponse == "" {
					log.Printf("Erro ao cadastrar cliente na Omie para ID %s: %v", idStr, err)
					mutex.Lock()
					idsComErro = append(idsComErro, idStr)
					mutex.Unlock()
				} else {
					log.Printf("Cliente cadastrado com sucesso na Omie para o ID %s", idStr)
				}
			}
		}(idStr)
	}
	wg.Wait()

	if len(idsComErro) > 0 {
		utils.SalvarIDsComErro(idsComErro)
	}

	fmt.Println("Programa finalizado.")
}
