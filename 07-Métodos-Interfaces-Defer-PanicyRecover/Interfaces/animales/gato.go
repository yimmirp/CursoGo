package animales

import "fmt"

//Gato es una estructura
type Gato struct {
	Nombre string
}

//Comunicarse es un metodo exportado
func (g Gato) Comunicarse() {
	fmt.Println("Miauuuu")
}

//Moverse es un metodo exportado
func (g Gato) Moverse() {
	fmt.Println("El gato se esta moviendo")
}
