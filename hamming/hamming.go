package hamming

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var paridad = map[int]int{
	528:  10,
	1040: 11,
	2060: 12,
}

//Hamming es el programa de prueba
func Hamming() {
	return
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
	p := paridad[ent]
	if p != 0 {
		return p
	}
	n := 0
	if ent == 7 {
		return 3
	}
	if ent == 31 {
		return 5
	}
	for i := 1; i <= ent; i = i * 2 {
		n++
	}
	paridad[ent] = n
	return n
}

//Proteger funcion que toma de entrada el path de un archivo y lo codifica segun un valor de entrada
func Proteger(url string, salida string, codificacion int) error {
	var archivoTemporal string
	if existeArchivo(url) {
		file, err := os.Open(url)
		manejoError(err)
		defer file.Close()
		fileTemp, err := ioutil.TempFile("", "./ham")
		manejoError(err)
		archivoTemporal = fileTemp.Name()

		fileOut, err := os.Create(salida)
		manejoError(err)

		fileOut.Close()
		fileOut, err = os.OpenFile(salida, os.O_WRONLY, 0666)
		manejoError(err)
		defer fileOut.Close()
		defer os.Remove(fileTemp.Name())

		bufferReader := bufio.NewReader(file)
		bufferWriter := bufio.NewWriter(fileTemp)
		buf := make([]byte, codificacion/8)

		byteLeidos, err := bufferReader.Read(buf)
		if byteLeidos != 0 {
			manejoError(err)
		}
		contadorBloques := 0
		for err != io.EOF {
			contadorBloques++
			codificadoByte := codificar(buf, byteLeidos)
			_, err = bufferWriter.Write(codificadoByte)
			manejoError(err)
			if byteLeidos < len(buf) {
				break
			}
			byteLeidos, err = bufferReader.Read(buf)
			if byteLeidos != 0 {
				manejoError(err)
			}
		}
		bufferWriter.Flush()
		fileTemp.Close()
		fileTemp, err = os.Open(archivoTemporal)
		manejoError(err)
		defer fileTemp.Close()
		return escribirArchivo(fileOut, fileTemp, codificacion, contadorBloques, byteLeidos)
	}
	return fmt.Errorf("El archivo no existe, %s", url)
}
func codificar(buf []byte, byteLeidos int) []byte {
	matrizG := matrizGeneradora(byteLeidos * 8)
	matrizEntrada := ByteToMatriz(buf[:byteLeidos])
	matrizResultante, err := matrizG.MultiplicarOpt(matrizEntrada)
	manejoError(err)
	codificadoByte := matrizResultante.ToByte()
	return codificadoByte
}

//matrizGeneradora Funcion que crea la matriz generadora
// tomando la cantidad de bits de informacion que hay
func matrizGeneradora(codificacion int) Matriz {
	anchoMatriz := codificacion + bitsParidad(codificacion)
	altoMatriz := codificacion
	matriz := NuevaMatriz(anchoMatriz, altoMatriz)
	indiceFilaIdentidad := 0
	p := -1
	for indiceColumna := 0; indiceColumna < anchoMatriz; indiceColumna++ {
		if !esPotenciaDeDos(indiceColumna + 1) {
			matriz.datos[indiceColumna][indiceFilaIdentidad] = true
			indiceFilaIdentidad++
		} else {
			p++
			f := float64(p)
			potencia := int(math.Pow(2, f))
			r := 1
			uno := false
			contadorBinario := 1
			for indiceFila := 0; indiceFila < altoMatriz; indiceFila++ {
				for esPotenciaDeDos(r) {
					r++
					if contadorBinario == potencia {
						uno = !uno
						contadorBinario = 0
					}
					contadorBinario++
				}
				if contadorBinario == potencia {
					uno = !uno
					contadorBinario = 0
				}
				contadorBinario++
				if uno {
					matriz.datos[indiceColumna][indiceFila] = true
				}
				r++
			}
		}
	}
	return matriz
}

