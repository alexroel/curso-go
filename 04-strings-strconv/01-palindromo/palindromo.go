package main

import (
	"fmt"
	"strings"
)

func reverse(cadena string) string {

	arrayCadena := strings.Split(cadena, "")
	arrayInvertida := make([]string, 0)

	for i := len(arrayCadena) - 1; i >= 0; i-- {
		arrayInvertida = append(arrayInvertida, arrayCadena[i])
	}

	return strings.Join(arrayInvertida, "")
}

func esPalindromo(palabra string) bool {
	//Oso
	palabra = strings.ToLower(palabra)
	palabra = strings.ReplaceAll(palabra, " ", "")

	palabraInvertida := reverse(palabra)

	return palabra == palabraInvertida
}

func main() {

	if esPalindromo("oso") {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}

}
