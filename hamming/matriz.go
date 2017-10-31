package hamming

import (
	"fmt"
)

type matriz struct {
	datos [][]bool
}

func Matriz() {
	k := [][]bool{{false, true, true}, {true, true, false}}
	ma := matriz{datos: k}
	fmt.Println("Matriz")
	fmt.Println(ma.ToString())
}
func (this *matriz) Multiplicar(n matriz) matriz {
	return n
}

//Funcion que convierte una matriz en un string para imprimir en consola
func (this *matriz) ToString() string {
	var resultado string
	for _, r := range this.datos {
		resultado = resultado + "|"
		for _, d := range r {
			if d {
				resultado = resultado + "1"
			} else {
				resultado = resultado + "0"
			}
		}
		resultado = resultado + "|\n"
	}
	return resultado
}

//Funcion que multiplica dos matrices y devuelve sus resultado
func (this *matriz) multiplicar(m *matriz) matriz {
	var v matriz
	return v
}
