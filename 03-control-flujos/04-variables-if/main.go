package main

import (
	"fmt"
)

func main() {

	if nombre, edad := "Roel", 26; nombre == "Alex" {
		fmt.Println("Hola,", nombre)
	} else {
		fmt.Printf("Nombre: %s - Edad: %d\n", nombre, edad)
	}

	//Obtener valor de mapa
	users := make(map[string]string)

	users["Alex"] = "alex@gmail.con"
	users["Roel"] = "roel@gmail.con"

	//correo, ok := users["Juan"]

	if correo, ok := users["Roel"]; ok {
		fmt.Println(correo)
	} else {
		fmt.Println("No fue posible obtener el valor")
	}

}
