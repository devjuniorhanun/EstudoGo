// Programas executáveis iniciam pelo pacote main
package main

/*
Os códigos em Go são organizados em pacotes
e para usá-los é necessário declarar um ou vários imports
*/
import (
	"fmt"     // Pacote de impressão
	"os"      // Pacote de Sistema. Recupera argumentos passado por linha de comanado
	"strconv" // Pacote de conversão de String em outros tipos
)

// A porta de entrada de um programa Go é a função main
func main() {
	// Verifica a se foi passado 3 argumentos
	if len(os.Args) < 3 {
		// Imprime
		fmt.Println("Uso: conversor <valores> <unidade>")
		// Finalizada o programa
		os.Exit(1)
	}

	// Recebe o primeiro argumento digitado pelo usuário
	valoresOrigem := os.Args[1 : len(os.Args)-1]
	// Recebe o segundo argumento digitado pelo usuário
	unidadeOrigem := os.Args[len(os.Args)-1]

	// Instancida uma variável string
	var unidadeDestino string

	// Verifica se foi digitado celsius
	if unidadeOrigem == "celsius" {
		// Muda o valor da unidadeDestino
		unidadeDestino = "fahrenheit"
		// Verifica se foi passado quilometros
	} else if unidadeOrigem == "quilometros" {
		unidadeDestino = "milhas"
		// Se não foi passado um valor valido
	} else {
		fmt.Printf("%s não é uma unidade conhecida!", unidadeOrigem)
		os.Exit(1)
	}

	// Percorre o Slice
	for i, v := range valoresOrigem {
		// Converte string em Ponto flutuante, ou gera um erro
		valorOrigem, err := strconv.ParseFloat(v, 64)
		// Verificar se foi gerando um erro
		if err != nil {
			// Imprime
			fmt.Printf("O valor %s na posição %d não é um número válido!\n", v, i)
			// Finalize o programa
			os.Exit(1)
		}

		// Cria ma variável em ponto flutuante
		var valorDestino float64

		// Verifica se foi passado celsius
		if unidadeOrigem == "celsius" {
			// Efetua a operação
			valorDestino = valorOrigem*1.8 + 32
		} else {
			// Efetua a operação
			valorDestino = valorOrigem / 1.60934
		}

		// Imprime o resultado
		fmt.Printf("%.2f %s = %.2f %s\n",
			valorOrigem, unidadeOrigem, valorDestino, unidadeDestino)
	}
}
