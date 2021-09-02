package main

import (
	"fmt"
	"net/http"
	"time"
)

func revisarServidor(servidor string, canal chan string) {
	_, err := http.Get(servidor)

	if err != nil {
		//fmt.Println(servidor, "No esta disponible")
		canal <- servidor + " No esta disponible"
	} else {
		//fmt.Println(servidor, "Esta funcionando")
		canal <- servidor + " Esta funcionando"
	}
}

func main() {
	inicio := time.Now()

	canal := make(chan string)

	servidores := []string{
		"https://oregoom.com/",
		"https://www.udemy.com/",
		"https://www.youtube.com/",
		"https://www.facebook.com/",
		"https://www.google.com/",
	}

	for _, sevidor := range servidores {
		go revisarServidor(sevidor, canal)
	}

	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-canal)
	}

	tiempoPaso := time.Since(inicio)
	fmt.Println("Tiempo de ejecución: ", tiempoPaso)
}
