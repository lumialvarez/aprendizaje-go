package main

import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	var nombre string
	var edad int
	fmt.Printf("Ingresa tu nombre: ")
	fmt.Scanf("%s\n", &nombre)

	fmt.Printf("Ingresa tu edad: ")
	fmt.Scanf("%d\n", &edad)

	fmt.Printf("Nombre: %s\n", nombre)
	fmt.Printf("Edad: %d\n", edad)


	//Paquete bufio
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingresa Apellido:")
	apellido,err := reader.ReadString('\n')
	if(err != nil){
		fmt.Println(err)
	} else {
		fmt.Println("apellido: " + apellido)
	}

}