package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

type Persona struct {
	Nombre string
	Email  []string
}

func main() {

	persona := Persona{
		Nombre: "alejandro",
		Email: []string{
			"correos academicos ",
			"correos personales",
		},
	}

	for j := 0; j <= 11; j++ {
		fmt.Println(j)
		go cliente(persona)
	}

	var input string
	fmt.Scanln(&input)
}

func cliente(persona Persona) {
	c, err := net.Dial("tcp", ":9999")

	if err != nil {
		fmt.Println(err)
		return
	}

	err = gob.NewEncoder(c).Encode(persona)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}
