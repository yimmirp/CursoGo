package main

import "fmt"

func main() {
	//Arrays
	//Todos los valores deben ser del mismo tipo de dato
	//Tamaño fijo
	var nombres [3]string
	var numeros [3]int
	fmt.Println(numeros)
	fmt.Println(nombres)
	nombres[0] = "Yimmi"
	nombres[1] = "Daniel"
	nombres[2] = "Ruano"

	nombres2 := [3]string{
		"Daniel",
		"Ruano",
		"Pernillo",
	}
	size := len(nombres2)
	variable := nombres2[1]
	fmt.Println(nombres)
	fmt.Println(nombres2)
	fmt.Println("El tamaño es:", size)
	fmt.Println("El nombre es", variable)
}
