package main

//import "paquetes/figuras"

//import "github.com/donvito/hellomod"
import "github.com/alexroel/figuras"

func main() {

	cua1 := figuras.Cuadrado{Lado: 10}
	figuras.Medidas(&cua1)

	//hellomod.SayHello()
	/*
		mesajes.Hola()
		mesajes.Imprimir()
	*/

	/*
		cua1 := figuras.Cuadrado{Lado: 8}
		figuras.Medidas(&cua1)

		cir1 := figuras.Circulo{Radio: 5}
		figuras.Medidas(&cir1)
	*/

	/**
	p1 := models.Persona{}
	p1.Constructor("Alex", 26)

	fmt.Println(p1)
	fmt.Println(p1.GetNombre())
	p1.SetNombre("Roel")

	fmt.Println(p1)
	*/

}
