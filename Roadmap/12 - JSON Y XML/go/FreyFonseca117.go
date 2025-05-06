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

import "encoding/json"

type User struct {
	name      string   `json:"name"`
	age       int      `json:"age"`
	birthDate string   `json:"birtdate"`
	lenguajes []string `json:"lenguajes"`
}

func err(err error) {
	if err != nil {
		return
	}
}

func main() {
	user := User{
		name:      "Ana Programadora",
		age:       30,
		birthDate: "1993-05-15",
		lenguajes: []string{"Go", "Python", "JavaScript"},
	}

	// formatear json para que sea mas legible
	jsonIdent, err := json.MarshalIndent(user, "", "  ")

}

// package main

// import (
// 	"encoding/json"
// 	"encoding/xml"
// 	"fmt"
// 	"os"
// 	"time"
// )

// // Estructura para almacenar los datos que vamos a usar en XML y JSON
// type Persona struct {
// 	Nombre      string   `xml:"nombre" json:"nombre"`
// 	Edad        int      `xml:"edad" json:"edad"`
// 	FechaNac    string   `xml:"fecha_nacimiento" json:"fecha_nacimiento"`
// 	Lenguajes   []string `xml:"lenguajes>lenguaje" json:"lenguajes"`
// }

// // Función para verificar errores (patrón común en Go)
// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func main() {
// 	// =============================================
// 	// PARTE PRINCIPAL DEL EJERCICIO
// 	// =============================================

// 	// Crear una instancia de Persona con datos de ejemplo
// 	persona := Persona{
// 		Nombre:    "Ana Programadora",
// 		Edad:      30,
// 		FechaNac:  "1993-05-15",
// 		Lenguajes: []string{"Go", "Python", "JavaScript"},
// 	}

// 	// =====================
// 	// Manejo del archivo JSON
// 	// =====================

// 	// MarshalIndent para formatear el JSON con sangría
// 	jsonData, err := json.MarshalIndent(persona, "", "  ")
// 	checkError(err)

// 	// Escribir el JSON a un archivo
// 	err = os.WriteFile("datos.json", jsonData, 0644)
// 	checkError(err)

// 	// Leer y mostrar el contenido del JSON
// 	jsonLeido, err := os.ReadFile("datos.json")
// 	checkError(err)
// 	fmt.Println("Contenido del archivo JSON:")
// 	fmt.Println(string(jsonLeido))

// 	// =====================
// 	// Manejo del archivo XML
// 	// =====================

// 	// MarshalIndent para formatear el XML con sangría
// 	xmlData, err := xml.MarshalIndent(persona, "", "  ")
// 	checkError(err)

// 	// Añadir cabecera XML y escribir a archivo
// 	xmlData = []byte(xml.Header + string(xmlData))
// 	err = os.WriteFile("datos.xml", xmlData, 0644)
// 	checkError(err)

// 	// Leer y mostrar el contenido del XML
// 	xmlLeido, err := os.ReadFile("datos.xml")
// 	checkError(err)
// 	fmt.Println("\nContenido del archivo XML:")
// 	fmt.Println(string(xmlLeido))

// 	// =====================
// 	// Limpieza: Borrar archivos
// 	// =====================

// 	err = os.Remove("datos.json")
// 	checkError(err)

// 	err = os.Remove("datos.xml")
// 	checkError(err)

// 	fmt.Println("\nArchivos JSON y XML borrados correctamente.")

// 	// =============================================
// 	// DIFICULTAD EXTRA (opcional)
// 	// =============================================

// 	fmt.Println("\n=== Parte Extra ===")

// 	// Volver a crear los archivos para la parte extra
// 	err = os.WriteFile("datos.json", jsonData, 0644)
// 	checkError(err)

// 	err = os.WriteFile("datos.xml", xmlData, 0644)
// 	checkError(err)

// 	// Leer y transformar ambos archivos a la estructura Persona
// 	var personaDesdeJSON Persona
// 	var personaDesdeXML Persona

// 	// Leer y decodificar JSON
// 	jsonDataLeido, err := os.ReadFile("datos.json")
// 	checkError(err)
// 	err = json.Unmarshal(jsonDataLeido, &personaDesdeJSON)
// 	checkError(err)

// 	// Leer y decodificar XML
// 	xmlDataLeido, err := os.ReadFile("datos.xml")
// 	checkError(err)
// 	err = xml.Unmarshal(xmlDataLeido, &personaDesdeXML)
// 	checkError(err)

// 	// Mostrar los datos decodificados
// 	fmt.Println("\nDatos desde JSON:")
// 	mostrarPersona(personaDesdeJSON)

// 	fmt.Println("\nDatos desde XML:")
// 	mostrarPersona(personaDesdeXML)

// 	// Borrar archivos nuevamente
// 	err = os.Remove("datos.json")
// 	checkError(err)

// 	err = os.Remove("datos.xml")
// 	checkError(err)
// }

// // Función auxiliar para mostrar los datos de una Persona
// func mostrarPersona(p Persona) {
// 	fmt.Printf("Nombre: %s\n", p.Nombre)
// 	fmt.Printf("Edad: %d\n", p.Edad)
// 	fmt.Printf("Fecha de Nacimiento: %s\n", p.FechaNac)
// 	fmt.Printf("Lenguajes: %v\n", p.Lenguajes)
// }
