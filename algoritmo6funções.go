package main

import (
	"fmt"
)

func concatenacaoComVirgula(x []string) (string, error) {

	soma := ""
	i := 0

	if len(x) == 0 {
		return "", fmt.Errorf("slice vazio")
	}
	for i = 0; len(x) > i; i++ {

		if i == len(x)-1 {
			soma += x[i]
		} else {
			soma += x[i] + ","
		}
	}

	return soma, nil

}

func main() {
	frases := []string{"fafsafa", "ebaaaaaaaa", "kkkkkk", "putz"}
	conjVazio := []string{}

	fmt.Print("Sua string concatenada com vírgulas é: ")
	fmt.Println(concatenacaoComVirgula(frases))
	fmt.Print("Sua string concatenada com vírgulas é: ")
	fmt.Println(concatenacaoComVirgula(conjVazio))

}
