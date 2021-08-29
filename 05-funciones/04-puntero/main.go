package main

import "fmt"

//Punteros
func modificarValor(cadena *string) {
	fmt.Printf("%p\n", cadena)
	*cadena = "Hola desde la funcion"

}

func main() {
	cadena := "Hola Mundo de Go"
	fmt.Printf("%p\n", &cadena)
	fmt.Println("Antes de la función:", cadena)

	modificarValor(&cadena) //Referencia

	fmt.Println("Después de la función:", cadena)
}
