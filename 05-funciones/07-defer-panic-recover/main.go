package main

import (
	"fmt"
	"os"
)

func main() {

	//Uso de funcion recover
	defer func() {
		if error := recover(); error != nil {
			fmt.Println("Ups!, al parecer el programa no finalizo de forma correcta")
		}
	}()

	if file, error := os.Open("hoa.txt"); error != nil {
		//Uso de función panic
		panic("No fue posible óbtener el archivo!!")
	} else {

		//Uso de defer
		defer func() {
			fmt.Println("Cerrando el archivo!")
			file.Close()
		}()

		contenido := make([]byte, 254)
		long, _ := file.Read(contenido)
		contenidoArchivo := string(contenido[:long])
		fmt.Println(contenidoArchivo)
	}
}