//Desproteger le doy un url de entrada y uno de salida
func Desproteger(url string, salida string) error {
	file, err := os.Open(url)
	manejoError(err)
	defer file.Close()
	fileO, err := os.Create(salida)
	manejoError(err)
	fileO.Close()
	fileO, err = os.OpenFile(salida, os.O_WRONLY, 0666)
	manejoError(err)
	defer fileO.Close()

	bufferReader := bufio.NewReader(file)
	bufferWriter := bufio.NewWriter(fileO)

	codificacion, bloqueCodificados, bitsUltimo := obtenerInformacion(bufferReader)

	tamBuf := (codificacion)/8 + 2
	buf := make([]byte, tamBuf)

	for bloqueCodificados != 0 {
		bloqueCodificados--
		byteLeidos, err := bufferReader.Read(buf)
		manejoError(err)
		if byteLeidos < len(buf) {
			buf2 := make([]byte, tamBuf-byteLeidos)
			_, err := bufferReader.Read(buf2)
			manejoError(err)
			buf = buf[:byteLeidos]
			buf = append(buf, buf2...)
		}
		decodificadoBytes := decodificar(buf)
		if bloqueCodificados == 0 {
			decodificadoBytes = decodificadoBytes[:bitsUltimo]
		}
		_, err = bufferWriter.Write(decodificadoBytes)
		manejoError(err)
		if bufferWriter.Available() < len(decodificadoBytes) {
			bufferWriter.Flush()
		}
	}
	bufferWriter.Flush()
	return nil
}

func decodificar(buf []byte) []byte {
	tamBuf := len(buf)
	entradaBool := (ByteToBool(buf))
	matrizEntrada := MatrizColumna(entradaBool)
	r := matrizDecodificadora(tamBuf * 8)
	decodificado, err := r.MultiplicarOpt(matrizEntrada)
	manejoError(err)
	decodificadoBytes := decodificado.ToByte()
	decodificadoBytes = decodificadoBytes[:tamBuf-2]

	return decodificadoBytes
}

//matrizDecodificadora Funcion que crea la matriz decodificadora
func matrizDecodificadora(codificacion int) Matriz {
	anchoMatriz := bitsInformacion(codificacion)
	altoMatriz := codificacion
	matriz := NuevaMatriz(anchoMatriz, altoMatriz)
	indiceColumna := 0
	for indiceFila := 0; indiceFila < altoMatriz; indiceFila++ {
		if !esPotenciaDeDos(indiceFila + 1) {
			matriz.datos[indiceColumna][indiceFila] = true
			indiceColumna++
		}
	}
	return matriz
}

//TieneErrores toma como parametros un archivo .ham con su archivo de informacion y verifica si tiene error
func TieneErrores(url string) (bool, int, int) {
	if existeArchivo(url) {
		file, err := os.Open(url)
		manejoError(err)
		defer file.Close()
		bufferReader := bufio.NewReader(file)
		codificacion, bloqueCodificados, bitsUltimo := obtenerInformacion(bufferReader)

		buf := make([]byte, (codificacion)/8+2)
		tamEscrito := codificacion + bitsParidad(codificacion)
		tamUltimo := bitsUltimo*8 + bitsParidad(bitsUltimo*8)
		tam := tamEscrito

		contadorBloques := 0
		for contadorBloques != bloqueCodificados {

			byteLeidos, err := bufferReader.Read(buf)
			manejoError(err)
			if byteLeidos < len(buf) {
				buf2 := make([]byte, (codificacion)/8+2-byteLeidos)
				byteLeidos2, err := bufferReader.Read(buf2)
				if byteLeidos2 != 0 {
					manejoError(err)
				}
				buf = buf[:byteLeidos]
				buf = append(buf, buf2...)
			}
			contadorBloques++
			if bloqueCodificados == contadorBloques {
				tam = tamUltimo
			}
			if contieneError, posicionError := errorEnBloque(buf, tam); contieneError {
				return true, contadorBloques - 1, posicionError
			}
		}
	} else {
		fmt.Println("No existe el archivo ", url)
	}
	return false, -1, -1
}

