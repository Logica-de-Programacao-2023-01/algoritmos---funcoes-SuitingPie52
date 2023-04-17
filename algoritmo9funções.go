package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func palavrasContidasNoString(x string) ([]string, error) {

	conjPalavras := strings.Split(x, " ")

	if len(x) == 0 {
		return conjPalavras, fmt.Errorf("slice vazio")
	}

	return conjPalavras, nil

}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Informe um string: ")
	scanner.Scan()
	x := scanner.Text()

	fmt.Print("Suas palavras são: ")
	fmt.Println(palavrasContidasNoString(x))

}
