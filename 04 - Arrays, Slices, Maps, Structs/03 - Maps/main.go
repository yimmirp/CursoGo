package main

import "fmt"

func main() {

	/*idiomas := make(map[string]string)
	idiomas["es"] = "Español"
	idiomas["en"] = "Ingles"
	idiomas["fr"] = "Frances"
	fmt.Println(idiomas)*/

	idiomas := map[string]string{
		"es": "Español",
		"en": "Ingles",
		"fr": "Frances"}
	//fmt.Println(idiomas["es"])
	fmt.Println(idiomas)

	//Borrar datos del mapa
	delete(idiomas, "en")
	fmt.Println(idiomas)

}
