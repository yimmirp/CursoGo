package main

import "fmt"

func main() {
	/*
		for i := 0; i < 5; i++ {
			fmt.Printf("Iteracion #%d\n", i)
		}*/

	/*	for i := 4; i >= 0; i-- {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}*/
	/*
		for i := 0; i < 5; i++ {
			if i == 2 {
				break
			}
			fmt.Printf("Iteracion #%d\n", i)
		}*/

	var matriz [3][3]int
	valor := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matriz[i][j] = valor
			valor = valor + 1
		}
	}

	fmt.Println(matriz)

}
