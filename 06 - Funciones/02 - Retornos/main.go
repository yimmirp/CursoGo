package main

import "fmt"

func main() {
	/*
		var n1, n2 int8
		n1 = 19
		n2 = 5
		r := suma(n1, n2)
		fmt.Printf("El resultado es: %d\n", r)*/

	n := []int8{4, 7, 34, -4, 100, 3, 9, 56}
	max, min := maxmin2(n)
	fmt.Println("El maximo es:", max)
	fmt.Println("El minimo es", min)

}

func suma(a, b int8) int8 {

	return a + b
}

func maxmin(numeros []int8) (int8, int8) {
	var max, min int8

	for _, v := range numeros {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	return max, min

}

func maxmin2(numeros []int8) (max int8, min int8) {
	for _, v := range numeros {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	return

}
