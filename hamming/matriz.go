package hamming

import (
	"fmt"
)

//Matriz estrucutra para almacenar bits en matrices
type Matriz struct {
	datos [][]bool
}

//ToString funcion que convierte una matriz en un string para imprimir en consola
func (operando *Matriz) ToString() string {
	var resultado string
	fmt.Println(fmt.Sprintf("M: %v, N: %v", len(operando.datos), len(operando.datos[0])))
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

//Multiplicar Funcion que multiplica dos matrices y devuelve sus resultado
func (operando *Matriz) Multiplicar(m Matriz) Matriz {
	if operando.datos != nil || operando.datos[0] != nil || m.datos != nil || m.datos[0] != nil {
		fmt.Println(fmt.Sprintf("M- Ancho: %v, Alto: %v ", len(operando.datos), len(operando.datos[0])))
		fmt.Println(fmt.Sprintf("N-Ancho: %v, Alto: %v", len(m.datos), len(m.datos[0])))

		if len(operando.datos) == len(m.datos[0]) {
			//La matriz resultante es de la forma rW x rH
			rW := len(operando.datos)
			rH := len(m.datos[0])
			aux := make([][]bool, rW)
			for i := range operando.datos {
				aux[i] = make([]bool, rH)
			}
			for k := 0; k < rW; k++ {
				for j := 0; j < rH; j++ {
					for i := 0; i < rW; i++ {
						aux[k][j] = (operando.datos[k][i] && m.datos[i][j] != aux[k][j])
					}
				}
			}
			v := Matriz{datos: aux}
			return v
		}
		fmt.Println("Error:El ancho y el alto de las matrices no concuerdan")
	} else {
		fmt.Println("Error:no hay datos cargados en las matrices.")
	}
	c := [][]bool{{}}
	var v = Matriz{datos: c}
	return v
}

//TieneUnos controla si algun valor de la matriz es igual a 1
func (operando *Matriz) TieneUnos() bool {
	for i := range operando.datos {
		for j := range operando.datos[i] {
			if operando.datos[i][j] {
				return true
			}
		}
	}
	return false
}

func NuevaMatriz(ancho int, alto int) *Matriz{
	aux:= make([][]bool ,ancho)
	for i range aux{
		aux[i] := make([]bool, alto)
	}
	return Matriz{datos:aux}
}
