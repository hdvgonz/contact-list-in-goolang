package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

//Guardar contactos en archivo JSON

func saveContactToFile(contact []Contact) error {

	file, err := os.Create("contacts.json")

	if err != nil{
		return err;
	}

	defer file.Close();
	//encoder convierte estructuras de datos en formato JSON
	encoder := json.NewEncoder(file); //Escribe en el archivo
	err = encoder.Encode (contact);//intenta codificar el slice de Contact en JSON y escribirlo en el archivo.

	if err != nil{
		return err;
	}

	return nil;
}

//Cargar contactos desde archivo JSON

func loadContactFromFile( contacts *[]Contact) error{

	file, err := os.Open("contacts.json");
	
	if err != nil{
		return err;
	}

	defer file.Close();

	decoder := json.NewDecoder(file);
	err = decoder.Decode(&contacts);

	if err != nil{
		return err;
	}

	return nil;

}


func main() {

	//slice contacts

	var contacts []Contact;

	//Cargar contactos existentes desde archivo JSON

	err := loadContactFromFile(&contacts);
	if err != nil {
		fmt.Print("Error al cargar los contactos", err);
	}

	//Insertar un nuevo contacto / crear instancia de bufio

	reader := bufio.NewReader(os.Stdin);

	for{
		fmt.Print("==== Gestor de contactos ====\n",
		"1. Agregar contacto\n",
		"2. Mostrar contactos\n",
		"3. Salir\n",
		"Elige una opcion: ");

		var option int;
		_, err = fmt.Scanln(&option);
		if err != nil{
			fmt.Print("Error al leer la opcion: ", err);
			continue;
		}
		//Manejar opciones del menu

		switch option{
		case 1:
			var c Contact;
			fmt.Print("Nombre: ");
			c.Name,_ = reader.ReadString('\n');
			fmt.Print("Email: ");
			c.Email, _ = reader.ReadString('\n');
			fmt.Print("Telefono: ");
			c.Phone, _ = reader.ReadString('\n');

			contacts = append(contacts, c);

			//Guardar en un archivo JSON
			err := saveContactToFile(contacts);
			if  err != nil {
				fmt.Print("Error al guardar los contactos: ", err);
			}

		case 2:
			//Mostrar toddos los contactos
			for index, contact := range contacts{
				fmt.Printf("%d, Nombre: %s Email: %s Phone: %s\n", index+1, contact.Name, contact.Email, contact.Phone)
			}
		
		case 3:
			return
		
		default:
			fmt.Println("Saliendo del programa...")
			return;
		}



		
	}

}