package main

import "fmt"

func main() {
	//SLICES
	//1. Apuntador a un array
	//2. Tamaño
	//3. Capacidad
	//var nombres []string .. Pero hay otra forma de declarar y es la siguiente
	// make(tipo del slice, tamaño inicial, capacidad inicial)
	nombres := make([]string, 0)
	fmt.Printf("Su tamaño es %d y su capacidad es de %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Yimmi")
	fmt.Printf("Su tamaño es %d y su capacidad es de %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Daniel")
	fmt.Printf("Su tamaño es %d y su capacidad es de %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Ruano")
	fmt.Printf("Su tamaño es %d y su capacidad es de %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Pernillo")
	fmt.Printf("Su tamaño es %d y su capacidad es de %d\n", len(nombres), cap(nombres))

	fmt.Println(nombres)

	nombres2 := []string{
		"Karina",
		"Raquel",
		"Pernillo",
		"Dieguez",
	}

	fmt.Println(nombres2)
}
