package hamming

import (
	"fmt"
	"log"
	"os"
	"sync"
)

//Matriz estructura para almacenar bits en matrices
type Matriz struct {
	datos [][]bool
}

//ToByte transforma una matriz en un arreglo de bytes (solo funciona para las matrices columna)
func (matrizEntrada Matriz) ToByte() []byte {
	ancho := len(matrizEntrada.datos)
	alto := len(matrizEntrada.datos[0])
	tam := ancho * alto / 8
	if (ancho*alto)%8 != 0 {
		tam++
	}
	auxB := make([]byte, tam)
	for indice := range auxB {
		var auxByte byte
		mascara := []byte{1, 2, 4, 8, 16, 32, 64, 128}
		for i, n := range mascara {
			if indice*8+i < ancho {
				if matrizEntrada.datos[indice*8+i][0] {
					auxByte = auxByte | n
				}
			}
		}
		auxB[indice] = auxByte
	}
	return auxB
}

//Multiplicar Funcion que multiplica dos matrices y devuelve sus resultado
func (matrizEntrada Matriz) Multiplicar(mE Matriz) (bool, Matriz) {
	if matrizEntrada.datos != nil || matrizEntrada.datos[0] != nil || mE.datos != nil || mE.datos[0] != nil {
		nO := len(matrizEntrada.datos[0])
		mEn := len(mE.datos)

		if nO == mEn {
			m := len(matrizEntrada.datos) //alto
			n := len(mE.datos[0])         //ancho
			aux := NuevaMatriz(m, n)
			for k := 0; k < m; k++ { //indiceFila = k
				for j := 0; j < n; j++ { //indiceColumna =j
					for i := 0; i < nO; i++ { //indice = i
						aux.datos[k][j] =
							((matrizEntrada.datos[k][i] && mE.datos[i][j]) != aux.datos[k][j])
					}
				}
			}
			return false, aux
		}

		fmt.Println("Error:El ancho y el alto de las matrices no concuerdan", nO, mEn)
	} else {
		fmt.Println("Error:no hay datos cargados en las matrices.")
	}
	c := [][]bool{{}}
	var v = Matriz{datos: c}
	return true, v
}

//MultiplicarOpt Funcion que multiplica dos matrices y devuelve sus resultado
func (matrizEntrada Matriz) MultiplicarOpt(mE Matriz) (Matriz, error) {
	labError := "Error:no hay datos cargados en las matrices."

	if matrizEntrada.datos != nil || matrizEntrada.datos[0] != nil || mE.datos != nil || mE.datos[0] != nil {
		nO := len(matrizEntrada.datos[0])
		mEn := len(mE.datos)
		if nO == mEn {
			m := len(matrizEntrada.datos) //alto
			n := len(mE.datos[0])         //ancho
			var w sync.WaitGroup
			w.Add(m)
			aux := NuevaMatriz(m, n)
			for k := 0; k < m; k++ { //indiceFila = k
				go func(k int) {
					for j := 0; j < n; j++ { //indiceColumna =j
						for i := 0; i < nO; i++ { //indice = i
							valor := matrizEntrada.datos[k][i] && mE.datos[i][j]
							aux.xor(k, j, valor)
						}
					}
					w.Done()
				}(k)

			}
			w.Wait()
			return aux, nil
		}

		labError = fmt.Sprint("Error:El ancho y el alto de las matrices no concuerdan", nO, mEn)
	}

	c := [][]bool{{}}
	var v = Matriz{datos: c}
	return v, fmt.Errorf(labError)
}

func (matrizEntrada Matriz) xor(k, j int, valor bool) {
	valorV := matrizEntrada.datos[k][j]
	matrizEntrada.datos[k][j] = valor != valorV
}

//TieneUnos controla si algun valor de la matriz es igual a 1
func (matrizEntrada Matriz) TieneUnos() bool {
	for i := range matrizEntrada.datos {
		for j := range matrizEntrada.datos[i] {
			if matrizEntrada.datos[i][j] {
				return true
			}
		}
	}
	return false
}

//NuevaMatriz funcion que crea la matriz y le asigna espacio dato a su ancho x alto
func NuevaMatriz(ancho int, alto int) Matriz {
	aux := make([][]bool, ancho)
	for i := range aux {
		aux[i] = make([]bool, alto)
	}
	m := Matriz{datos: aux}
	return m
}

//MatrizColumna apartir de una cadena de bytes crea la matriz columna
func MatrizColumna(matrizEntrada []bool) Matriz {
	dat := make([][]bool, len(matrizEntrada))
	for i, b := range matrizEntrada {
		dat[i] = make([]bool, 1)
		dat[i][0] = b
	}
	m := Matriz{datos: dat}
	return m
}

//ToString funcion que convierte una matriz en un string para imprimir en consola
func (matrizEntrada Matriz) ToString() string {
	var resultado string
	var aux uint64
	var contador uint64
	for _, r := range matrizEntrada.datos {
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
			resultado = resultado + fmt.Sprintf("%v", aux) //\n
			aux = 0
			contador = 0
		}
	}
	return resultado
}

//ToStringConInfo funcion que convierte una matriz en un string para imprimir en consola
func (matrizEntrada Matriz) ToStringConInfo() string {
	var resultado string
	fmt.Println(fmt.Sprintf("M: %v, N: %v", len(matrizEntrada.datos), len(matrizEntrada.datos[0])))
	for _, r := range matrizEntrada.datos {
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

//ByteToBool convierte un arreglo de byte's en uno de booleanos
func ByteToBool(entrada []byte) []bool {
	auxB := make([]bool, len(entrada)*8)
	for i, b := range entrada {
		m := []byte{1, 2, 4, 8, 16, 32, 64, 128}
		for j := 0; j < 8; j++ {
			auxB[i*8+j] = ((b & m[j]) != 0)
		}
	}
	return auxB

}

//ToFile graba una Matriz en un archivo binario
func (matrizEntrada Matriz) ToFile(url string) {
	file, err := os.Create(url)
	manejoError(err)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	alto := len(matrizEntrada.datos)
	ancho := len(matrizEntrada.datos[0])
	_, err = file.WriteString(fmt.Sprintf("%v\n%v\n", alto, ancho))
	manejoError(err)
	_, err = file.WriteString(matrizEntrada.ToString())
	manejoError(err)
	return
}

//ToFileConInfo graba una Matriz en un archivo binario retornando a la vez la información
func (matrizEntrada Matriz) ToFileConInfo(url string) {
	file, err := os.Create(url)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString(matrizEntrada.ToStringConInfo())
	return
}

//ByteToMatriz genera una matriz columna apartir de un slice de bytes
func ByteToMatriz(entrada []byte) Matriz {
	entradaBool := ByteToBool(entrada)
	matrizEntrada := MatrizColumna(entradaBool)
	return matrizEntrada
}
