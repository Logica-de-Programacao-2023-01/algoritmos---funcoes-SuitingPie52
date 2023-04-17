ALGORITMO INCOMPLETO

package main

func ordemAlfabetica(x []string) (string, error) {

	var menor uint8
	menor = 99

	for _, ranX := range x {
		if ranX[0] < menor {
			menor = ranX[0]
		}
	}
}
