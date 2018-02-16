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
	512:  10,
	1024: 11,
	2048: 12,
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
		fileOut, err := os.OpenFile(salida, os.O_WRONLY, 0666)
		manejoError(err)
		defer fileOut.Close()
		defer os.Remove(fileTemp.Name())

		bufferReader := bufio.NewReader(file)
		bufferWriter := bufio.NewWriter(fileTemp)
		buf := make([]byte, codificacion/8)
		matrizG := g(codificacion + bitsParidad(codificacion))

		byteLeidos, err := bufferReader.Read(buf)
		if byteLeidos != 0 {
			manejoError(err)
		}
		numB := 0
		contadorBloques := 0
		for err != io.EOF {
			matrizEntrada := ByteToMatriz(buf)
			codificado, err := matrizG.MultiplicarOpt(matrizEntrada)
			manejoError(err)

			contadorBloques++
			codificadoByte := codificado.ToByte()
			if byteLeidos+1 < codificacion/8 {
				marcador := byteLeidos + bitsParidad(byteLeidos*8)/8
				if bitsParidad(byteLeidos*8)%8 != 0 {
					marcador++
				}
				codificadoByte = codificadoByte[:marcador]
			}
			if len(codificadoByte)+1 > bufferWriter.Available() {
				bufferWriter.Flush()
			}
			numB, err = bufferWriter.Write(codificadoByte)
			manejoError(err)

			if numB == 0 {
				fmt.Println(contadorBloques, ":No se escribio nada")
			}
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

//g Funcion que crea la matriz generadora
func g(codificacion int) Matriz {
	anchoMatriz := codificacion
	altoMatriz := bitsInformacion(codificacion)
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
	r := r(tamBuf * 8)

	byteLeidos, err := bufferReader.Read(buf)
	if byteLeidos != 0 {
		manejoError(err)
	}
	contadorBloques := 0
	for bloqueCodificados != 0 {
		bloqueCodificados--
		entradaBool := (ByteToBool(buf))
		matrizEntrada := MatrizColumna(entradaBool)
		decodificado, err := r.MultiplicarOpt(matrizEntrada)
		manejoError(err)

		contadorBloques++
		decodificadoBytes := decodificado.ToByte()
		if bloqueCodificados == 0 {
			decodificadoBytes = decodificadoBytes[:bitsUltimo]
		} else {
			decodificadoBytes = decodificadoBytes[:codificacion/8]
		}
		numB, err := bufferWriter.Write(decodificadoBytes)

		if bufferWriter.Available() < len(decodificadoBytes) {
			bufferWriter.Flush()
		}
		if numB == 0 {
			fmt.Println(contadorBloques, ":No se escribio nada")
		}
		manejoError(err)
		byteLeidos, err = bufferReader.Read(buf)
		if byteLeidos != 0 {
			manejoError(err)
		}
		if byteLeidos < len(buf) {
			buf2 := make([]byte, tamBuf-byteLeidos)
			byteLeidos2, err := bufferReader.Read(buf2)
			if byteLeidos2 != 0 {
				manejoError(err)
			}
			buf = buf[:byteLeidos]
			buf = append(buf, buf2...)

		}
	}
	bufferWriter.Flush()
	return nil
}

//r Funcion que crea la matriz decodificadora
func r(codificacion int) Matriz {
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
		matrizChequeo := h(tamEscrito)

		byteLeidos, err := bufferReader.Read(buf)
		if byteLeidos != 0 {
			manejoError(err)
		}

		contadorBloques := 0
		for bloqueCodificados != 0 {
			bloqueCodificados--
			entradaBool := (ByteToBool(buf))
			if bloqueCodificados == 0 {
				entradaBool = entradaBool[:tamUltimo]
				matrizChequeo = h(tamUltimo)
			} else {
				entradaBool = entradaBool[:tamEscrito]
			}
			auxMatriz := MatrizColumna(entradaBool)
			sindrome, err := matrizChequeo.MultiplicarOpt(auxMatriz)
			if err != nil {
				manejoError(fmt.Errorf("tieneErrror: error al multiplicar matrices, %s", err))
			}
			if sindrome.TieneUnos() {

				auxInt := 0
				for i, fila := range sindrome.datos {
					mascara := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1028, 2048}
					if fila[0] {
						auxInt = auxInt | mascara[i]
					}
				}
				auxInt = auxInt - 1
				resultado2 := errorEnSindrome(sindrome.datos)
				fmt.Println("Menor-Mayor matriz: ", auxInt, " Mayor-Menor matriz: ", resultado2)
				if bloqueCodificados == 0 {
					fmt.Println("Tamaño ultimo", tamUltimo)
					if auxInt < tamUltimo {
						return true, contadorBloques, auxInt
					}
				} else {
					fmt.Println("Tamaño escrito", tamEscrito)

					return true, contadorBloques, auxInt
				}
			}
			byteLeidos, err = bufferReader.Read(buf)
			if byteLeidos != 0 {
				manejoError(err)
			}
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
		}
	} else {
		fmt.Println("No existe el archivo ", url)
	}
	return false, -1, -1
}

