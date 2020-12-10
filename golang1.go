package main

import ("fmt")

func main(){

	var nombre string = "Diego Alejandro Tobon"
	var edad int = 20

	fmt.Println("Hola Mundo")
	fmt.Println("Bienvenido ",nombre)
	if (edad >= 18){
		fmt.Println("Mayor de edad")
	} else{
		fmt.Println("No puedes ... :( ")
	}
	var apellido,ciudad,celular string = "Tobon Henao","Manizales","3136702173"
	fmt.Println(nombre,apellido,edad,celular, ciudad)


}

