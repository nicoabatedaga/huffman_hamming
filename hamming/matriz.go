package hamming

import (
	"fmt"
	"log"
	"os"
)

//Matriz estructura para almacenar bits en matrices
type Matriz struct {
	datos [][]bool
}

//ToByte transforma una matriz en un arreglo de bytes
func (operando *Matriz) ToByte() []byte {
	ancho := len(operando.datos)
	alto := len(operando.datos[0])
	auxB := make([]byte, ancho*alto/8)
	for indice, _ := range auxB { 
		var auxByte byte 
		if operando.datos[indice*8][0 ]{ 
			auxByte = auxByte | 1

		}
		if operando.datos[indice*8 +1][0]{
			auxByte= auxByte |2
		}
		if operando.datos[indice*8 +2][0]{
			auxByte = auxByte | 4
		}
		if operando.datos[indice*8 +3][0]{
			auxByte = auxByte | 8
		}
		if operando.datos[indice*8 +4][0]{
			auxByte = auxByte | 16
		}
		if operando.datos[indice*8 + 5][0]{
			auxByte = auxByte | 32
		}
		if operando.datos[indice*8 + 6][0]{
			auxByte = auxByte | 64
		}
		if operando.datos[indice*8 +7][0]{
			auxByte = auxByte | 128
		}
		auxB[indice] = auxByte
	}
	return auxB
}

//ToString funcion que convierte una matriz en un string para imprimir en consola
func (operando *Matriz) ToString() string {
	var resultado string
	var aux uint64
	var contador uint64
	for _, r := range operando.datos {
		for _, d := range r {
			if d {
				aux = aux | (1 << contador)
			}
			contador++
			if contador%64 == 0 {
				resultado = resultado + fmt.Sprintf("%v", aux)
				aux = 0
				contador = 0
			}
		}
		if contador != 0 {
			resultado = resultado + fmt.Sprintf("%v", aux)
			aux = 0
			contador = 0
		}

		//resultado = resultado + "\n"
	}
	return resultado
}

//ToStringConInfo funcion que convierte una matriz en un string para imprimir en consola
func (operando *Matriz) ToStringConInfo() string {
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
func (operando *Matriz) Multiplicar(mE *Matriz) (bool, Matriz) { 
	if operando.datos != nil || operando.datos[0] != nil || mE.datos != nil || mE.datos[0] != nil {
		nO := len(operando.datos[0])
		mEn := len(mE.datos)

		if nO == mEn {
			m := len(operando.datos)
			n := len(mE.datos[0])
			aux := NuevaMatriz(m, n)
			for k := 0; k < m; k++ {
				for j := 0; j < n; j++ {
					for i := 0; i < nO; i++ {
						aux.datos[k][j] = ((operando.datos[k][i] && mE.datos[i][j]) != aux.datos[k][j])
					}
				}
			} 
			return false, *aux
		}

		fmt.Println("Error:El ancho y el alto de las matrices no concuerdan", nO, mEn)
	} else {
		fmt.Println("Error:no hay datos cargados en las matrices.")
	}
	c := [][]bool{{}}
	var v = Matriz{datos: c}
	return true, v
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
func NuevaMatriz(ancho int, alto int) *Matriz {
	aux := make([][]bool, ancho)
	for i := range aux {
		aux[i] = make([]bool, alto)
	}
	m := Matriz{datos: aux}
	return &m
}

//MatrizColumna apartir de una cadena de bytes crea la matriz columna
func MatrizColumna(entrada []bool) *Matriz {
	aux := NuevaMatriz(len(entrada), 1)
	for i, b := range entrada {
		aux.datos[i][0] = b
	}
	return aux
}

//ToFile graba una Matriz en un archivo binario
func (entrada *Matriz) ToFile(url string) {
	file, err := os.Create(url)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString(entrada.ToString())
	return
}

//ByteToBool convierte un arreglo de byte's en uno de booleanos
func ByteToBool(entrada []byte) []bool {
	auxB := make([]bool, len(entrada)*8)
	for i, b := range entrada {
		auxB[i*8] = ((b & 1) != 0)
		auxB[i*8+1] = ((b & 2) != 0)
		auxB[i*8+2] = ((b & 4) != 0)
		auxB[i*8+3] = ((b & 8) != 0)
		auxB[i*8+4] = ((b & 16) != 0)
		auxB[i*8+5] = ((b & 32) != 0)
		auxB[i*8+6] = ((b & 64) != 0)
		auxB[i*8+7] = ((b & 128) != 0)
	}
	return auxB
}
