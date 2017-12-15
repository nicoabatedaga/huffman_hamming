package hamming

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"syscall"
	"time"
)

//Hamming es el programa de prueba
func Hamming() {
	start := time.Now()
	fmt.Println("-Hamming-")

	codificacion := 522
	//codificacion = 1035
	//codificacion := 2060

	fmt.Println(fmt.Sprintf("Codificacion: %v, Bits Paridad: %v, Bits Información: %v",
		codificacion,
		bitsParidad(codificacion),
		bitsInformacion(codificacion)))

	Proteger("./prueba.txt", "./prueba.ham", codificacion)
	error, b, l := TieneErrores("./prueba.ham")
	fmt.Println(error, b, l)
	//IntroducirError("./alicia.ham", "./aliciaConError.ham")
	//error, b, l = TieneErrores("./alicia.ham")
	//fmt.Println(error, b, l)
	Desproteger("./prueba.ham", "./pruebaDesprotegido.txt")

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
func h(codificacion int) *Matriz {
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
	return aux
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

//Proteger funcion que toma de entrada el path de un archivo y lo codifica segun un valor de entrada
func Proteger(url string, salida string, codificacion int) {
	fmt.Println("\nProtección Archivo:")
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

	buf := make([]byte, bitsInformacion(codificacion)/8)
	g := g(codificacion)

	byteLeidos, err := bufferReader.Read(buf)
	if byteLeidos != 0 {
		manejoError(err)
	}
	contadorBloques := 0
	for byteLeidos > 0 {
		auxMatriz := MatrizColumna(ByteToBool(buf))
		b, m := g.Multiplicar(auxMatriz)
		if !b {
			contadorBloques++
			bin := m.ToByte()
			if byteLeidos+1 < codificacion/8 {
				marcador := byteLeidos + bitsParidad(byteLeidos*8)/8
				if bitsParidad(byteLeidos*8)%8 != 0 {
					marcador++
				}
				bin = bin[:marcador]
			}
			if len(bin)+1 > bufferWriter.Available() {
				bufferWriter.Flush()
			}
			numB, err := bufferWriter.Write(bin)
			if numB == 0 {
				fmt.Println(contadorBloques, ":No se escribio nada")
			}
			manejoError(err)
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

	salida = salida + "info"

	ocultarFile(salida, false)
	fileOinfo, err := os.Create(salida)
	manejoError(err)
	fileOinfo.Close()
	fileOinfo, err = os.OpenFile(salida, os.O_WRONLY, 0666)
	manejoError(err)
	fileOinfo.WriteString(fmt.Sprintf("%v\n%v\n%v\n", codificacion, contadorBloques, byteLeidos))
	fileOinfo.Close()
	ocultarFile(salida, true)
}

//Desproteger le doy un url de entrada y uno de salida
func Desproteger(url string, salida string) {

	file, err := os.Open(url)
	manejoError(err)
	defer file.Close()
	fileinfo, err := os.Open(url + "info")
	manejoError(err)
	defer fileinfo.Close()
	fileO, err := os.Create(salida)
	manejoError(err)
	fileO.Close()
	fileO, err = os.OpenFile(salida, os.O_WRONLY, 0666)
	manejoError(err)
	defer fileO.Close()

	bufferReaderInfo := bufio.NewReader(fileinfo)
	bufferReader := bufio.NewReader(file)
	bufferWriter := bufio.NewWriter(fileO)

	line, err := bufferReaderInfo.ReadString('\n')
	manejoError(err)
	codificacion, err := strconv.Atoi(line[:len(line)-1])
	manejoError(err)
	line, err = bufferReaderInfo.ReadString('\n')
	manejoError(err)
	bloqueCodificados, err := strconv.Atoi(line[:len(line)-1])
	manejoError(err)
	line, err = bufferReaderInfo.ReadString('\n')
	manejoError(err)
	bitsUltimo, err := strconv.Atoi(line[:len(line)-1])
	manejoError(err)

	fmt.Println("\nDesprotección Archivo:", codificacion, bloqueCodificados, bitsUltimo)

	buf := make([]byte, (codificacion)/8+1)
	bitesInfo := bitsInformacion(codificacion)
	r := r(len(buf) * 8)

	byteLeidos, err := bufferReader.Read(buf)
	if byteLeidos != 0 {
		manejoError(err)
	}
	contadorBloques := 0
	for bloqueCodificados != 0 {
		bloqueCodificados--
		auxBool := (ByteToBool(buf))
		//[:bitesInfo]
		auxMatriz := MatrizColumna(auxBool)
		b, m := r.Multiplicar(auxMatriz)
		if !b {
			contadorBloques++
			bin := m.ToByte()
			if bloqueCodificados == 0 {
				bin = bin[:bitsUltimo]
			} else {
				bin = bin[:bitesInfo/8]
			}
			numB, err := bufferWriter.Write(bin)

			if bufferWriter.Available() < len(bin) {
				bufferWriter.Flush()
			}
			if numB == 0 {
				fmt.Println(contadorBloques, ":No se escribio nada")
			}
			manejoError(err)
		}
		byteLeidos, err = bufferReader.Read(buf)
		if byteLeidos != 0 {
			manejoError(err)
		}
		if byteLeidos < len(buf) {
			//Se agrego este segmento, para evitar una lectura menor que el buff., se continua leyendo hasta completarlo
			buf2 := make([]byte, (codificacion)/8+1-byteLeidos)
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

//IntroducirError toma como parametros un archivo .ham y devuelve un .ham con un erro introducido
func IntroducirError(url string, salida string) {
	error, b, l := TieneErrores(url)
	if error {
		fmt.Println("El archivo ya contiene un error en el bloque", b, " en ", l)
	} else {
		file, err := os.Open(url)
		manejoError(err)
		defer file.Close()
		fileinfo, err := os.Open(url + "info")
		manejoError(err)
		defer fileinfo.Close()
		fileO, err := os.Create(salida)
		manejoError(err)
		fileO.Close()
		fileO, err = os.OpenFile(salida, os.O_WRONLY, 0666)
		manejoError(err)
		defer fileO.Close()

		bufferReaderInfo := bufio.NewReader(fileinfo)
		bufferReader := bufio.NewReader(file)
		bufferWriter := bufio.NewWriter(fileO)

		line, err := bufferReaderInfo.ReadString('\n')
		manejoError(err)
		codificacion, err := strconv.Atoi(line[:len(line)-1])
		manejoError(err)
		line, err = bufferReaderInfo.ReadString('\n')
		manejoError(err)
		bloqueCodificados, err := strconv.Atoi(line[:len(line)-1])
		manejoError(err)
		line, err = bufferReaderInfo.ReadString('\n')
		manejoError(err)
		bitsUltimo, err := strconv.Atoi(line[:len(line)-1])
		manejoError(err)

		fmt.Println("\nIntroducción error:")
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		bloqueError := r.Intn(bloqueCodificados)
		fmt.Println("Bloque Error:", bloqueError)

		buf := make([]byte, (codificacion)/8+1)
		bitesInfo := bitsInformacion(codificacion)

		byteLeidos, err := bufferReader.Read(buf)
		if byteLeidos != 0 {
			manejoError(err)
		}
		contadorBloques := 0
		for bloqueCodificados != 0 {
			bloqueCodificados--
			auxBool := (ByteToBool(buf))
			//[:bitesInfo]
			auxMatriz := MatrizColumna(auxBool)
			if contadorBloques == bloqueError {
				n := len(auxMatriz.datos)
				m := len(auxMatriz.datos[0])
				i := r.Intn(n)
				j := r.Intn(m)
				fmt.Println("Posicion error:", i)
				auxMatriz.datos[i][j] = !auxMatriz.datos[i][j]
			}
			contadorBloques++
			bin := auxMatriz.ToByte()
			if bloqueCodificados == 0 {
				bin = bin[:bitsUltimo]
			} else {
				bin = bin[:bitesInfo/8]
			}
			numB, err := bufferWriter.Write(bin)

			if bufferWriter.Available() < len(bin) {
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
				buf2 := make([]byte, (codificacion)/8+1-byteLeidos)
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

//TieneErrores toma como parametros un archivo .ham  y verifica si tiene error
func TieneErrores(url string) (bool, int, int) {
	file, err := os.Open(url)
	manejoError(err)
	defer file.Close()
	fileinfo, err := os.Open(url + "info")
	manejoError(err)
	defer fileinfo.Close()

	bufferReaderInfo := bufio.NewReader(fileinfo)
	bufferReader := bufio.NewReader(file)

	line, err := bufferReaderInfo.ReadString('\n')
	manejoError(err)
	codificacion, err := strconv.Atoi(line[:len(line)-1])
	manejoError(err)
	line, err = bufferReaderInfo.ReadString('\n')
	manejoError(err)
	bloqueCodificados, err := strconv.Atoi(line[:len(line)-1])
	manejoError(err)
	line, err = bufferReaderInfo.ReadString('\n')
	manejoError(err)
	bitsUltimo, err := strconv.Atoi(line[:len(line)-1])
	manejoError(err)

	fmt.Println("\nControlo error:")

	buf := make([]byte, (codificacion)/8+1)
	hM := h(len(buf) * 8)

	byteLeidos, err := bufferReader.Read(buf)
	if byteLeidos != 0 {
		manejoError(err)
	}
	contadorBloques := 0
	for bloqueCodificados != 0 {
		bloqueCodificados--
		auxBool := (ByteToBool(buf))
		if bloqueCodificados == 0 {
			marcardor := bitsUltimo * 8
			auxBool = auxBool[:marcardor]
			hM = h(marcardor)
			fmt.Println("MARCADOR: ", marcardor, bitsParidad(marcardor), len(auxBool))
		}
		auxMatriz := MatrizColumna(auxBool)
		b, sindrome := hM.Multiplicar(auxMatriz)
		if !b {
			contadorBloques++
			if sindrome.TieneUnos() {
				fmt.Println(sindrome.ToString())
				auxInt := 0
				for i, fila := range sindrome.datos {
					mascara := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048}
					for _, b := range fila {
						if b {
							auxInt = auxInt | mascara[i]
						}
					}
				}
				auxInt = auxInt - 1
				fmt.Println("Hay error en el bloque: ", contadorBloques, " en el bit ", auxInt)
				return true, contadorBloques, auxInt
			}
		}
		byteLeidos, err = bufferReader.Read(buf)
		if byteLeidos != 0 {
			manejoError(err)
		}
		if byteLeidos < len(buf) {
			//Se agrego este segmento, para evitar una lectura menor que el buff., se continua leyendo hasta completarlo
			buf2 := make([]byte, (codificacion)/8+1-byteLeidos)
			byteLeidos2, err := bufferReader.Read(buf2)
			if byteLeidos2 != 0 {
				manejoError(err)
			}
			buf = buf[:byteLeidos]
			buf = append(buf, buf2...)
		}
	}
	return false, -1, -1
}

func ocultarFile(url string, ocultar bool) {
	if _, err := os.Stat(url); err == nil {
		nameptr, err := syscall.UTF16PtrFromString(url)
		manejoError(err)
		if ocultar {
			err = syscall.SetFileAttributes(nameptr, syscall.FILE_ATTRIBUTE_HIDDEN)
		} else {
			err = syscall.SetFileAttributes(nameptr, syscall.FILE_ATTRIBUTE_NORMAL)
		}
		manejoError(err)
	}
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
