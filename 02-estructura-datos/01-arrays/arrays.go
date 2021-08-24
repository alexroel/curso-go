package main

import "fmt"

func main() {
	//arrays
	var numeros [5]int
	fmt.Println(numeros)

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30

	fmt.Println(numeros)
	fmt.Println(numeros[4])

	//Arrays con valores
	nombres := [3]string{"Alex", "Roel", "Juan"}

	fmt.Println(nombres)

	colores := [...]string{
		"Rojo",
		"Verde",
		"Negro",
		"Azul",
	}

	fmt.Println(colores, len(colores))

	//Indices definidos
	momedas := [...]string{
		0: "Dolar",
		2: "Soles",
		3: "Pesos",
		5: "Euro",
	}

	fmt.Println(momedas, len(momedas))

	//sub Array
	subMoneda := momedas[3:]
	fmt.Println(subMoneda)

}
