package hamming

import (
	"fmt"
)

//MatrizBytes estructura para almacenar bits en matrices de bytes
type MatrizBytes struct {
	datos [][]byte
}

//NuevaMatrizBytes genera la matriz de bytes del ancho y alto correspondiente
func NuevaMatrizBytes(ancho int, alto int) *MatrizBytes {
	arregloBytes := make([][]byte, alto)
	for indice := range arregloBytes {
		arregloBytes[indice] = make([]byte, ancho)
	}
	return &MatrizBytes{datos: arregloBytes}
}

//MatrizFila a partir de una slice de bytes genera la matriz fila correspondiente
func MatrizFila(bytesEntrada []byte) *MatrizBytes {
	matrizAux := NuevaMatrizBytes(len(bytesEntrada), 1)
	matrizAux.datos[0] = bytesEntrada
	return matrizAux
}

//Set coloca un valor a un bit, dado una fila y una posicion de bit
func (matriz MatrizBytes) Set(fila, posicionBit int, valor bool) error {

	if fila > matriz.Alto() || fila < 0 {
		return fmt.Errorf("error set:no corresponde a un indice de fila valido %v", fila)
	}
	if posicionBit > matriz.Ancho()*8 || posicionBit < 0 {
		fmt.Println("ERROR")
		return fmt.Errorf("error set:no corresponde a un indice de bit valido %v", posicionBit)
	}
	byte := matriz.datos[fila][posicionBit/8]
	byteAux, error := SetBit(byte, uint(posicionBit%8), valor)
	if error != nil {
		return fmt.Errorf("error set: %s", error)
	}
	matriz.datos[fila][posicionBit/8] = byteAux
	return nil
}

//Get retorna el  valor del bit, dado una fila y una posicion de bit
func (matriz MatrizBytes) Get(fila, posicionBit int) (bool, error) {
	if fila > matriz.Alto() || fila < 0 {
		return false, fmt.Errorf("error get:no corresponde a un indice de fila valido %v", fila)
	}
	if posicionBit > matriz.Ancho()*8 || posicionBit < 0 {
		return false, fmt.Errorf("error get:no corresponde a un indice de bit valido %v", posicionBit)
	}
	byte := matriz.datos[fila][posicionBit/8]
	valor, error := GetBit(byte, uint(posicionBit%8))
	if error != nil {
		return false, fmt.Errorf("error get: %s", error)
	}
	return valor, nil
}

//Xor realiza la operacion xor en un bit y setea el valor
func (matriz MatrizBytes) Xor(fila, posicionBit int, valor bool) error {

	if fila > matriz.Alto() || fila < 0 {
		return fmt.Errorf("error xor:no corresponde a un indice de fila valido %v", fila)
	}
	if posicionBit > matriz.Ancho()*8 || posicionBit < 0 {
		fmt.Println("ERROR")
		return fmt.Errorf("error xor:no corresponde a un indice de bit valido %v", posicionBit)
	}
	byte := matriz.datos[fila][posicionBit/8]
	valorViejo, error := GetBit(byte, uint(posicionBit%8))
	if error != nil {
		return fmt.Errorf("error xor:no se puede retornar valor anterior %s", error)
	}

	byteAux, error := SetBit(byte, uint(posicionBit%8), valor != valorViejo)
	if error != nil {
		return fmt.Errorf("error xor: %s", error)
	}
	matriz.datos[fila][posicionBit/8] = byteAux
	return nil
}

//SetBit setea un bit de una byte a un valor
func SetBit(ent byte, posicion uint, valor bool) (byte, error) {
	if posicion > 7 || posicion < 0 {
		return ent, fmt.Errorf("set bit:%b,la posicion %v excede el limite", ent, posicion)
	}
	movimiento := 7 - posicion
	mascaraBinaria := byte(1 << movimiento)
	if valor {
		return ent | mascaraBinaria, nil
	}
	return ent & ^mascaraBinaria, nil
}

