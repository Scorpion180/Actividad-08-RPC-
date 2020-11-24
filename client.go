package main

import (
	"fmt"
	"net/rpc"
)

type Request struct {
	Materia string
	Nombre  string
	Cal     float64
}

func client() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var op int64
	for {
		fmt.Println("1.- Agregar calificación de una materia")
		fmt.Println("2.- Mostrar el promedio de un alumno")
		fmt.Println("3.- Mostrar el promedio general")
		fmt.Println("4.- Mostrar el promedio de una materia")
		fmt.Println("0.- Salir")
		fmt.Scanln(&op)

		switch op {
		case 1:
			slice := []string{"", ""}
			var cal float64
			fmt.Println("Ingrese materia")
			fmt.Scanln(&slice[0])
			fmt.Println("Ingrese nombre")
			fmt.Scanln(&slice[1])
			fmt.Println("Ingrese calificación")
			fmt.Scanln(&cal)
			request := Request{slice[0], slice[1], cal}
			var result string
			err = c.Call("Server.AddCalMateria", request, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.AddCalMateria", request, "=", result)
			}
		case 2:
			var name string
			fmt.Print("Nombre del alumno: ")
			fmt.Scanln(&name)

			var result float64
			err = c.Call("Server.PromedioAlumno", name, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.PromedioAlumno", name, "=", result)
			}
		case 3:
			var name string = ""

			var result float64
			err = c.Call("Server.PromedioGeneral", name, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.PromedioGeneral", name, "=", result)
			}
		case 4:
			var name string
			fmt.Print("Nombre de la materia: ")
			fmt.Scanln(&name)
			var result float64
			err = c.Call("Server.PromedioMateria", name, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.PromedioMateria", result)
			}
		case 0:
			return
		}
	}
}

func main() {
	client()
}