func errorEnBloque(buf []byte, tam int) (bool, int) {
	matrizChequeo := matrizChequeoParidad(tam)

	entradaBool := (ByteToBool(buf))
	entradaBool = entradaBool[:tam]

	matrizEntrada := MatrizColumna(entradaBool)
	sindrome, err := matrizChequeo.MultiplicarOpt(matrizEntrada)
	manejoError(err)

	if sindrome.TieneUnos() {
		posicionError := errorEnSindrome(sindrome.datos)
		if posicionError < tam {
			return true, posicionError
		}
	}
	return false, -1
}

func errorEnSindrome(sindrome [][]bool) int {
	posicionError := 0
	mascara := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048}

	for indice, fila := range sindrome {
		if fila[0] {
			posicionError = posicionError | mascara[indice]
		}
	}
	posicionError = posicionError - 1
	return posicionError
}

//matrizChequeoParidad metodo que devuelve la matriz que se multiplica para codificar una entrada, matriz de chequeo
func matrizChequeoParidad(codificacion int) Matriz {
	altoMatriz := codificacion
	anchoMatriz := bitsParidad(codificacion)
	matriz := NuevaMatriz(anchoMatriz, altoMatriz)
	for i := 0; i < anchoMatriz; i++ {
		contadorBinario := 1
		uno := false
		f := float64(i)
		potencia := int(math.Pow(2, f))
		for j := 0; j < altoMatriz; j++ {
			if potencia == contadorBinario {
				contadorBinario = 0
				uno = !uno
			}
			contadorBinario++
			if uno {
				matriz.datos[i][j] = true
			}
		}
	}
	return matriz
}

//IntroducirError toma como parametros un archivo .ham y devuelve un .ham con un erro introducido
func IntroducirError(url string, salida string) {
	error, b, l := TieneErrores(url)
	if error {
		fmt.Println("ERROR: El archivo ya contiene un error en el bloque", b, " en ", l)
	} else {
		file, err := os.Open(url)
		manejoError(err)
		defer file.Close()
		fileO, err := os.Create(salida)
		manejoError(err)
		fileO.Close()
		fileO, err = os.OpenFile(salida, os.O_WRONLY, 0666)
		manejoError(err)
		defer fileO.Close()

		bufferReader := bufio.NewReader(file)
		codificacion, bloquesCodificados, bitsUltimo := obtenerInformacion(bufferReader)
		bufferWriter := bufio.NewWriter(fileO)
		_, err = bufferWriter.WriteString(fmt.Sprintf("%v\n%v\n%v\n", codificacion, bloquesCodificados, bitsUltimo))
		manejoError(err)

		tamEscrito := codificacion + bitsParidad(codificacion)
		tamUltimo := bitsUltimo*8 + bitsParidad(bitsUltimo*8)

		//CONTROLAR
		marcador := bitsUltimo*8 + bitsParidad(bitsUltimo*8)
		ultimoBloque := marcador / 8
		if marcador%8 != 0 {
			ultimoBloque++
		}

		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		bloqueError := 0
		if bloquesCodificados != 0 {
			bloqueError = r.Intn(bloquesCodificados)
		}
		buf := make([]byte, (codificacion)/8+2)

		contadorBloques := -1
		for bloquesCodificados != contadorBloques {
			contadorBloques++
			byteLeidos, err := bufferReader.Read(buf)
			manejoError(err)
			if byteLeidos < len(buf) {
				buf2 := make([]byte, (codificacion)/8+2-byteLeidos)
				byteLeidos2, err := bufferReader.Read(buf2)
				if byteLeidos2 != 0 {
					manejoError(err)
				}
				buf = buf[:byteLeidos]
				buf = append(buf, buf2...)
			}
			if contadorBloques == bloqueError {
				maximo := tamEscrito - 1
				if bloquesCodificados-1 == contadorBloques {
					maximo = tamUltimo - 1
				}
				buf = agregarErrorByte(buf, maximo)
			}
			if bloquesCodificados == contadorBloques {
				buf = buf[:ultimoBloque]
			}

			if bufferWriter.Available() < len(buf) {
				bufferWriter.Flush()
			}
			_, err = bufferWriter.Write(buf)
			manejoError(err)
		}
		bufferWriter.Flush()
	}
}
func agregarErrorByte(buf []byte, tam int) []byte {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r.Intn(tam)
	mascara := []int{128, 64, 32, 16, 8, 4, 2, 1}
	buf[i/8] = byte(int(buf[i/8]) ^ mascara[i%8])
	return buf
}

