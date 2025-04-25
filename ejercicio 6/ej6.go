package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Contact struct {
	Name  string `json:"nombre"`
	Phone string `json:"telefono"`
	Email string `json:"email"`
}

func saveContactFile(contacto Contact) error {
	file, err := os.Create("lista-contactos.json")

	if err != nil {
		return errors.New("Error al crear el archivo")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacto)
	if err != nil {
		return errors.New("Error al formatear el contacto")
	}

	return nil
}

func loadContactList(contactList *[]Contact) error {
	file, err := os.Open("lista-contactos.json")
	if err != nil {
		return errors.New("Error al abrir el archivo")
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(contactList)
	if err != nil {
		return errors.New("Error al decodificar el contacto")
	}

	return nil
}

func main() {
	var opcion int
	var contactList []Contact

	err := loadContactList(&contactList)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	for {
		fmt.Println("GESTOR DE CONTACTOS")
		fmt.Println("1. Agregar contacto")
		fmt.Println("2. Listar contactos")
		fmt.Println("3. Buscar contacto")
		fmt.Println("0. Salir")

		fmt.Scanln(&opcion)

		switch opcion {
		case 1:

			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Ingrese el nombre del contacto:")
			nombre, _ := reader.ReadString('\n')
			fmt.Println("Ingrese el teléfono del contacto:")
			telefono, _ := reader.ReadString('\n')
			fmt.Println("Ingrese el email del contacto:")
			email, _ := reader.ReadString('\n')

			nombre = strings.TrimSpace(nombre)
			telefono = strings.TrimSpace(telefono)
			email = strings.TrimSpace(email)

			con := Contact{Name: nombre, Phone: telefono, Email: email}
			contactList = append(contactList, con)

			err := saveContactFile(con)

			if err != nil {
				fmt.Println("Error: ", err)
				break
			}

		case 2:
			fmt.Println("Lista de contactos:")
			for index, contacto := range contactList {
				fmt.Printf("%d. Nombre: %s, Teléfono: %s, Email: %s \n",
					index+1, contacto.Name, contacto.Phone, contacto.Email)
			}
		case 3:
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Ingrese el nombre del contacto a buscar:")
			nombre, _ := reader.ReadString('\n')
			nombre = strings.TrimSpace(nombre)
			contactoEncontrado := false
			for _, contacto := range contactList {
				if strings.EqualFold(contacto.Name, nombre) {
					fmt.Printf("Nombre: %s, Teléfono: %s, Email: %s \n",
						contacto.Name, contacto.Phone, contacto.Email)
					contactoEncontrado = true
					break
				}
			}
			if !contactoEncontrado {
				fmt.Println("Contacto no encontrado")
			}
		case 0:
			fmt.Println("Saliendo del programa...")
			return
		default:
			println("Opción no válida")
		}
	}

}
