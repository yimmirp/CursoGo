package main

import (
	"errors"
	"fmt"
)

func main() {
	r, err := division(100, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r, err)
}

func division(dividendo, divisor float64) (resultado float64, err error) {
	if divisor == 0 {
		err = errors.New("El divisor es O")
	} else {
		resultado = dividendo / divisor
	}
	return
}
