package main

import "fmt"

//Persona es una estructura
type Persona struct {
	Nombre string
	Edad   uint
}

func main() {

	/*var persona1 Persona
	persona1.Nombre = "Yimmi"
	persona1.Edad = 3
	fmt.Println(persona1)*/

	persona2 := Persona{
		Nombre: "Yimm",
		Edad:   23,
	}
	fmt.Println(persona2)

	persona3 := Persona{
		"Jose",
		25,
	}

	fmt.Println(persona3)

}
