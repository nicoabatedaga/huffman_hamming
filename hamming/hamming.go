package hamming

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

//Hamming es el programa de prueba
func Hamming() {
	start := time.Now()
	fmt.Println("-Hamming-")

	//codificacion := 522
	//codificacion = 1035
	codificacion := 2060

	fmt.Println(fmt.Sprintf("Codificacion: %v, Bits Paridad: %v, Bits Información: %v",
		codificacion,
		bitsParidad(codificacion),
		bitsInformacion(codificacion)))

	//pruebaHGR(codificacion)
	Proteger("prueba.txt", "prueba.ham", 522)
	elapsed := time.Since(start)
	fmt.Println(fmt.Sprintf("Estuvo ejecutando: %s", elapsed))
	return
}

func manejoError(err error) {
	if err != nil {
		panic(fmt.Sprintf("hubo un error, %v", err))
	}
}

func pruebaHGR(codificacion int) {
	fmt.Println("H:")
	h := h(codificacion)
	h.ToFile("h2048.bin")

	fmt.Println("G:")
	g := g(codificacion)
	g.ToFile("g2048.bin")

	fmt.Println("R:")
	r := r(codificacion)
	r.ToFile("r2048.bin")

	fmt.Println("Entrada:")
	c := matrizEntradaPrueba(codificacion)
	c.ToFile("entrada.bin")

	error, codificada := g.Multiplicar(c)

	AgregarError(&codificada)
	error, posicion := TieneError(&codificada)
	fmt.Println(fmt.Sprintf("Tienen errores:%v en %v", error, posicion))

	CorregirError(&codificada)
	error, posicion = TieneError(&codificada)
	fmt.Println(fmt.Sprintf("Tienen errores:%v en %v", error, posicion))

	fmt.Println("Corregida:")

	fmt.Println("Decodificada:")
	error, decodificada := Decodificar(&codificada)
	if error {
		fmt.Println("Error al multiplicar por R")
	}
	decodificada.ToFile("decodificada.bin")

}

func esPotenciaDeDos(ent int) bool {
	return ((ent != 0) && ((ent & (ent - 1)) == 0))
}

func bitsInformacion(ent int) int {
	return (ent - (bitsParidad(ent)))
}

//bitsParidad Devuelve que candidad de bit corresponderian a bits de paridad para
//determinada codifciacion, si se desea para evitar iterar, podria hacerse un mapeo
//con los 5 valores de codificacion que exiten.
func bitsParidad(ent int) int {
	n := 0
	for i := 1; i <= ent; i = i * 2 {
		n++
	}
	if ent == 7 {
		return 3
	}
	if ent == 31 {
		return 5
	}
	return n
}

//h metodo que devuelve la matriz que se multiplica para codificar una entrada
func h(codificacion int) Matriz {
	alto := codificacion
	ancho := bitsParidad(codificacion)
	aux := NuevaMatriz(ancho, alto)
	for i := 0; i < ancho; i++ {
		b := 1
		uno := false
		for j := 0; j < alto; j++ {
			f := float64(i)
			if int(math.Pow(2, f)) == b {
				b = 0
				uno = !uno
			}
			b++
			if uno {
				aux.datos[i][j] = true
			}
		}
	}
	return *aux
}

//g Funcion que crea la matriz generadora
func g(codificacion int) *Matriz {
	n := codificacion
	m := bitsInformacion(codificacion)
	aux := NuevaMatriz(n, m)
	k := 0
	p := -1
	for i := 0; i < n; i++ {
		if !esPotenciaDeDos(i + 1) {
			aux.datos[i][k] = true
			k++
		} else {
			p++
			r := 1
			uno := false
			b := 1
			for j := 0; j < m; j++ {
				for esPotenciaDeDos(r) {
					r++
					f := float64(p)
					if b == int(math.Pow(2, f)) {
						uno = !uno
						b = 0
					}
					b++
				}
				f := float64(p)
				if b == int(math.Pow(2, f)) {
					uno = !uno
					b = 0
				}
				b++
				if uno {
					aux.datos[i][j] = true
				}
				r++
			}
		}
	}
	return aux
}

