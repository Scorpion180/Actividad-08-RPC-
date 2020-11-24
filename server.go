package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Request struct {
	Materia string
	Nombre  string
	Cal     float64
}

var materias = make(map[string]map[string]float64)

type Server struct{}

func (this *Server) PromedioAlumno(name string, reply *float64) error {
	var promedio float64 = 0
	var cont float64 = 0
	for key := range materias {
		if val, ok := materias[key][name]; ok {
			promedio += val
			cont += 1
		}
	}
	*reply = promedio / cont
	return nil
}

func (this *Server) PromedioGeneral(name string, reply *float64) error {
	var promedio float64 = 0
	var cont float64 = 0
	for _, item := range materias {
		for _, value := range item {
			promedio += value
			cont += 1
		}
	}
	*reply = promedio / cont
	return nil
}

func (this *Server) PromedioMateria(name string, reply *float64) error {
	var promedio float64 = 0
	var cont float64 = 0
	for _, value := range materias[name] {
		promedio += value
		cont += 1
	}
	*reply = promedio / cont
	return nil
}

func (this *Server) AddCalMateria(s Request, reply *string) error {
	if _, ok := materias[s.Materia]; ok {
		materias[s.Materia][s.Nombre] = s.Cal
	} else {
		materias[s.Materia] = make(map[string]float64)
		materias[s.Materia][s.Nombre] = s.Cal
	}
	*reply = "Exito"
	return nil
}

func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main() {

	go server()

	var input string
	fmt.Scanln(&input)
}
