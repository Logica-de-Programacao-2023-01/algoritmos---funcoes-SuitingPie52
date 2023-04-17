package main

import "fmt"

func concatenecaoString(x []string) string {
	var soma string

	for _, ranX := range x {
		soma += ranX
	}

	return soma
}

func main() {

	frases := []string{"hduiajsopdk", "ebaaaaaaa", "uepa"}

	fmt.Println("Sua concatenação é: ", concatenecaoString(frases))

}
