/**
  Exemplo de como "incluir um método a um float64"
*/

package main

import "fmt"

type nota float64

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceito(n nota) string {
	if n.entre(9.0, 10.0) {
		return "A"
	} else if n.entre(7.0, 8.99) {
		return "B"
	}

	return "E"
}

func main() {
	fmt.Println(notaParaConceito(10))
	fmt.Println(notaParaConceito(8))
	fmt.Println(notaParaConceito(5))
	fmt.Println(notaParaConceito(0))
}
