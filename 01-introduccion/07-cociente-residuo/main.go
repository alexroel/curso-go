package main

import "fmt"

//función que calcula cociente
func cociente(a, b int) int {
	return a / b
}

//Función que calcula residuo
func residuo(a, b int) int {
	return a % b
}

//Función principal
func main() {

	//Variables
	var a, b, c, r int

	//Entrada  de datos
	fmt.Print("Ingrese N01: ")
	fmt.Scanln(&a)

	fmt.Print("Ingrese N02: ")
	fmt.Scanln(&b)

	//Llamar la función cociente
	c = cociente(a, b)

	//LLamar la función residuo
	r = residuo(a, b)

	//Salida de datos
	fmt.Println("Cociente:", c)
	fmt.Println("Residuo:", r)
}
