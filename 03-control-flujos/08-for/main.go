package main

import "fmt"

func main() {
	//Bucle infinito
	/*for {
		fmt.Println("Hola Mundo")
	}*/

	//Bucle tipo while
	/**
	numeros := 12455458
	c := 0
	for numeros > 0 {
		numeros /= 10
		c++
	}

	fmt.Println("Cantidad de digitos es: ", c)
	*/

	//Bucle for

	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

}
