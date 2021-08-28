package main

import (
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
	mgs := "hola mundo"
	fmt.Println(mgs)
	c.Write([]byte(mgs))
	c.Close()
}
