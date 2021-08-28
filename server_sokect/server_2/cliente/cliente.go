package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func main() {
	go cliente()
	var input string
	fmt.Scanln(&input)
}

func cliente() {
	c, err := net.Dial("tcp", ":9999")

	if err != nil {
		fmt.Println(err)
		return
	}
	msg := "hola mundo"
	fmt.Println(msg)
	err = gob.NewEncoder(c).Encode(msg)

	if err != nil {
		fmt.Println(err)
	}

	c.Close()
}
