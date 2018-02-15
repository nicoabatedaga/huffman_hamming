package hamming

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var informacion = map[int]int{
	522:  512,
	1035: 1024,
	1040: 1024,
	2060: 2048,
}
var paridad = map[int]int{
	522:  10,
	1035: 11,
	1040: 11,
	2060: 12,
}

//GenerarArchivosMatrices genera los distintos archivos para las matrices
func GenerarArchivosMatrices(conInfo bool) {
	codificaciones := []int{522, 1035, 2060}
	for _, cod := range codificaciones {
		h := h(cod)
		h.ToFile(fmt.Sprintf("./H%v.matriz", cod))
		g := g(cod)
		g.ToFile(fmt.Sprintf("./G%v.matriz", cod))
		r := r(cod)
		r.ToFile(fmt.Sprintf("./R%v.matriz", cod))
		if conInfo {
			h.ToFileConInfo(fmt.Sprintf("./H%vconInf.matriz", cod))
			g.ToFileConInfo(fmt.Sprintf("./G%vconInf.matriz", cod))
			r.ToFileConInfo(fmt.Sprintf("./R%vconInfo.matriz", cod))
		}

	}
}

//Hamming es el programa de prueba
func Hamming() {
	//codificacion := 522
	//codificacion = 1035
	//codificacion := 2060
	comienzoPrograma := time.Now()
	/*f, err := os.Create("./cpuSinPunt.trace")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	trace.Start(f)
	defer trace.Stop()*/

	/*f, err := os.Create("./cpu.profile")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()*/
	//testProteccionDesproteccionArchivoByte()
	testProteccionDesproteccionArchivo()
	fmt.Println("Tiempo ejecucion:\t", tiempoStr(time.Now().Sub(comienzoPrograma)))

	/*f, err = os.Create("./memoria.profile")
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	runtime.GC() // get up-to-date statistics
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
	f.Close()*/
	return
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
	for i := 1; i <= ent; i = i * 2 {
		n++
	}
	if ent == 7 {
		return 3
	}
	if ent == 31 {
		return 5
	}
	paridad[ent] = n
	return n
}

func matrizChequeoDeParidadH(codificacion int) *MatrizBytes {
	c := [][]byte{{}}
	var matrizVacia = MatrizBytes{datos: c}
	if codificacion == 522 || codificacion == 1040 || codificacion == 2060 {
		faltante := 7 - codificacion%8
		alto := codificacion + faltante
		ancho := 2
		matriz := NuevaMatrizBytes(ancho, alto)
		for posicion := 0; posicion < paridad[codificacion]; posicion++ {
			contadorBinario := 1
			uno := false
			f := float64(posicion)
			potencia := int(math.Pow(2, f))
			for fila := 0; fila < alto; fila++ {
				if potencia == contadorBinario {
					contadorBinario = 0
					uno = !uno
				}
				contadorBinario++
				if uno {
					err := matriz.Set(fila, posicion, true)
					if err != nil {
						fmt.Printf("h: no se pudo settiar %s\n", err)
						return &matrizVacia
					}
				}
			}
		}
		return matriz

	}
	fmt.Printf("h: no corresponde a una codificacion aceptada %v\n", codificacion)
	return &matrizVacia
}

//h metodo que devuelve la matriz que se multiplica para codificar una entrada
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

func matrizGeneradoraG(codificacion int) *MatrizBytes {
	c := [][]byte{{}}
	var matrizVacia = MatrizBytes{datos: c}
	if codificacion == 522 || codificacion == 1035 || codificacion == 2060 {
		ancho := informacion[codificacion] / 8
		alto := codificacion
		matriz := NuevaMatrizBytes(ancho, alto)
		indiceColumnaIdentidad := 0
		p := -1
		for indiceFila := 0; indiceFila < alto; indiceFila++ {
			if !esPotenciaDeDos(indiceFila + 1) {
				matriz.Set(indiceFila, indiceColumnaIdentidad, true)
				indiceColumnaIdentidad++
			} else {
				p++
				f := float64(p)
				potencia := int(math.Pow(2, f))
				r := 1
				uno := false
				contadorBinario := 1
				for indiceColumna := 0; indiceColumna < ancho*8; indiceColumna++ {
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
						matriz.Set(indiceFila, indiceColumna, true)
					}
					r++
				}

			}
		}

		return matriz
	}
	fmt.Printf("g: no corresponde a una codificacion aceptada %v", codificacion)
	return &matrizVacia

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

