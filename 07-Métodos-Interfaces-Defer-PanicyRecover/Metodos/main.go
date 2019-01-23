package main

import "fmt"

type Persona struct {
	nombre string
	edad   int8
}

type Numero int

func (p Persona) saludar() {
	fmt.Println("Hola soy una persona")
}

func (n Numero) presentarse() {
	fmt.Println("Hola, yo soy un número")
}

func (p *Persona) asignarEdad(i int8) {
	if i >= 0 {
		p.edad = i
	} else {
		fmt.Println("No es válida la edad negativa")
	}
}
func (p *Persona) asignarNombre(nombre string) {
	p.nombre = nombre
}
func main() {
	var persona Persona
	// var numero Numero
	persona.saludar()
	// numero.presentarse()
	persona.asignarEdad(35)
	persona.asignarNombre("Yimmi")
	fmt.Println(persona)
}
