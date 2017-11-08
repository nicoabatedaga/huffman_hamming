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
	fmt.Println(fmt.Sprintf("M: %v, N: %v", len(operando.datos[0]),len(operando.datos)))
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
func (operando *Matriz) Multiplicar(mE *Matriz) Matriz {
	if operando.datos != nil || operando.datos[0] != nil || mE.datos != nil || mE.datos[0] != nil {	
		nO := len(operando.datos[0])
		mEn := len(mE.datos)

		if nO == mEn {
			m := len(operando.datos)
			n := len(mE.datos[0])
			aux := NuevaMatriz(m,n)
			for k := 0; k < m; k++ {
				for j := 0; j < n; j++ {
					for i := 0; i < nO; i++ {
						aux.datos[k][j] = ((operando.datos[k][i] && mE.datos[i][j]) != aux.datos[k][j])
					}
				}
			}
			return *aux
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

//NuevaMatriz funcion que crea la matriz y le asigna espacio dato a su ancho x alto
func NuevaMatriz(ancho int, alto int) *Matriz{
	aux:= make([][]bool ,ancho)
	for i:=range aux{
		aux[i] = make([]bool, alto)
	}
	m:=Matriz{datos:aux}
	return &m
}

//MatrizColumna apartir de una cadena de bytes crea la matriz columna
func MatrizColumna(entrada []bool) *Matriz{
	aux:= NuevaMatriz(len(entrada),1)
	for i,b := range entrada{
		aux.datos[i][0] = b
	}
	return aux

}