func errorEnSindrome(sindrome [][]bool) int {
	auxInt := 0
	mascara := []int{2048, 1028, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}
	mascara = mascara[12-len(sindrome):]
	for i, b := range sindrome {
		if b[0] {
			auxInt = auxInt | mascara[i]
		}
	}
	auxInt = auxInt - 1
	return auxInt
}

//h metodo que devuelve la matriz que se multiplica para codificar una entrada, matriz de chequeo
func h(codificacion int) Matriz {
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
		if bloquesCodificados == 0 && bitsUltimo == 0 {
			fmt.Println("Es un archivo vacio")
		}
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		bloqueError := 0
		if bloquesCodificados != 0 {
			bloqueError = r.Intn(bloquesCodificados)
		}
		buf := make([]byte, (codificacion)/8+2)
		byteLeidos, err := bufferReader.Read(buf)
		if byteLeidos != 0 {
			manejoError(err)
		}
		contadorBloques := -1
		for bloquesCodificados != contadorBloques {
			contadorBloques++
			if contadorBloques == bloqueError {
				maximo := byteLeidos*8 - 1 + bitsParidad(byteLeidos*8)
				if bloquesCodificados == contadorBloques {
					maximo = bitsUltimo*8 - 1 + bitsParidad(bitsUltimo*8)
				}
				i := r.Intn(maximo)
				mascara := []int{128, 64, 32, 16, 8, 4, 2, 1}
				fmt.Println("Bytes Leidos: ", byteLeidos, " len buf: ", len(buf), " i: ", i, "posicion mascara ", i%8, " mascara ", mascara[i%8], "valor: ", int(buf[i/8]))
				fmt.Printf("Antes %8b ", int(buf[i/8]))
				buf[i/8] = byte(int(buf[i/8]) ^ mascara[i%8])
				fmt.Printf(" Despues %8b ", int(buf[i/8]))

			}
			if bloquesCodificados == contadorBloques {
				buf = buf[:bitsUltimo]
			}

			if bufferWriter.Available() < len(buf) {
				bufferWriter.Flush()
			}
			numB, err := bufferWriter.Write(buf)

			if numB == 0 {
				fmt.Println(":No se escribio nada")
			}
			manejoError(err)
			byteLeidos, err = bufferReader.Read(buf)
			if byteLeidos != 0 {
				manejoError(err)
			}
			if byteLeidos < len(buf) {
				buf2 := make([]byte, (codificacion)/8+2-byteLeidos)
				byteLeidos2, err := bufferReader.Read(buf2)
				if byteLeidos2 != 0 {
					manejoError(err)
				}
				buf = buf[:byteLeidos]
				buf = append(buf, buf2...)
			}
		}
		bufferWriter.Flush()
	}
}

