package main

import "fmt"

func primo(x int) (bool, error) {

	if x < 0 {
		return false, fmt.Errorf("seu número é negativo")
	}

	primoBool := true
	i := 0

	if x == 1 || x == 0 {

		primoBool = false
	}

	for i = 2; i < x; i++ {

		if x%i == 0 {

			primoBool = false

		}

	}

	return primoBool, nil

}

func main() {

	var x int

	fmt.Println("Informe um número:")
	fmt.Scanln(&x)

	fmt.Print("Seu número é primo está: ")
	fmt.Print(primo(x))

}
