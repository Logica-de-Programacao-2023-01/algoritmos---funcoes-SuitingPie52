package main

import (
	"fmt"
	"strconv"
	"strings"
)

func somaAlgarismos(x int) (int, error) {
	soma := 0
	xNum := strings.Split(strconv.Itoa(x), "")

	if x < 0 {
		return x, fmt.Errorf("número negativos")
	}

	for i := 0; i < len(xNum); i++ {
		algarismo, _ := strconv.Atoi(xNum[i])
		soma += algarismo
	}

	return soma, nil
}

func main() {
	x := 3428495

	fmt.Printf("A soma dos algorismos de %d é: ", x)
	fmt.Println(somaAlgarismos(x))
}
