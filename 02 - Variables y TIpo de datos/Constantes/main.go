package main

import "fmt"

func main(){
	//const constante = "Yimmi"
	//fmt.Println(constante)

	var a int
	var b int8

	a= 121211
	b=5

	//casting
	c:= a+ int(b)
	fmt.Println(c)
	n:="Pedro"
	p:="Gomez"

	fmt.Printf("Hola %s %s como estas?",n,p)
}

//El comentario de linea se utiliza para documentar el codigo

/*
	Y el comentario de bloque para dejar codigo que ya no se usara
	a:=5
*/