package main

import "fmt"

func main() {
	saludarVarios(18, "pedro", "Daniel", "Juan")
}

func saludarVarios(edad uint8, nombres ...string) {
	for _, v := range nombres {
		fmt.Printf("El nombre es: %s y la edad %d\n", v, edad)
	}
}
