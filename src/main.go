package main

import "fmt"

func main() {
	println("Boletim FIAP")

	for _, class := range grades {
		fmt.Println(class.ToString())
	}
}
