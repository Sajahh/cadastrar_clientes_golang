package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"meu_projeto_omie_gennera/pkg/api"
	"os"
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
	fmt.Println("Iniciando o aplicativo...")

	listaIDPerson := []string{"2156551", "2484803"}
	pessoasDados := api.RetornarPessoaEspecificaGennera(listaIDPerson, tokenGennera)
	fmt.Println(pessoasDados)
}
