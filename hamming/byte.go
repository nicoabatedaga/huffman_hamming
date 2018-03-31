package hamming

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	bits "math/bits"
	"os"
	"sync"
)

//MatrizB estrutura para trasnformar los arreglos de byte's
type MatrizB struct {
	datos [][]byte
}

//TieneUnos retorna verdad si la matriz tiene algun 1, falso si son todos 0's
func (m MatrizB) TieneUnos() bool {
	for _, f := range m.datos {
		for _, c := range f {
			if c != 0 {
				return true
			}
		}
	}
	return false
}

//Set le da a una posicion de bits el valor 1
func (m MatrizB) Set(i, j int) {
	alto, ancho := m.obtenerTam()
	if i < alto && j < ancho*8 {
		mascara := []byte{128, 64, 32, 16, 8, 4, 2, 1}
		m.datos[i][j/8] |= mascara[j%8]
	} else {
		fmt.Println("ERROR: seteo un numero fuera de la matriz", alto, " ", ancho*8, " ", i, " ", j, " ")
		panic("Error: set")
	}
}

//NuevaMatrizB funcion que crea la matrizByte y le asigna
//espacio dado a su ancho x alto
func NuevaMatrizB(alto, ancho int) MatrizB {
	aux := make([][]byte, alto)
	for i := range aux {
		aux[i] = make([]byte, ancho)
	}
	m := MatrizB{datos: aux}
	return m
}

//MatrizFila apartir de una cadena de bytes crea la matriz fila
func MatrizFila(arregloEntrada []byte) MatrizB {
	dat := make([][]byte, 1)
	dat[0] = arregloEntrada
	m := MatrizB{datos: dat}
	return m
}

//MutiplicarByte multiplicar una matriz contra una matriz fila
func MultiplicarByte(matrizEntrada, mE MatrizB) (MatrizB, error) {
	altoM, anchoM := matrizEntrada.obtenerTam()
	altoE, anchoE := mE.obtenerTam()
	if altoE != 1 {
		dato := make([][]byte, 1)
		m := MatrizB{datos: dato}
		return m, fmt.Errorf("error: la matriz no es una matriz fila")
	}
	mascara := []byte{128, 64, 32, 16, 8, 4, 2, 1}
	if anchoE == anchoM {
		ancho := (altoM) / 8
		if altoM%8 != 0 {
			ancho++
		}
		aux := NuevaMatrizB(1, ancho)
		for k := 0; k < altoM; k++ {
			valor := matrizEntrada.datos[k][0] & mE.datos[0][0]
			for i := 1; i < anchoM; i++ {
				valor ^= matrizEntrada.datos[k][i] & mE.datos[0][i]
			}
			if bits.OnesCount8(valor)%2 == 1 {
				aux.datos[0][k/8] |= mascara[k%8]
			}
		}
		return aux, nil
	}
	dato := make([][]byte, 1)
	m := MatrizB{datos: dato}
	return m, fmt.Errorf("error: los anchos no coinciden %v %v %v %v", altoM, anchoM, altoE, anchoE)
}

const numbProc int = 8

