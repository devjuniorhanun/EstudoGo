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
	// Recebe os argumentos digitados pelo usuário
	entrada := os.Args[1:]
	// Cria um Slice com os argumentos
	numeros := make([]int, len(entrada))

	// Percorre o Slice
	for i, n := range entrada {
		// Converter os argumentos em Números Inteiro, ou retorna um erro
		numero, err := strconv.Atoi(n)
		// Verifica se foi gerado um erro
		if err != nil {
			// Imprimi
			fmt.Printf("%s não é um número válido!\n", n)
			// Finaliza o Programa
			os.Exit(1)
		}
		// Passa os numeros convertidos para numeros[i]
		numeros[i] = numero
	}
	// Imprimi o Método
	fmt.Println(quicksort(numeros))
}

// Método
func quicksort(numeros []int) []int {
	// Verifica se foi passado algum número
	if len(numeros) <= 1 {
		// Retorna numeros
		return numeros
	}
	// Cria uma copia do Slice original
	n := make([]int, len(numeros))
	copy(n, numeros)

	// Cria um Pivo , sendo um número que fica no meio da lista
	indicePivo := len(n) / 2
	pivo := n[indicePivo]
	// Remove o Pivo criado
	n = append(n[:indicePivo], n[indicePivo+1:]...)

	menores, maiores := particionar(n, pivo)

	return append(append(quicksort(menores), pivo), quicksort(maiores)...)
}

func particionar(numeros []int, pivo int) (menores []int, maiores []int) {
	for _, n := range numeros {
		if n <= pivo {
			menores = append(menores, n)
		} else {
			maiores = append(maiores, n)
		}
	}

	return menores, maiores
}