//CorregirError dado un archivo, su archivo de info y un path de salida genera la salida de los mismos
func CorregirError(url string, salida string) {
	if error, bloqueError, posicionError := TieneErrores(url); !error {
		fmt.Println("El archivo no tiene error")
	} else {
		file, err := os.Open(url)
		manejoError(err)
		defer file.Close()
		fileO, err := os.Create(salida)
		manejoError(err)
		fileO.Close()
		fileO, err = os.OpenFile(salida, os.O_WRONLY, 0666)
		manejoError(err)
		defer fileO.Close()

		bufferReader := bufio.NewReader(file)
		codificacion, bloqueCodificados, bitsUltimo := obtenerInformacion(bufferReader)
		bufferWriter := bufio.NewWriter(fileO)

		_, err = bufferWriter.WriteString(fmt.Sprintf("%v\n%v\n%v\n", codificacion, bloqueCodificados, bitsUltimo))
		manejoError(err)

		buf := make([]byte, (codificacion)/8+2)
		contadorBloques := -1
		for bloqueCodificados != 0 {
			byteLeidos, err := bufferReader.Read(buf)
			manejoError(err)
			if byteLeidos < len(buf) {
				buf2 := make([]byte, (codificacion)/8+2-byteLeidos)
				byteLeidos2, err := bufferReader.Read(buf2)
				if byteLeidos2 != 0 {
					manejoError(err)
				}
				buf = buf[:byteLeidos]
				buf = append(buf, buf2...)
			}
			bloqueCodificados--
			contadorBloques++
			if contadorBloques == bloqueError {
				mascara := []int{128, 64, 32, 16, 8, 4, 2, 1}
				buf[posicionError/8] = byte(int(buf[posicionError/8]) ^ mascara[posicionError%8])
			}
			if bufferWriter.Available() < len(buf) {
				bufferWriter.Flush()
			}
			_, err = bufferWriter.Write(buf)
			manejoError(err)

		}
		bufferWriter.Flush()
	}
}

func escribirArchivo(file io.Writer, fileTemp io.Reader, codificacion, contadorBloques, byteLeidos int) error {
	bufferReader := bufio.NewReader(fileTemp)
	bufferWriter := bufio.NewWriter(file)
	bufR := make([]byte, bufferWriter.Available())
	_, err := bufferWriter.WriteString(fmt.Sprintf("%v\n%v\n%v\n", codificacion, contadorBloques, byteLeidos))
	if err != nil {
		return fmt.Errorf("No se pudo ingresar la informacion %s,%v", file, err)
	}
	err = bufferWriter.Flush()
	if err != nil {
		return fmt.Errorf("No se escribir archivo %v", err)
	}
	_, err = bufferReader.Read(bufR)
	for err != io.EOF {
		bufferWriter.Write(bufR)
		err = bufferWriter.Flush()
		if err != nil {
			return fmt.Errorf("No se escribir archivo %v", err)
		}
		_, err = bufferReader.Read(bufR)

	}
	return nil
}

func obtenerInformacion(buffer *bufio.Reader) (int, int, int) {
	line, err := buffer.ReadString('\n')
	manejoError(err)
	codificacion, err := strconv.Atoi(line[:len(line)-1])
	manejoError(err)
	line, err = buffer.ReadString('\n')
	manejoError(err)
	bloqueCodificados, err := strconv.Atoi(line[:len(line)-1])
	manejoError(err)
	line, err = buffer.ReadString('\n')
	manejoError(err)
	bitsUltimo, err := strconv.Atoi(line[:len(line)-1])
	manejoError(err)

	return codificacion, bloqueCodificados, bitsUltimo
}

func manejoError(err error) {
	if err != nil {
		panic(fmt.Sprintf("hubo un error, %v", err))
	}
}

func existeArchivo(pathArchivo string) bool {
	if _, err := os.Stat(pathArchivo); err == nil {
		return true
	}
	return false
}
