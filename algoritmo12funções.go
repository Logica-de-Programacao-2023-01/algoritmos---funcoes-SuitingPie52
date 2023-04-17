package main

import (
	"fmt"
	"strings"
)

func comecoMaiusculo(x []string) (string, error) {

	if len(x) == 0 {
		return "", fmt.Errorf("stringvazio")
	}

	y := ""

	for i := 0; i < len(x); i++ {
		if string(x[i][0]) == strings.ToUpper(string(x[i][0])) {
			y += x[i]
		}
	}

	return y, nil
}

func main() {

	palavras := []string{"eba", "Tubarão", "Deus", "deus", "Brasil", "kkkkkk"}

	fmt.Print("As palavras que começam com letra maiúscula são: ")
	fmt.Println(comecoMaiusculo(palavras))

}
