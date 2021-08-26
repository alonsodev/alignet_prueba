package main

import (
	"fmt"
	"unicode/utf8"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Es requerido un parametro con una cadena de caracteres")
		os.Exit(3)
	}
	cadenaStr := os.Args[1]

    cantidadCaracteres := utf8.RuneCountInString(cadenaStr)

    fmt.Println("Cantidad de caracteres:", cantidadCaracteres)
}
