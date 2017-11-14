package hamming

import (
	"math"
	"fmt"
)

//Matriz estructura para almacenar bits en matrices
type Matriz struct {
	datos [][]bool
}

//Matriz32 estrutura para almacenar bits en enteros
type Matriz32 struct{
	datos [][]uint32
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

//ToString convierte una matriz en su representacion en String
func (operando *Matriz32) ToString() string{
	var resultado string
	fmt.Println(fmt.Sprintf("M: %v, N: %v", len(operando.datos[0]),len(operando.datos)))
	for _, r := range operando.datos {
		resultado = resultado + "|"
		for _, d := range r {
			resultado += fmt.Sprintf("%v", d)
		}
		resultado = resultado + "|\n"
	}
	return resultado
}

//ToStringB funcion que convierte una matriz en un string para imprimir en consola
// donde como parametro toma el tamaño maximo de las filas
func (operando *Matriz32) ToStringB(size int) string {
	var resultado string
	fmt.Println(fmt.Sprintf("M: %v, N: %v", len(operando.datos[0]),len(operando.datos)))
	for _, r := range operando.datos {
		resultado = resultado + "|"
		for _, d := range r {
			aux :=  fmt.Sprintf("%b",d)
			for i:=len(aux); i <size; i ++ {
				resultado+= "0"
			}
			resultado +=aux
		}
		resultado = resultado + "|\n"
	}
	return resultado
}


//Multiplicar Funcion que multiplica dos matrices y devuelve sus resultado
func (operando *Matriz) Multiplicar(mE *Matriz)( bool,Matriz ){
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
			return false,*aux
		}
		fmt.Println("Error:El ancho y el alto de las matrices no concuerdan")
	} else {
		fmt.Println("Error:no hay datos cargados en las matrices.")
	}
	c := [][]bool{{}}
	var v = Matriz{datos: c}
	return true,v
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
//TieneUnos controla si algun valor de la matriz es igual a 1
func (operando *Matriz32) TieneUnos() bool {
	for i := range operando.datos {
		for j := range operando.datos[i] {
			if operando.datos[i][j] != 0 {
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

//NuevaMatriz32 funcion que genera la matriz vacia de tal ancho y alto
func NuevaMatriz32(ancho int,alto int) *Matriz32{
	aux := make([][]uint32 , ancho)
	for i:= range aux{
		aux[i] = make([]uint32, alto)
	}
	m:= Matriz32{datos:aux}
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

//MatrizFila a partir de una cadena de enteros de 32 bits genera la matriz fila
func MatrizFila(entrada []uint32) *Matriz32{
	aux:= NuevaMatriz32(1,len(entrada))
	for i,b := range entrada{
		aux.datos[0][i] = b
	}
	return aux
}


//MatrizTo32 convierte una matriz de booleanos en una de 64 bits
func (entrada *Matriz)MatrizTo32()*Matriz32{
	ancho := len(entrada.datos[0])/64
	auxM := NuevaMatriz32(len(entrada.datos),ancho)

	for indice,fila := range entrada.datos{
		for indiceC, celda :=range fila{
			if celda {
				f :=math.Pow(2, float64(len(fila)-1-indiceC))
				auxM.datos[indice][indiceC/64] = uint32(f)| auxM.datos[indice][indiceC/64]
			}	
		}
	}
	return auxM
}