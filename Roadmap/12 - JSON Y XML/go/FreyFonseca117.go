/*
 * IMPORTANTE: Sólo debes subir el fichero de código como parte del ejercicio.
 *
 * EJERCICIO:
 * Desarrolla un programa capaz de crear un archivo XML y JSON que guarde los
 * siguientes datos (haciendo uso de la sintaxis correcta en cada caso):
 * - Nombre
 * - Edad
 * - Fecha de nacimiento
 * - Listado de lenguajes de programación
 * Muestra el contenido de los archivos.
 * Borra los archivos.
 *
 * DIFICULTAD EXTRA (opcional):
 * Utilizando la lógica de creación de los archivos anteriores, crea un
 * programa capaz de leer y transformar en una misma clase custom de tu
 * lenguaje los datos almacenados en el XML y el JSON.
 * Borra los archivos.
 */
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	BirthDate string   `json:"birthdate"`
	Lenguajes []string `json:"lenguajes"`
}

func checkError(err error) {
	if err != nil {
		return
	}
}

func main() {
	user := user{
		Name:      "Ana Programadora",
		Age:       30,
		BirthDate: "1993-05-15",
		Lenguajes: []string{"Go", "Python", "JavaScript"},
	}

	// formatear json para que sea mas legible
	jsonIdent, err := json.MarshalIndent(user, "", "  ")
	checkError(err)

	// Crear el archivo y escribir el archivo

	err = os.WriteFile("datos.json", jsonIdent, 0644)
	checkError(err)

	//Leer el archivo

	leerJson, err := os.ReadFile("datos.json")
	checkError(err)
	fmt.Println("El archivo contiene los siguientes datos")
	fmt.Println(string(leerJson))

	//borrar el archivo creado
	os.Remove("datos.json")
	fmt.Println("Luego de leido se ha borrado el archivo")

}
