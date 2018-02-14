package hamming

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

//Matriz estructura para almacenar bits en matrices
type Matriz struct {
	datos [][]bool
}

//ToByte transforma una matriz en un arreglo de bytes (solo funciona para las matrices columna)
func (matrizEntrada *Matriz) ToByte() []byte {
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

//ToString funcion que convierte una matriz en un string para imprimir en consola
func (matrizEntrada *Matriz) ToString() string {
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
			resultado = resultado + fmt.Sprintf("%v\n", aux)
			aux = 0
			contador = 0
		}
	}
	return resultado
}

//GenerarMatrizViaString apartir de un String, su ancho y alto genera la matriz correspodiente, estando codificada en 64 bits la matriz
/*func GenerarMatrizViaString(cadenaMatriz string, ancho, alto int) Matriz {
	matriz := NuevaMatriz(ancho, alto)
	bits := strings.Split(cadenaMatriz, " ")
	for indiceAlto := 0; indiceAlto < alto; indiceAlto++ {
		for indiceAncho := 0; indiceAncho < ancho; indiceAncho++ {

		}
	}
	return *matriz
}
*/
//ToStringConInfo funcion que convierte una matriz en un string para imprimir en consola
func (matrizEntrada *Matriz) ToStringConInfo() string {
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

//Multiplicar Funcion que multiplica dos matrices y devuelve sus resultado
func (matrizEntrada *Matriz) Multiplicar(mE *Matriz) (bool, Matriz) {
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

//MultiplicarOpt Funcion que multiplica dos matrices y devuelve sus resultado
func (matrizEntrada *Matriz) MultiplicarOpt(mE *Matriz) (bool, Matriz) {
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

func (matrizEntrada *Matriz) xor(k, j int, valor bool) {
	valorV := matrizEntrada.datos[k][j]
	matrizEntrada.datos[k][j] = valor != valorV
}

//TieneUnos controla si algun valor de la matriz es igual a 1
func (matrizEntrada *Matriz) TieneUnos() bool {
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
func NuevaMatriz(ancho int, alto int) *Matriz {
	aux := make([][]bool, ancho)

	for i := range aux {
		aux[i] = make([]bool, alto)
	}
	m := Matriz{datos: aux}
	return &m
}

//MatrizColumna apartir de una cadena de bytes crea la matriz columna
func MatrizColumna(matrizEntrada []bool) *Matriz {
	dat := make([][]bool, len(matrizEntrada))
	for i, b := range matrizEntrada {
		dat[i] = make([]bool, 1)
		dat[i][0] = b
	}
	m := Matriz{datos: dat}
	return &m
}

//ToFile graba una Matriz en un archivo binario
func (matrizEntrada *Matriz) ToFile(url string) {
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

//ToMatriz carga la matriz desde un archivo
func ToMatriz(url string) Matriz {
	archivo, err := os.Open(url)
	manejoError(err)
	defer archivo.Close()
	bufferReader := bufio.NewReader(archivo)
	line, err := bufferReader.ReadString('\n')
	manejoError(err)
	alto, err := strconv.Atoi(line[:len(line)-1])
	manejoError(err)
	line, err = bufferReader.ReadString('\n')
	manejoError(err)
	ancho, err := strconv.Atoi(line[:len(line)-1])
	manejoError(err)
	matriz := *NuevaMatriz(ancho, alto)
	for indiceAlto := 0; indiceAlto < alto; indiceAlto++ {
		line, err = bufferReader.ReadString('\n')
		manejoError(err)
		for indiceAncho, caracter := range line {
			valor, err := strconv.Atoi(string(caracter))
			manejoError(err)
			matriz.datos[indiceAlto][indiceAncho] = valor == 1
		}
	}
	return matriz
}

//ToFileConInfo graba una Matriz en un archivo binario retornando a la vez la información
func (matrizEntrada *Matriz) ToFileConInfo(url string) {
	file, err := os.Create(url)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString(matrizEntrada.ToStringConInfo())
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

func compararMatrices(matriz1, matriz2 Matriz) bool {
	ancho1 := len(matriz1.datos)
	ancho2 := len(matriz2.datos)
	alto1 := len(matriz1.datos[0])
	alto2 := len(matriz2.datos[0])
	if ancho1 != ancho2 || alto1 != alto2 {
		return false
	}
	return true
}
