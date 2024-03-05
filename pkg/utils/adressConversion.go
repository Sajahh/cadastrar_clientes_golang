package utils

import (
	"regexp"
	"strings"
)

func AbreviarEndereco(endereco string) string {
	abreviacoes := map[string]string{
		"avenida": "Av.", "beco": "Bc.", "boulevard": "Blvd.", "calcadao": "Calç.",
		"caminho": "Cam.", "esplanada": "Esp.", "estrada": "Estr.", "largo": "Lg.",
		"parque": "Pq.", "passarela": "Pas.", "patio": "Pt.", "praca": "Pç.",
		"rotatoria": "Rot.", "rotunda": "Rot.", "rua": "R.", "travessa": "Tv.",
		"viaduto": "Vd.", "vila": "Vl.", "via expressa": "V. Exp.", "alameda": "Al.",
		"rodovia": "Rod.", "autoestrada": "AE", "marginal": "Marg.", "ponte": "Pte.",
		"tunel": "Tun.", "alto": "Alt.", "complexo": "Comp.", "condominio": "Cond.",
		"corredor": "Corr.", "descida": "Desc.", "subida": "Sub.",
	}

	reg, _ := regexp.Compile(`"[^\\w\\s]+`)

	palavras := strings.Fields(endereco)

	for i, palavra := range palavras {

		palavraNormalizada := reg.ReplaceAllString(strings.ToLower(palavra), "")

		if abreviacao, ok := abreviacoes[palavraNormalizada]; ok {
			palavras[i] = abreviacao
		}
	}

	return strings.Join(palavras, " ")
}
