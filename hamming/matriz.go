package hamming

import (
	"fmt"
)

type matriz struct {
	dato [][]bool
}

func Matriz() {
	k := [][]bool{{false, true, false}, {true, true, false}}
	ma := matriz{dato: k}
	fmt.Println("Matriz")
	fmt.Println(ma)
	fmt.Println(ma.ToString())
}
func (this *matriz) Multiplicar(n matriz) matriz {
	return n
}

//Funcion que convierte una matriz en un string para imprimir en consola
func (this *matriz) ToString() string {
	var resultado string
	for i, r := range this.dato {
		resultado = resultado + "|"
		for j, d := range r {
			fmt.Println(j)
			fmt.Println(d)
			if d {
				resultado = resultado + "1"
			} else {
				resultado = resultado + "0"
			}
		}
		fmt.Println(i)
		resultado = resultado + "|\n"
	}
	return resultado
}
