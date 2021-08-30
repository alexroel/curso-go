package main

import "fmt"

// Struct persona
type Persona struct {
	nombre string
	edad   int
}

//metodos
func (p *Persona) imprimir() {
	fmt.Printf("Nombre: %s\nEdad: %d\n", p.nombre, p.edad)
}

func (p *Persona) editNombre(nuevoNombre string) {
	p.nombre = nuevoNombre
}

//Herencia
type Empleado struct {
	Persona
	sueldo float64
}

func main() {
	p1 := Persona{"Alex", 26}

	p1.imprimir()
	//fmt.Println(p1)
	//p1.nombre = "Juan"
	p1.editNombre("Juan")

	p1.imprimir()
	//fmt.Println(p1)

	p2 := Persona{
		nombre: "Roel",
		edad:   27,
	}

	p2.imprimir()
	p2.editNombre("Carlos")
	p2.imprimir()

	//fmt.Println(p2)

	em1 := Empleado{
		sueldo: 500,
	}
	em1.nombre = "Pedro"
	em1.edad = 30
	em1.imprimir()
	fmt.Println(em1)

}
