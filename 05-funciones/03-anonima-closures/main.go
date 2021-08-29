package main

import (
	"fmt"
	"strings"
)

//Closures y funciones anonima
func repeat(n int) func(cadena string) string {

	return func(cadena string) string {
		return strings.Repeat(cadena, n)
	}
}

func main() {

	repeat3 := repeat(3)
	fmt.Println(repeat3("Hola"))
	fmt.Println(repeat3("Mundo"))

	repeat5 := repeat(5)
	fmt.Println(repeat5("Alex"))
	fmt.Println(repeat5("Roel"))

	//Función anonima
	/*
			func() {
				fmt.Println("Hola desde la función anónima")
			}()

		myfunc := func(nombre string) string {
			return fmt.Sprintf("Hola, %s desde la función anónima", nombre)
		}

		fmt.Println(myfunc("Alex"))
	*/
}
