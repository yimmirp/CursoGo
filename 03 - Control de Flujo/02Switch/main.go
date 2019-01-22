package main

import "fmt"

func main(){
/*	a:= 5

	switch a {
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		fallthrough
	case 4:
		fallthrough
	case 5:
		fmt.Println("Estas entre semana")
	case 6: 
		fallthrough
	case 7:
		fmt.Println("Estas en fin de semana")
	default:
		fmt.Println("No es un dia Valido")
	}
*/


switch a:= 5;{
 case a >= 0 && a <= 5 :
	fmt.Println("Estas entre semana")
 case a >= 6 && a <= 7:
	fmt.Println("Estas en Fin de semana")
 default:
	fmt.Println("No es un dia Valido")
 }
}