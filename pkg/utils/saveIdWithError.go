package utils

import (
	"log"
	"os"
)

func SalvarIDsComErro(ids []string) {
	file, err := os.Create("ids_com_erro.txt")
	if err != nil {
		log.Fatalf("Erro ao criar arquivo para salvar IDs com erro: %v", err)
	}
	defer file.Close()

	for _, id := range ids {
		_, err := file.WriteString(id + "\n")
		if err != nil {
			log.Printf("Erro ao escrever ID no arquivo: %v", err)
		}
	}
}
