package main

import (
	"fmt" // Importamos el paquete fmt para imprimir en consola
	"math"
)

const (
	// Constantes para el programa
	pi      = 3.14
	Gravity = 9.81
	mensaje = "Hola Mundo"
)

var (
	// Variables para el programa
	Nombre      = "Agustina Bosia"
	Edad        = 21
	Lado, Lado2 float64
)

func main() {
	fmt.Println(mensaje)
	Edad = 22 // Cambiamos el valor de la variable Edad
	fmt.Println("Nombre:", Nombre, "Edad:", Edad)

	fmt.Println("Ingrese lado 1 de un triangulo rectangulo")
	fmt.Scanln(&Lado)
	fmt.Println("Ingrese lado 2 de un triangulo rectangulo")
	fmt.Scanln(&Lado2)

	fmt.Println("lado 1 de un triangulo rectangulo: ", Lado)
	fmt.Println("lado 2 de un triangulo rectangulo: ", Lado2)

	var hipotenusa = math.Sqrt(Lado*Lado + Lado2*Lado2)
	fmt.Println("La hipotenusa es:", hipotenusa)

	var area = (Lado * Lado2) / 2
	fmt.Println("El area es:", area)

	var perimetro = Lado + Lado2 + hipotenusa
	fmt.Println("EL perimetro es: ", perimetro)

}
