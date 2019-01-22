package main

import "fmt"

func saludar(nombre string, edad uint8){
	fmt.Printf("Mi nombre es %s y tengo %d años",nombre ,edad)
	if edad > 30 {
		return
	}
	fmt.Println("Eres Joven")
}

func main(){
	saludar("Yimmi",5)
}