//IntroducirErrorNR toma como parametros un archivo .ham y devuelve un .ham con un erro introducido en una posicion determinada
func IntroducirErrorNR(url string, salida string, bloque, posicion int) {
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
		if bloquesCodificados == 0 && bitsUltimo == 0 {
			fmt.Println("Es un archivo vacio")
		}
		bloqueError := bloque
		buf := make([]byte, (codificacion)/8+2)
		byteLeidos, err := bufferReader.Read(buf)
		if byteLeidos != 0 {
			manejoError(err)
		}
		contadorBloques := -1
		for bloquesCodificados != contadorBloques {
			contadorBloques++
			if contadorBloques == bloqueError {
				i := posicion
				mascara := []int{128, 64, 32, 16, 8, 4, 2, 1}
				fmt.Println("Bytes Leidos: ", byteLeidos, " len buf: ", len(buf), " i: ", i, "posicion mascara ", i%8, " mascara ", mascara[i%8], "valor: ", int(buf[i/8]))
				fmt.Printf("Antes %8b", int(buf[i/8]))
				buf[i/8] = byte(int(buf[i/8]) ^ mascara[i%8])
				fmt.Printf("Despues %8b", int(buf[i/8]))

			}
			if bloquesCodificados == contadorBloques {
				buf = buf[:bitsUltimo]
			}

			if bufferWriter.Available() < len(buf) {
				bufferWriter.Flush()
			}
			numB, err := bufferWriter.Write(buf)

			if numB == 0 {
				fmt.Println(":No se escribio nada")
			}
			manejoError(err)
			byteLeidos, err = bufferReader.Read(buf)
			if byteLeidos != 0 {
				manejoError(err)
			}
			if byteLeidos < len(buf) {
				buf2 := make([]byte, (codificacion)/8+2-byteLeidos)
				byteLeidos2, err := bufferReader.Read(buf2)
				if byteLeidos2 != 0 {
					manejoError(err)
				}
				buf = buf[:byteLeidos]
				buf = append(buf, buf2...)
			}
		}
		bufferWriter.Flush()
	}
}

//CorregirError dado un archivo, su archivo de info y un path de salida genera la salida de los mismos
func CorregirError(url string, salida string) {
	fmt.Println("Corregir Error")
	error, bloqueError, posicionError := TieneErrores(url)
	if !error {
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

		fmt.Println("Bloque Error:", bloqueError)

		buf := make([]byte, (codificacion)/8+2)
		byteLeidos, err := bufferReader.Read(buf)
		if byteLeidos != 0 {
			manejoError(err)
		}
		contadorBloques := -1
		for bloqueCodificados != 0 {
			bloqueCodificados--
			contadorBloques++
			if contadorBloques == bloqueError {
				fmt.Println("Posición: ", posicionError)
				mascara := []int{1, 128, 64, 32, 16, 8, 4, 2}
				fmt.Println("Antes: ", buf[posicionError/8])
				fmt.Println("Mascara: ", mascara[posicionError%8])
				buf[posicionError/8] = byte(int(buf[posicionError/8]) ^ mascara[posicionError%8])
				fmt.Println("Despues: ", buf[posicionError/8])

			}
			if bloqueCodificados == 0 {
				buf = buf[:bitsUltimo]
			} else {
				buf = buf[:byteLeidos]
			}
			numB, err := bufferWriter.Write(buf)

			if bufferWriter.Available() < len(buf) {
				bufferWriter.Flush()
			}
			if numB == 0 {
				fmt.Println(contadorBloques, ":No se escribio nada")
			}
			manejoError(err)
			byteLeidos, err = bufferReader.Read(buf)
			if byteLeidos != 0 {
				manejoError(err)
			}
			if byteLeidos < len(buf) {
				//Se agrego este segmento, para evitar una lectura menor que el buff., se continua leyendo hasta completarlo
				buf2 := make([]byte, (codificacion)/8+2-byteLeidos)
				byteLeidos2, err := bufferReader.Read(buf2)
				if byteLeidos2 != 0 {
					manejoError(err)
				}
				buf = buf[:byteLeidos]
				buf = append(buf, buf2...)
			}
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
