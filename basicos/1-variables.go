package main

import "fmt"
import "strconv"

func main() {
	//Variables basicas
	var edad, edad2 int
	var cadena string 
	var bandera bool
	var cadenas [2]string

	edad = 15
	cadena = "Hola"
	cadena2 := "Mundo"
	cadenas[0] = "Primer Elemento"
	cadenas[1] = "Segundo Elemento"

	fmt.Println(cadena + " " + cadena2)

	//Int to String
	strEdad := strconv.Itoa(edad)
	fmt.Println("Edad 1 :" + strEdad)
	
	//String to Int 
	//intValue,error = strconv.Atoi(string)
	edad2,_ = strconv.Atoi(strEdad)
	fmt.Printf("Edad 2: %d\n", edad2)
	fmt.Printf("Edad 2: %v, %v %v\n", edad2, cadena, cadena2)
	fmt.Printf("Bandera: %t\n", bandera)
	fmt.Printf("Array: %v\n", cadenas)



	
}