//MultiplicarByteO multiplicar una matriz contra una matriz fila
func MultiplicarByteO(matrizEntrada, mE MatrizB) (MatrizB, error) {
	altoM, anchoM := matrizEntrada.obtenerTam()
	altoE, anchoE := mE.obtenerTam()
	if altoE != 1 {
		dato := make([][]byte, 1)
		m := MatrizB{datos: dato}
		return m, fmt.Errorf("error: la matriz no es una matriz fila")
	}
	if anchoE == anchoM {
		ancho := (altoM) / 8
		if altoM%8 != 0 {
			ancho++
		}
		aux := NuevaMatrizB(1, ancho)

		mascara := []byte{128, 64, 32, 16, 8, 4, 2, 1}
		var wg sync.WaitGroup
		var mu sync.Mutex
		wg.Add(numbProc)
		for g := 0; g < numbProc; g++ {
			go func(g int) {
				for k := g * altoM / numbProc; k < (1+g)*altoM/numbProc; k++ {
					valor := matrizEntrada.datos[k][0] & mE.datos[0][0]
					for i := 1; i < anchoM; i++ {
						valor ^= matrizEntrada.datos[k][i] & mE.datos[0][i]
					}
					if bits.OnesCount8(valor)%2 == 1 {
						mu.Lock()
						aux.datos[0][k/8] |= mascara[k%8]
						mu.Unlock()
					}
				}
				wg.Done()
			}(g)
		}
		wg.Wait()

		return aux, nil
	}
	dato := make([][]byte, 1)
	m := MatrizB{datos: dato}
	return m, fmt.Errorf("error: los anchos no coinciden %v %v %v %v", altoM, anchoM, altoE, anchoE)
}

//MultiplicarByteFila multiplicar una matriz contra una matriz fila
func MultiplicarByteFila(matrizEntrada, mE MatrizB) (MatrizB, error) {
	altoM, anchoM := matrizEntrada.obtenerTam()
	altoE, anchoE := mE.obtenerTam()
	if altoE != 1 {
		dato := make([][]byte, 1)
		m := MatrizB{datos: dato}
		return m, fmt.Errorf("error: la matriz no es una matriz fila")
	}
	if anchoE == anchoM {
		ancho := (altoM) / 8
		if altoM%8 != 0 {
			ancho++
		}
		aux := NuevaMatrizB(1, ancho)

		var wg sync.WaitGroup
		wg.Add(altoM)
		for k := 0; k < altoM; k++ {
			go func(k int) {
				valor := matrizEntrada.datos[k][0] & mE.datos[0][0]
				for i := 1; i < anchoM; i++ {
					valor ^= matrizEntrada.datos[k][i] & mE.datos[0][i]
				}
				if bits.OnesCount8(uint8(valor))%2 != 1 {
					aux.Set(0, k)
				}
				wg.Done()
			}(k)
		}
		wg.Wait()
		return aux, nil
	}
	dato := make([][]byte, 1)
	m := MatrizB{datos: dato}
	return m, fmt.Errorf("error: los anchos no coinciden %v %v %v %v", altoM, anchoM, altoE, anchoE)
}

func (m MatrizB) obtenerTam() (int, int) {
	alto := len(m.datos)
	if alto != 0 {
		return alto, len(m.datos[0])
	}
	return 0, 0
}

func contarUnos(ent byte) int {
	return bits.OnesCount8(uint8(ent))
}

//ToImage graba la imagen en una imagen
func (m MatrizB) ToImage(url string, expAltura int) {
	ancho := len(m.datos[0])
	alto := len(m.datos)
	imagen := image.NewRGBA(image.Rect(0, 0, alto*expAltura, ancho*8))
	for i, f := range m.datos {
		for j, c := range f {
			for b := 0; b < 8; b++ {
				cB := c & (uint8(1) << uint8(b))
				if cB == 0 {
					cB = 255
				} else {
					cB = 0
					//cB = cB - 254

				}
				col := color.RGBA{R: cB, G: 255, B: 255, A: 255}
				for k := 0; k < expAltura; k++ {
					imagen.Set(i*expAltura+k, j*8+b, col)
				}

			}

		}
	}

	file, err := os.Create(url)
	if err != nil {
		fmt.Println("Error:", err)
	}

	jpeg.Encode(file, imagen, nil)
	file.Close()
}

//ByteToBool convierte un arreglo de byte's en uno de booleanos
func ByteToBool(entrada []byte) []bool {
	salida := make([]bool, len(entrada)*8)
	mascara := []byte{128, 64, 32, 16, 8, 4, 2, 1}
	for indice, b := range entrada {
		for j := 0; j < 8; j++ {
			salida[indice*8+j] = ((b & mascara[j]) != 0)
		}
	}
	return salida

}
