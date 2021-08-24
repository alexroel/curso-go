package main

import "fmt"

func main() {

	numeros := make([]int, 3, 3)

	numeros[0] = 100
	numeros[1] = 200
	numeros[2] = 300

	numeros = append(numeros, 400)

	fmt.Println(numeros, len(numeros), cap(numeros))

}