//r Funcion que crea la matriz decodificadora
func r(codificacion int) *Matriz {
	n := bitsInformacion(codificacion)
	m := codificacion
	aux := NuevaMatriz(n, m)
	k := 0
	for i := 0; i < m; i++ {
		if !esPotenciaDeDos(i + 1) {
			aux.datos[k][i] = true
			k++
		}
	}
	return aux
}

//Codificar  Cifra el archivo de entrada
func Codificar(operando *Matriz) (bool, Matriz) {
	cod := len(operando.datos)
	h := h(cod)
	error, codificada := h.Multiplicar(operando)
	if error {
		return true, *operando
	}
	return false, codificada
}

//Proteger funcion que toma de entrada el path de un archivo y lo codifica segun un valor de entrada
func Proteger(url string, salida string, codificacion int) bool {
	if _, err := os.Stat(url); os.IsNotExist(err) {
		return false
	}
	file, err := os.Open(url)
	manejoError(err)
	defer file.Close()

	fileO, err := os.OpenFile(salida, os.O_WRONLY, 0666)
	manejoError(err)
	defer fileO.Close()

	bufferReader := bufio.NewReader(file)

	buf := make([]byte, bitsInformacion(codificacion)/8)
	h := h(codificacion)

	byteLeidos, err := bufferReader.Read(buf)

	manejoError(err)
	for byteLeidos > 0 {
		auxMatriz := MatrizColumna(ByteToBool(buf))
		b, m := h.Multiplicar(auxMatriz)
		if !b {
			fileO.Write(m.ToString())
		}

	}

	h.ToFile("h.bin")
	return true
}

//TieneError verifica si tiene error una matriz, y devuelve la posicion del mismo.
func TieneError(operando *Matriz) (bool, int) {
	codificacion := len(operando.datos)
	h := h(codificacion)
	error, sindrome := h.Multiplicar(operando)
	if error {
		fmt.Println("Error al multiplicar")
		return true, -1
	}
	resultB := sindrome.TieneUnos()
	resultI := -1
	for i, fila := range sindrome.datos {
		for _, b := range fila {
			if b {
				f := float64(i)
				resultI = resultI + int(math.Pow(2, f))
			}
		}
	}

	return resultB, resultI
}

//Decodificar descrifra el archivo de entrada
func Decodificar(operando *Matriz) (bool, *Matriz) {
	cod := len(operando.datos)
	r := r(cod)
	error, decodificada := r.Multiplicar(operando)
	if error {
		return true, operando
	}
	return false, &decodificada

}

//AgregarError a un archivo, si devuelve true es porque agrego error, sino ya habia antes un error
func AgregarError(operando *Matriz) bool {
	tieneError, _ := TieneError(operando)
	if tieneError {
		return false
	}
	n := len(operando.datos)
	m := len(operando.datos[0])
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r.Intn(n)
	j := r.Intn(m)
	operando.datos[i][j] = !operando.datos[i][j]
	return true
}

//CorregirError toma como entrada una matriz, si la corrige retorna verdadero, sino retorna false
func CorregirError(operando *Matriz) bool {
	tieneError, posicion := TieneError(operando)
	if tieneError {
		if posicion != -1 {
			operando.datos[posicion][0] = !operando.datos[posicion][0]
			return true
		}
	}
	return false
}

//matrizEntradaPrueba genera una matriz de entrada aleatoria del tamaño de la codificacion
func matrizEntradaPrueba(codificacion int) *Matriz {
	aux := make([]bool, bitsInformacion(codificacion))
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range aux {
		j := r.Intn(100)
		aux[i] = (j%2 == 1)
	}
	m := MatrizColumna(aux)
	return m
}
