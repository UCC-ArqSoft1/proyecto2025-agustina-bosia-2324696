package main

import (
	"fmt" // Importamos el paquete fmt para imprimir en consola
	"math"
	"math/rand"
)

var (
	// Variables para el programa
	Num    int
	Number int
)

func main() {

	var randomNumber = rand.Intn(100)

	Number = math.Round(float64(randomNumber))

	fmt.Println("Número aleatorio:", randomNumber)

	for i := 0; i < 10; i++ {
		fmt.Println("Ingrese un numero")
		fmt.Scanln(&Num)

		if Num == randomNumber {
			fmt.Println("Acertaste. Felicitaciones")
			i = 10
		} else if Num > randomNumber {
			fmt.Println("No le pegate. El número es Menor. Te quedan ", 9-i, "intentos")
		} else {
			fmt.Println("No le pegate. El número es Mayor. Te quedan ", 9-i, "intentos")
		}
	}

}
