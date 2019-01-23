package main

import (
	"github.com/ypernilloo/07-Métodos-Interfaces-Defer-PanicyRecover/Interfaces/animales"
)

func main() {
	var p animales.Perro
	var g animales.Gato
	p.Nombre = "Firulais"
	g.Nombre = "Manchas"
	// AdoptarPerro(p)
	// AdoptarGato(g)
	AdoptarMascota(p)
	AdoptarMascota(g)
}

/*
func AdoptarPerro(p animales.Perro) {
	p.Comunicarse()
}
func AdoptarGato(g animales.Gato) {
	g.Comunicarse()
}
*/

func AdoptarMascota(m animales.Mascota) {
	m.Comunicarse()
	m.Moverse()
}
