/*
* File              : main.go
* Author			: Eliazar
* Creation date     : 10/03/2024
* Last modified by  : Eliazar
* Last modified date: 10/03/2024
* Description       : This file contains go with mySQL functions
 */

package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"go-mysql/handlers"
	"go-mysql/models"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"go-mysql/database"
)

func main() {
	// Establecer conexión a la base de datos
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("Error al cerrar la conexión a la base de datos:", err)
		}
	}(db)

	for {
		fmt.Println("\nMenú:")
		fmt.Println("1. Listar contactos")
		fmt.Println("2. Obtener contacto por ID")
		fmt.Println("3. Crear nuevo contacto")
		fmt.Println("4. Actualizar contacto")
		fmt.Println("5. Eliminar contacto")
		fmt.Println("6. Salir")
		fmt.Print("Seleccione una opción: ")

		// Leer la opción seleccionada por el usuario
		var option int
		_, err := fmt.Scanln(&option)
		if err != nil {
			log.Fatal("Error al leer la opción seleccionada:", err)
		}

		// Ejecutar la opción seleccionada
		switch option {
		case 1:
			handlers.ListContacts(db)
		case 2:
			fmt.Print("Ingrese el ID del contacto: ")
			var idContact int
			_, err := fmt.Scanln(&option)
			if err != nil {
				log.Fatal("Error al leer el ID del contacto:", err)
			}
			handlers.GetContactByID(db, idContact)
		case 3:
			newContact := inputContactDetails()
			handlers.CreateContact(db, newContact)
		case 4:
			updatedContact := inputContactDetails()
			handlers.UpdateContact(db, updatedContact)
		case 5:
			fmt.Print("Ingrese el ID del contacto que quiere eliminar: ")
			var idContact int
			_, err := fmt.Scanln(&option)
			if err != nil {
				log.Fatal("Error al leer el ID del contacto a eliminar:", err)
			}
			handlers.DeleteContact(db, idContact)
		case 6:
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("Opción no válida. Por favor, seleccione una opción válida.")
		}
	}
}

// Función para ingresar los detalles del contacto desde la entrada estándar
func inputContactDetails() models.Contact {
	// Leer la entrada del usuario utilizando bufio
	reader := bufio.NewReader(os.Stdin)

	var contact models.Contact

	fmt.Print("Ingrese el nombre del contacto: ")
	name, _ := reader.ReadString('\n')
	contact.Name = strings.TrimSpace(name)

	fmt.Print("Ingrese el correo electrónico del contacto: ")
	email, _ := reader.ReadString('\n')
	contact.Email = strings.TrimSpace(email)

	fmt.Print("Ingrese el número de teléfono del contacto: ")
	phone, _ := reader.ReadString('\n')
	contact.Phone = strings.TrimSpace(phone)

	return contact
}
