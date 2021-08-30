package models

type Persona struct {
	nombre string
	edad   int
}

func (p *Persona) Constructor(nombre string, edad int) {
	p.nombre = nombre
	p.edad = edad
}

func (p *Persona) GetNombre() string {
	return p.nombre
}

func (p *Persona) SetNombre(nombre string) {
	p.nombre = nombre
}

func (p *Persona) GetEdad() int {
	return p.edad
}

func (p *Persona) SetEdad(edad int) {
	p.edad = edad
}
