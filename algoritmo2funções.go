package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func quantidadeVogais(frase string) int {
	soma := 0

	frase = strings.ToLower(frase)
	a := strings.Count(frase, "a")
	b := strings.Count(frase, "e")
	c := strings.Count(frase, "i")
	d := strings.Count(frase, "o")
	e := strings.Count(frase, "u")
	k := strings.Count(frase, "á")
	l := strings.Count(frase, "é")
	m := strings.Count(frase, "í")
	n := strings.Count(frase, "ó")
	o := strings.Count(frase, "ú")
	v := strings.Count(frase, "ã")
	w := strings.Count(frase, "õ")
	z := strings.Count(frase, "â")
	a1 := strings.Count(frase, "ê")
	b1 := strings.Count(frase, "î")
	c1 := strings.Count(frase, "ô")
	d1 := strings.Count(frase, "û")

	soma = a + b + c + d + e + k + l + m + n + o + v + w + z + a1 + b1 + c1 + d1
	return soma
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Escreva uma frase: ")
	scanner.Scan()
	x := scanner.Text()

	fmt.Println("Sua quantidade de vogais é", quantidadeVogais(x))
}
