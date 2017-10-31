package hamming

import (
	"fmt"
)

type matriz struct {
	datos [][]bool
}

//Matriz funcion de prueba
func Matriz() {
	k := [][]bool{{false, true, true}, {true, true, false}}
	ma := matriz{datos: k}
	fmt.Println("Matriz")
	fmt.Println(ma.ToString())
}

//Multiplicar sirve para multiplicar dos matrices
func (operando *matriz) Multiplicar(n matriz) matriz {
	return n
}

//ToString funcion que convierte una matriz en un string para imprimir en consola
func (operando *matriz) ToString() string {
	var resultado string
	fmt.Println(fmt.Sprintf("M: %v, N: %v", len(operando.datos), len(operando.datos)))
	for _, r := range operando.datos {
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
func (operando *matriz) multiplicar(m *matriz) matriz {
	var v matriz
	return v
}
