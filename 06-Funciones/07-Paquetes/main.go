package main

import (
	"fmt"

	"github.com/ypernilloo/06-Funciones/07-Paquetes/despedirse"
	"github.com/ypernilloo/06-Funciones/07-Paquetes/saludar"
)

func main() {
	nombre := "gente del futuro"
	saludar.Saludar(nombre)
	saludar.MeVes = "Escrito desde el main"
	fmt.Println(saludar.MeVes)
	despedirse.Despedir(nombre)
}
