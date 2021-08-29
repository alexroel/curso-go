package main

import "fmt"

//Función variádica -
func sumar(nombre string, numeros ...int) (string, int) {

	mensaje := fmt.Sprintf("La suma de %s es: ", nombre)
	var total int
	for _, num := range numeros {
		total += num
	}

	//Retornar multiples valores
	return mensaje, total
}

func main() {
	mensaje, result := sumar("Alex", 10, 20, 40, 70, 60)
	fmt.Println(mensaje, result)
}