//GetBit devuelve el valor del bit en la posicion correspondiente
func GetBit(ent byte, posicion uint) (bool, error) {
	if posicion > 7 || posicion < 0 {
		return false, fmt.Errorf("Error,get bit:%b,la posicion %v excede el limite", ent, posicion)
	}
	movimiento := 7 - posicion
	mascaraBinaria := byte(1 << movimiento)
	return ent&mascaraBinaria != 0, nil
}

//Multiplicar dos matrices, y retorna su resultado o error
func (matriz *MatrizBytes) Multiplicar(mE *MatrizBytes) (MatrizBytes, error) {
	ladoComunMatrices := matriz.Ancho() * 8
	c := [][]byte{{}}
	var matrizVacia = MatrizBytes{datos: c}
	if ladoComunMatrices != mE.Alto() {
		err := fmt.Errorf("los tamaños de las matrices no concuerdan para multiplicar\n alto:%v ancho:%v\nalto:%v ancho:%v", matriz.Alto(), ladoComunMatrices, mE.Alto(), mE.Ancho())
		return matrizVacia, err
	}
	alto := matriz.Alto()
	ancho := mE.Ancho()
	matrizResultante := NuevaMatrizBytes(ancho, alto)
	for indiceFila := 0; indiceFila < alto; indiceFila++ {
		for indiceColumna := 0; indiceColumna*8 < ancho; indiceColumna++ {
			for indice := 0; indice < ladoComunMatrices; indice++ {
				valorMatrizEntrada, error := matriz.Get(indiceFila, indice)
				if error != nil {
					err := fmt.Errorf("error multiplicar: %s", error)
					return matrizVacia, err
				}
				valorMatrizParamentro, error := mE.Get(indice, indiceColumna)
				if error != nil {
					err := fmt.Errorf("error multiplicar: %s", error)
					return matrizVacia, err
				}
				valor := valorMatrizEntrada && valorMatrizParamentro
				error = matrizResultante.Xor(indiceFila, indiceColumna, valor)
				if error != nil {
					err := fmt.Errorf("error multiplicar: %s", error)
					return matrizVacia, err
				}
			}
		}
	}
	return *matrizResultante, nil

}

//Alto devuelve el alto de la matriz
func (matriz *MatrizBytes) Alto() int {
	return len(matriz.datos)
}

//Ancho devuelve el ancho de la matriz
func (matriz *MatrizBytes) Ancho() int {
	return len(matriz.datos[0])
}

//ByToBo convierte un arreglo de byte's en uno de booleanos
func ByToBo(entrada []byte) []bool {
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

//MatrizColumnaByte apartir de una cadena de bytes crea la matriz columna
func MatrizColumnaByte(arregloBytes []byte) *MatrizBytes {
	alto := len(arregloBytes)
	matriz := NuevaMatrizBytes(1, alto*8)
	for indice := 0; indice < alto; indice++ {
		b := arregloBytes[indice]
		mascara := []byte{1, 2, 4, 8, 16, 32, 64, 128}
		for i, m := range mascara {
			matriz.Set(indice*8+i, 0, (b&m) != 0)
		}
	}
	return matriz
}

//ToByte transforma una matriz en un arreglo de bytes (solo funciona para las matrices columna)
func (matriz *MatrizBytes) ToByte() []byte {
	ancho := matriz.Ancho()
	alto := matriz.Alto()
	tam := ancho * alto / 8
	if (ancho*alto)%8 != 0 {
		tam++
	}
	arregloBytes := make([]byte, tam)
	for indice := range arregloBytes {
		var auxByte byte
		mascara := []byte{1, 2, 4, 8, 16, 32, 64, 128}
		for i, n := range mascara {
			if indice*8+i < alto {
				bo, err := matriz.Get(indice*8+i, 0)
				if err != nil {
					fmt.Printf("error ToByte: %s", err)
				}
				if bo {
					auxByte = auxByte | n

				}

			}
		}
		arregloBytes[indice] = auxByte
	}

	return arregloBytes
}

//TieneUnos controla si algun valor de la matriz es igual a 1
func (matriz *MatrizBytes) TieneUnos() bool {
	for i := range matriz.datos {
		for j := range matriz.datos[i] {
			if matriz.datos[i][j] != 0 {
				return true
			}
		}
	}
	return false
}
