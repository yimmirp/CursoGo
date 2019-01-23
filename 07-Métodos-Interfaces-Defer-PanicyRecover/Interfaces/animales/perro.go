package animales

import "fmt"

//Perro es una estructura
type Perro struct {
	Nombre string
}

// Comunicarse exportado
func (p Perro) Comunicarse() {
	fmt.Println("woffff")
}

//Moverse es un metodo exportado
func (p Perro) Moverse() {
	fmt.Println("El perro se esta moviendo")
}