func matrizDecodificadorR(codificacion int) *MatrizBytes {
	c := [][]byte{{}}
	var matrizVacia = MatrizBytes{datos: c}
	if codificacion == 522 || codificacion == 1035 || codificacion == 2060 {
		faltante := 8 - codificacion%8
		ancho := codificacion + faltante
		alto := informacion[codificacion]
		matriz := NuevaMatrizBytes(ancho/8, alto)
		indiceColumna := 0
		for indiceFila := 0; indiceFila < alto; indiceFila++ {
			for esPotenciaDeDos(indiceColumna + 1) {
				indiceColumna++
			}
			error := matriz.Set(indiceFila, indiceColumna, true)
			manejoError(error)
			indiceColumna++
		}
		return matriz
	}
	fmt.Println("r: no corresponde a una codificacion aceptada ", codificacion)
	return &matrizVacia

}

//ProtegerByte funcion que toma de entrada el path de un archivo y lo codifica segun un valor de entrada
func ProtegerByte(url string, info string, salida string, codificacion int) error {
	if !(codificacion == 2060 || codificacion == 1035 || codificacion == 522) {
		return fmt.Errorf("error:no corresponde a una codificacion aceptada, %v", codificacion)
	}
	if existeArchivo(url) {
		file, err := os.Open(url)
		if err != nil {
			return fmt.Errorf("error:no se pudo abrir el archivo %s %s", url, err)
		}
		defer file.Close()

		fileO, err := os.Create(salida)
		if err != nil {
			return fmt.Errorf("error:no se pudo crear el archivo de salida %s %s", salida, err)
		}
		fileO.Close()
		fileO, err = os.OpenFile(salida, os.O_WRONLY, 0666)
		if err != nil {
			return fmt.Errorf("error:no se pudo abrir el archivo de salida %s %s", salida, err)
		}
		defer fileO.Close()

		bufferReader := bufio.NewReader(file)
		bufferWriter := bufio.NewWriter(fileO)

		buf := make([]byte, informacion[codificacion]/8)
		g := matrizGeneradoraG(codificacion)
		/*fmt.Println("Matriz G:")

		fmt.Println(g.datos)

		fmt.Println("----------")*/

		byteLeidos, err := bufferReader.Read(buf)
		if byteLeidos != 0 {
			manejoError(err)
		}
		contadorBloques := 0
		for byteLeidos > 0 {
			auxMatriz := MatrizColumnaByte(buf)
			m, b := g.Multiplicar(auxMatriz)
			if b == nil {
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
			} else {
				manejoError(b)
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

		fileOinfo, err := os.Create(info)
		if err != nil {
			return fmt.Errorf("error:no se pudo crear el archivo %s %s", info, err)
		}
		fileOinfo.Close()
		fileOinfo, err = os.OpenFile(info, os.O_WRONLY, 0666)
		if err != nil {
			return fmt.Errorf("error:no se pudo abrir el archivo %s %s", info, err)
		}
		fileOinfo.WriteString(fmt.Sprintf("%v\n%v\n%v\n", codificacion, contadorBloques, byteLeidos))
		fileOinfo.Close()
		return nil
	}
	return fmt.Errorf("El archivo no existe, %s", url)
}

//Proteger funcion que toma de entrada el path de un archivo y lo codifica segun un valor de entrada
func Proteger(url string, info string, salida string, codificacion int) error {
	if !(codificacion == 2060 || codificacion == 1035 || codificacion == 522) {
		return fmt.Errorf("error:no corresponde a una codificacion aceptada, %v", codificacion)
	}
	if existeArchivo(url) {
		file, err := os.Open(url)
		if err != nil {
			return fmt.Errorf("error:no se pudo abrir el archivo %s %s", url, err)
		}
		defer file.Close()

		fileO, err := os.Create(salida)
		if err != nil {
			return fmt.Errorf("error:no se pudo crear el archivo de salida %s %s", salida, err)
		}
		fileO.Close()
		fileO, err = os.OpenFile(salida, os.O_WRONLY, 0666)
		if err != nil {
			return fmt.Errorf("error:no se pudo abrir el archivo de salida %s %s", salida, err)
		}
		defer fileO.Close()

		bufferReader := bufio.NewReader(file)
		bufferWriter := bufio.NewWriter(fileO)

		buf := make([]byte, informacion[codificacion]/8)
		matrizG := g(codificacion)

		byteLeidos, err := bufferReader.Read(buf)
		if byteLeidos != 0 {
			manejoError(err)
		}

		contadorBloques := 0
		for byteLeidos > 0 {
			arbool := ByteToBool(buf)
			auxMatriz := MatrizColumna(arbool)
			b, m := matrizG.MultiplicarOpt(auxMatriz)
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

		fileOinfo, err := os.Create(info)
		if err != nil {
			return fmt.Errorf("error:no se pudo crear el archivo %s %s", info, err)
		}
		fileOinfo.Close()
		fileOinfo, err = os.OpenFile(info, os.O_WRONLY, 0666)
		if err != nil {
			return fmt.Errorf("error:no se pudo abrir el archivo %s %s", info, err)
		}
		fileOinfo.WriteString(fmt.Sprintf("%v\n%v\n%v\n", codificacion, contadorBloques, byteLeidos))
		fileOinfo.Close()
		return nil
	}
	return fmt.Errorf("El archivo no existe, %s", url)
}

func readFile(url, info string, tamBuf, codificacion int, canalAProcesar chan datosAProcesar) {
	file, err := os.Open(url)
	manejoError(err)
	defer file.Close()

	bufferReader := bufio.NewReader(file)
	buf := make([]byte, tamBuf)

	byteLeidos, err := bufferReader.Read(buf)
	if byteLeidos != 0 {
		manejoError(err)
	}
	contadorBloques := 0
	for byteLeidos > 0 {
		contadorBloques++
		canalAProcesar <- datosAProcesar{datos: buf, bytesLeidos: byteLeidos}
		if byteLeidos < len(buf) {
			break
		}
		byteLeidos, err = bufferReader.Read(buf)
		if byteLeidos != 0 {
			manejoError(err)
		}
	}
	close(canalAProcesar)
	escribirInfo(info, codificacion, contadorBloques, byteLeidos)

}
func escribirInfo(path string, codificacion, contadorBloques, byteLeidos int) {
	fileOinfo, err := os.Create(path)
	manejoError(err)
	fileOinfo.Close()
	fileOinfo, err = os.OpenFile(path, os.O_WRONLY, 0666)
	manejoError(err)

	fileOinfo.WriteString(fmt.Sprintf("%v\n%v\n%v\n", codificacion, contadorBloques, byteLeidos))
	fileOinfo.Close()
}

type datosAProcesar struct {
	datos       []byte
	bytesLeidos int
}

func procesarCanal(codificacion int, canalAProcesar chan datosAProcesar, canalDatos chan []byte) {
	matrizG := g(codificacion)
	for d := range canalAProcesar {
		arbool := ByteToBool(d.datos)
		auxMatriz := MatrizColumna(arbool)
		b, m := matrizG.MultiplicarOpt(auxMatriz)
		if !b {
			bin := m.ToByte()
			if d.bytesLeidos+1 < codificacion/8 {
				marcador := d.bytesLeidos + bitsParidad(d.bytesLeidos*8)/8
				if bitsParidad(d.bytesLeidos*8)%8 != 0 {
					marcador++
				}
				bin = bin[:marcador]
			}
			canalDatos <- bin
		}

	}
	close(canalDatos)
}

func writeFile(salida string, canalDatos chan []byte, semaforo chan bool) {
	fileO, err := os.Create(salida)
	manejoError(err)
	fileO.Close()
	fileO, err = os.OpenFile(salida, os.O_WRONLY, 0666)
	manejoError(err)
	defer fileO.Close()
	bufferWriter := bufio.NewWriter(fileO)

	for bin := range canalDatos {
		if len(bin)+1 > bufferWriter.Available() {
			bufferWriter.Flush()
		}
		numB, err := bufferWriter.Write(bin)
		if numB == 0 {
			fmt.Println("No se escribio nada")
		}
		manejoError(err)
	}
	bufferWriter.Flush()
	semaforo <- true
}

//ProtegerChan funcion que toma de entrada el path de un archivo y lo codifica segun un valor de entrada utilizando canales
func ProtegerChan(url string, info string, salida string, codificacion int) error {
	if !(codificacion == 2060 || codificacion == 1035 || codificacion == 522) {
		return fmt.Errorf("error:no corresponde a una codificacion aceptada, %v", codificacion)
	}
	if existeArchivo(url) {
		canalDatos := make(chan []byte)
		canalAProcesar := make(chan datosAProcesar)
		semaforo := make(chan bool)
		go readFile(url, info, informacion[codificacion]/8, codificacion, canalAProcesar)
		go procesarCanal(codificacion, canalAProcesar, canalDatos)
		go writeFile(salida, canalDatos, semaforo)
		c := <-semaforo
		fmt.Println(c)
		return nil

	} else {
		return fmt.Errorf("El archivo no existe, %s", url)

	}
}

//Desproteger le doy un url de entrada y uno de salida
func Desproteger(url string, info string, salida string) error {

	file, err := os.Open(url)
	if err != nil {
		return fmt.Errorf("error:no se pudo abrir el archivo %s %s", url, err)
	}
	defer file.Close()
	fileO, err := os.Create(salida)
	if err != nil {
		return fmt.Errorf("error:no se pudo crear el archivo %s %s", salida, err)
	}
	fileO.Close()
	fileO, err = os.OpenFile(salida, os.O_WRONLY, 0666)
	if err != nil {
		return fmt.Errorf("error:no se pudo abrir el archivo %s %s", salida, err)
	}
	defer fileO.Close()

	bufferReader := bufio.NewReader(file)
	bufferWriter := bufio.NewWriter(fileO)

	codificacion, bloqueCodificados, bitsUltimo := obtenerInformacion(info)

	buf := make([]byte, (codificacion)/8+1)
	bitesInfo := informacion[codificacion]
	r := r(len(buf) * 8)

	byteLeidos, err := bufferReader.Read(buf)
	if byteLeidos != 0 {
		manejoError(err)
	}
	contadorBloques := 0
	for bloqueCodificados != 0 {
		bloqueCodificados--
		auxBool := (ByteToBool(buf))
		auxMatriz := MatrizColumna(auxBool)
		b, m := r.MultiplicarOpt(auxMatriz)
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
	return nil
}

//DesprotegerByte le doy un url de entrada y uno de salida
func DesprotegerByte(url string, info string, salida string) error {

	file, err := os.Open(url)
	if err != nil {
		return fmt.Errorf("error:no se pudo abrir el archivo %s %s", url, err)
	}
	defer file.Close()
	fileO, err := os.Create(salida)
	if err != nil {
		return fmt.Errorf("error:no se pudo crear el archivo %s %s", salida, err)
	}
	fileO.Close()
	fileO, err = os.OpenFile(salida, os.O_WRONLY, 0666)
	if err != nil {
		return fmt.Errorf("error:no se pudo abrir el archivo %s %s", salida, err)
	}
	defer fileO.Close()

	bufferReader := bufio.NewReader(file)
	bufferWriter := bufio.NewWriter(fileO)

	codificacion, bloqueCodificados, bitsUltimo := obtenerInformacion(info)

	buf := make([]byte, (codificacion)/8+1)
	bitesInfo := informacion[codificacion]
	r := matrizDecodificadorR(codificacion)
	byteLeidos, err := bufferReader.Read(buf)
	if byteLeidos != 0 {
		manejoError(err)
	}
	contadorBloques := 0
	for bloqueCodificados != 0 {
		bloqueCodificados--
		auxMatriz := MatrizColumnaByte(buf)
		m, error := r.Multiplicar(auxMatriz)
		if error == nil {

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
		} else {
			manejoError(error)
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
	return nil
}

//IntroducirError toma como parametros un archivo .ham y devuelve un .ham con un erro introducido
func IntroducirError(url string, info string, salida string) {
	error, b, l := TieneErrores(url, info)
	if error {
		fmt.Println("El archivo ya contiene un error en el bloque", b, " en ", l)
	} else {
		codificacion, bloqueCodificados, bitsUltimo := obtenerInformacion(info)
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

		fmt.Println("\nIntroducción error:")
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		bloqueError := r.Intn(bloqueCodificados)
		fmt.Println("Bloque Error:", bloqueError)

		buf := make([]byte, (codificacion)/8+1)
		byteLeidos, err := bufferReader.Read(buf)
		if byteLeidos != 0 {
			manejoError(err)
		}
		contadorBloques := -1
		for bloqueCodificados != 0 {
			bloqueCodificados--
			contadorBloques++
			if contadorBloques == bloqueError {
				maximo := len(buf)*8 - 1
				if bloqueCodificados == 0 {
					maximo = bitsUltimo - 1
				}
				i := r.Intn(maximo)
				fmt.Println("Posición: ", i)
				mascara := []int{1, 128, 64, 32, 16, 8, 4, 2}
				fmt.Println("Antes: ", buf[i/8])
				fmt.Println("Mascara: ", mascara[i%8])
				buf[i/8] = byte(int(buf[i/8]) ^ mascara[i%8])
				fmt.Println("Despues: ", buf[i/8])

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

/*
//CorregirError dado un archivo, su archivo de info y un path de salida genera la salida de los mismos
func CorregirError(url string, info string, salida string) {
	fmt.Println("Corregir Error")
	error, bloqueError, posicionError := TieneErrores(url, info)
	if !error {
		fmt.Println("El archivo no tiene error")
	} else {
		codificacion, bloqueCodificados, bitsUltimo := obtenerInformacion(info)
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

		fmt.Println("Bloque Error:", bloqueError)

		buf := make([]byte, (codificacion)/8+1)
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
}*/
func obtenerInformacion(info string) (int, int, int) {
	fileinfo, err := os.Open(info)
	manejoError(err)
	defer fileinfo.Close()
	bufferReaderInfo := bufio.NewReader(fileinfo)
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

	return codificacion, bloqueCodificados, bitsUltimo
}

//TieneErrores toma como parametros un archivo .ham con su archivo de informacion y verifica si tiene error
func TieneErrores(url string, info string) (bool, int, int) {
	if existeArchivo(url) && existeArchivo(info) {
		codificacion, bloqueCodificados, bitsUltimo := obtenerInformacion(info)

		file, err := os.Open(url)
		manejoError(err)
		defer file.Close()
		bufferReader := bufio.NewReader(file)

		buf := make([]byte, (codificacion)/8+1)
		hM := h(codificacion)

		byteLeidos, err := bufferReader.Read(buf)
		if byteLeidos != 0 {
			manejoError(err)
		}
		marcardor := bitsUltimo * 8
		contadorBloques := 0
		for bloqueCodificados != 0 {
			bloqueCodificados--
			auxBool := (ByteToBool(buf))
			if bloqueCodificados == 0 {
				auxBool = auxBool[:marcardor]
				hM = h(marcardor)
			}
			auxMatriz := MatrizColumna(auxBool)
			b, sindrome := hM.MultiplicarOpt(auxMatriz)
			if !b {
				if sindrome.TieneUnos() {
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

					if bloqueCodificados == 0 {
						if auxInt < bitsUltimo {
							return true, contadorBloques, auxInt

						}
					} else {
						return true, contadorBloques, auxInt
					}
				}
			}
			byteLeidos, err = bufferReader.Read(buf)
			if byteLeidos != 0 {
				manejoError(err)
			}
			if byteLeidos < len(buf) {
				//Cuando la lectura no completa el buffer.
				buf2 := make([]byte, (codificacion)/8+1-byteLeidos)
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
		fmt.Println("No existe uno de los archivos ", url, info)
	}
	return false, -1, -1
}

//TieneErroresByte toma como parametros un archivo .ham con su archivo de informacion y verifica si tiene error
func TieneErroresByte(url string, info string) (bool, int, int) {
	if existeArchivo(url) && existeArchivo(info) {
		codificacion, bloqueCodificados, bitsUltimo := obtenerInformacion(info)

		file, err := os.Open(url)
		manejoError(err)
		defer file.Close()
		bufferReader := bufio.NewReader(file)

		buf := make([]byte, (codificacion)/8+1)
		hM := matrizChequeoDeParidadH(codificacion)

		byteLeidos, err := bufferReader.Read(buf)
		if byteLeidos != 0 {
			manejoError(err)
		}
		marcardor := bitsUltimo
		contadorBloques := 0
		for bloqueCodificados != 0 {
			bloqueCodificados--
			if bloqueCodificados == 0 {
				buf = buf[:marcardor]
				hM = matrizChequeoDeParidadH(marcardor)
			}
			auxMatriz := MatrizColumnaByte(buf)
			sindrome, b := hM.Multiplicar(auxMatriz)
			if b != nil {
				if sindrome.TieneUnos() {
					auxInt := 0
					for i, fila := range sindrome.datos {
						mascara := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048}
						for _, b := range fila {
							valor, err := GetBit(b, uint(i))
							if err != nil {
								fmt.Println("Error en TieneError ", err)
							}
							if valor {
								auxInt = auxInt | mascara[i]
							}
						}
					}
					auxInt = auxInt - 1

					if bloqueCodificados == 0 {
						if auxInt < bitsUltimo {
							return true, contadorBloques, auxInt

						}
					} else {
						return true, contadorBloques, auxInt
					}
				}
			}
			byteLeidos, err = bufferReader.Read(buf)
			if byteLeidos != 0 {
				manejoError(err)
			}
			if byteLeidos < len(buf) {
				//Cuando la lectura no completa el buffer.
				buf2 := make([]byte, (codificacion)/8+1-byteLeidos)
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
		fmt.Println("No existe uno de los archivos ", url, info)
	}
	return false, -1, -1
}

func testProteccionDesproteccionArchivo() {
	ahora := time.Now()
	fmt.Println("Test Proteccion-Desproteccion ", ahora)
	var archivosPrueba = []struct {
		archivoEntrada       string
		archivoProtegido     string
		archivoProtegidoInfo string
		archivoDesprotegido  string
	}{
		//{"./prueba.txt", "./prueba.ham", "./prueba.haminfo", "./pruebaDesprotegido.txt"},
		//{"./alicia.txt", "./alicia.ham", "./alicia.haminfo", "./aliciaDesprotegido.txt"},
		{"./biblia.txt", "./biblia.ham", "./biblia.haminfo", "./bibliaDesprotegido.txt"},
	}
	var codificacionesPosibles = []int{522} //, 1035, 2060
	for _, tuplaArchivos := range archivosPrueba {
		for _, codificacion := range codificacionesPosibles {
			ahora := time.Now()
			fmt.Printf("%s\t%s\tCodificacion:%v \n", tiempo(), tuplaArchivos.archivoEntrada, codificacion)
			error := ProtegerChan(tuplaArchivos.archivoEntrada, tuplaArchivos.archivoProtegidoInfo, tuplaArchivos.archivoProtegido, codificacion)
			manejoError(error)
			fmt.Printf("%s\t\tDesprotejo\n", tiempo())
			error = Desproteger(tuplaArchivos.archivoProtegido, tuplaArchivos.archivoProtegidoInfo, tuplaArchivos.archivoDesprotegido)
			manejoError(error)
			if !compararArchivos(tuplaArchivos.archivoEntrada, tuplaArchivos.archivoDesprotegido) {
				fmt.Printf("\tError en el %s para la codificacion %d\tduracion:%s\n", tuplaArchivos.archivoEntrada, codificacion, tiempoStr(time.Now().Sub(ahora)))
			} else {
				fmt.Printf("\tExito en el %s para la codificacion %d\tduracion:%s\n", tuplaArchivos.archivoEntrada, codificacion, tiempoStr(time.Now().Sub(ahora)))
			}
		}
		fmt.Println("")
	}
}
func testProteccionDesproteccionArchivoByte() {
	ahora := time.Now()
	fmt.Println("Test Proteccion-Desproteccion Bytes", ahora)
	var archivosPrueba = []struct {
		archivoEntrada       string
		archivoProtegido     string
		archivoProtegidoInfo string
		archivoDesprotegido  string
	}{
		{"./prueba.txt", "./prueba.ham", "./prueba.haminfo", "./pruebaDesprotegido.txt"},
		{"./alicia.txt", "./alicia.ham", "./alicia.haminfo", "./aliciaDesprotegido.txt"},
		{"./biblia.txt", "./biblia.ham", "./biblia.haminfo", "./bibliaDesprotegido.txt"},
	}
	var codificacionesPosibles = []int{522, 1035, 2060} //
	for _, tuplaArchivos := range archivosPrueba {
		for _, codificacion := range codificacionesPosibles {
			ahora := time.Now()
			fmt.Printf("%s Archivo: %s Codificacion:%v \n", tiempo(), tuplaArchivos.archivoEntrada, codificacion)
			error := ProtegerByte(tuplaArchivos.archivoEntrada, tuplaArchivos.archivoProtegidoInfo, tuplaArchivos.archivoProtegido, codificacion)
			manejoError(error)
			fmt.Printf("%s Desprotejo \n", tiempo())
			error = DesprotegerByte(tuplaArchivos.archivoProtegido, tuplaArchivos.archivoProtegidoInfo, tuplaArchivos.archivoDesprotegido)
			manejoError(error)
			if !compararArchivos(tuplaArchivos.archivoEntrada, tuplaArchivos.archivoDesprotegido) {
				fmt.Printf("%s	 Error en el %s para la codificacion %d\n", tiempoStr(time.Now().Sub(ahora)), tuplaArchivos.archivoEntrada, codificacion)
			} else {
				fmt.Printf("%s	 Exito en el %s para la codificacion %d\n", tiempoStr(time.Now().Sub(ahora)), tuplaArchivos.archivoEntrada, codificacion)
			}
			fmt.Println("___________________________________________________________")

		}
	}
}
func tiempo() string {
	ahora := time.Now()
	return fmt.Sprintf("%v' %v'' %v-", ahora.Minute(), ahora.Second(), ahora.Nanosecond())
}
func tiempoStr(ahora time.Duration) string {

	return fmt.Sprintf("%v", ahora.String())

}
func testProteccionArchivosTieneErrores() {
	fmt.Println("Test Proteccion-Tiene Errores")
	var archivosPrueba = []struct {
		archivoEntrada string
		archivoSalida  string
	}{
		{"./prueba.txt", "./prueba.ham"},
		//{"./alicia.txt", "./alicia.ham"},
		//	{"./biblia.txt", "./biblia.ham"},
	}
	var codificacionesPosibles = []int{522, 1035, 2060} //
	for _, parArchivos := range archivosPrueba {
		for _, codificacion := range codificacionesPosibles {
			erro := Proteger(parArchivos.archivoEntrada, parArchivos.archivoSalida+"info", parArchivos.archivoSalida, codificacion)
			manejoError(erro)
			error, bloque, posicion := TieneErrores(parArchivos.archivoSalida, parArchivos.archivoSalida+"info")
			if error {
				fmt.Printf("Error en %s(%d):en el bloque %d en la posición %d\n", parArchivos.archivoEntrada, codificacion, bloque, posicion)
			} else {
				fmt.Printf("Exito en el %s para la codificacion %d\n", parArchivos.archivoEntrada, codificacion)
			}
		}
		fmt.Println("\n")
	}
}
func compararArchivos(path1, path2 string) bool {
	archivo1, error := os.Open(path1)
	manejoError(error)
	archivo2, error := os.Open(path2)
	manejoError(error)
	for {
		buffer1 := make([]byte, 64000)
		_, error1 := archivo1.Read(buffer1)
		buffer2 := make([]byte, 64000)
		_, error2 := archivo2.Read(buffer2)
		if error1 != nil || error2 != nil {
			if error1 == io.EOF && error2 == io.EOF {
				return true
			} else if error1 == io.EOF || error2 == io.EOF {
				return false
			} else {
				manejoError(error1)
				manejoError(error2)
			}
		}
		if !bytes.Equal(buffer1, buffer2) {
			return false
		}
	}

}
