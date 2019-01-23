package saludar

import "fmt"

//MeVes es exportable
var MeVes string
var noMeVes string

//Saludar es una funcion exportada
func Saludar(nombre string) {
	fmt.Println("Saludos ", nombre)

}
