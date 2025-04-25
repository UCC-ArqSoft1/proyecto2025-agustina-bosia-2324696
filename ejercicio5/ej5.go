package main

import (
	"fmt" // Importamos el paquete fmt para imprimir en consola
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
	ListaTareas []Tarea
)

type Tarea struct {
	nombreT     string
	descripcion string
	estadoT     bool
}

func (t *Tarea) Mostrar() {
	fmt.Println("- ", t.nombreT)
	fmt.Println("Descripción:", t.descripcion)
	fmt.Println("Completado:", t.estadoT)
	fmt.Println("========================================")
}

//func caso1 () {

func main() {
	var listaTareas = make([]Tarea, 0)
	var opcion int
	var nombreTarea, descripcion string
	var estadoT bool

	fmt.Println("Seleccione una opción: \n1. Agregar tarea\n2. Marcar Tarea como Completada\n3. Editar Tarea\n4. Eliminar Tarea\n5. Salir\nIngrese su opción: ")

	fmt.Scanln(&opcion)

	switch opcion {
	case 1:
		fmt.Println("Ingrese el nombre de la tarea:")
		fmt.Scanln(&nombreTarea)
		fmt.Println("Ingrese la descripción de la tarea:")
		fmt.Scanln(&descripcion)
		estadoT = false
		tarea := Tarea{nombreT: nombreTarea, descripcion: descripcion, estadoT: estadoT}
		listaTareas = append(listaTareas, tarea)
		fmt.Println("Tarea agregada")
		fmt.Println("========================================")
		for i, tarea := range listaTareas {
			fmt.Printf("Tarea %d.\n", i+1)
			tarea.Mostrar()
		}
	case 2:
		fmt.Println("Ingrese el nombre de la tarea a marcar como completada:")
		fmt.Scanln(&nombreTarea)
		for i, tarea := range listaTareas {
			if tarea.nombreT == nombreTarea {
				listaTareas[i].estadoT = true
				fmt.Println("Tarea marcada como completada")
				fmt.Println("========================================")
				for i, tarea := range listaTareas {
					fmt.Printf("Tarea %d.\n", i+1)
					tarea.Mostrar()
				}
				break
			}
		}

	case 3:
	case 4:
	case 5:
		fmt.Println("Saliendo del programa...")
		return
	default:
		fmt.Println("Opción no válida. Intente nuevamente.")
	}

}
