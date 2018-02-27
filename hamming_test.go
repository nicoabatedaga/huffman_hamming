package main

import (
	"bytes"
	"fmt"
	"huffman_hamming/hamming"
	"io"
	"os"
	"testing"
	"time"
)

func TestProteccionDesproteccionArchivo(t *testing.T) {
	ahora := time.Now()
	fmt.Println("Test Proteccion-Desproteccion ", ahora)
	var archivosPrueba = []struct {
		archivoEntrada      string
		archivoProtegido    string
		archivoDesprotegido string
	}{
		{"./prueba.txt", "./prueba.ham", "./pruebaDesprotegido.txt"},
		{"./alicia.txt", "./alicia.ham", "./aliciaDesprotegido.txt"},
		{"./biblia.txt", "./biblia.ham", "./bibliaDesprotegido.txt"},
		{"./moby.txt", "./moby.ham", "./mobyDesprotegido.txt"},
		{"./metamorfosis.txt", "./metamorfosis.ham", "./metamorfosisDesprotegido.txt"},
		{"./hyde.txt", "./hyde.ham", "./hydeDesprotegido.txt"},
	}
	var codificacionesPosibles = []int{512, 1024, 2048} //, 1024, 2048
	for _, tuplaArchivos := range archivosPrueba {
		for _, codificacion := range codificacionesPosibles {
			ahora := time.Now()
			fmt.Printf("%s\t%s\tCodificacion:%v \n", tiempo(), tuplaArchivos.archivoEntrada, codificacion)
			error := hamming.Proteger(tuplaArchivos.archivoEntrada, tuplaArchivos.archivoProtegido, codificacion)
			manejoError(error)
			fmt.Printf("%s\t\tDesprotejo\n", tiempo())
			error = hamming.Desproteger(tuplaArchivos.archivoProtegido, tuplaArchivos.archivoDesprotegido)
			manejoError(error)
			if !compararArchivos(tuplaArchivos.archivoEntrada, tuplaArchivos.archivoDesprotegido) {
				fmt.Printf("\tError en el %s para la codificacion %d\tduracion:%s\n", tuplaArchivos.archivoEntrada, codificacion, tiempoStr(time.Now().Sub(ahora)))
				t.Errorf("\tError en el %s para la codificacion %d\t\n", tuplaArchivos.archivoEntrada, codificacion)
			} else {
				fmt.Printf("\tExito en el %s para la codificacion %d\tduracion:%s\n", tuplaArchivos.archivoEntrada, codificacion, tiempoStr(time.Now().Sub(ahora)))
			}
		}
		fmt.Println("")
	}
}

func tiempo() string {
	ahora := time.Now()
	return fmt.Sprintf("%v' %v'' %v-", ahora.Minute(), ahora.Second(), ahora.Nanosecond())
}
func tiempoStr(ahora time.Duration) string {

	return fmt.Sprintf("%v", ahora.String())

}

func TestCorregirError(t *testing.T) {
	var archivosPrueba = []struct {
		archivoEntrada      string
		archivoProtegido    string
		archivoConError     string
		archivoCorregido    string
		archivoDesprotegido string
	}{
	//	{"./prueba.txt", "./prueba.ham", "./pruebaConError.ham", "./pruebaCorregido.ham", "./pruebaDesprotegido.txt"},
	//	{"./alicia.txt", "./alicia.ham", "./aliciaConError.ham", "./aliciaCorregido.ham", "./aliciaDesprotegido.txt"},
	//	{"./biblia.txt", "./biblia.ham", "./bibliaConError.ham", "./bibliaCorregido.ham", "./bibliaDesprotegido.txt"},
	//	{"./moby.txt", "./moby.ham", "./mobyConError.ham", "./mobyCorregido.ham", "./mobyDesprotegido.txt"},
	//	{"./metamorfosis.txt", "./metamorfosis.ham", "./metamorfosisConError.ham", "./metamorfosisCorregido.ham", "./metamorfosisDesprotegido.txt"},
	//	{"./hyde.txt", "./hyde.ham", "./hydeConError.ham", "./hydeCorregido.ham", "./hydeDesprotegido.txt"},
	}
	var codificacionesPosibles = []int{512} //, 1024, 2048
	for _, tuplaArchivos := range archivosPrueba {
		for _, codificacion := range codificacionesPosibles {
			error := hamming.Proteger(tuplaArchivos.archivoEntrada, tuplaArchivos.archivoProtegido, codificacion)
			manejoError(error)
			fmt.Println("-----Introducir Errores")
			hamming.IntroducirError(tuplaArchivos.archivoProtegido, tuplaArchivos.archivoConError)
			fmt.Println("-----Corregir Errores")
			hamming.CorregirError(tuplaArchivos.archivoConError, tuplaArchivos.archivoCorregido)
			fmt.Println("-----Desproteger")
			hamming.Desproteger(tuplaArchivos.archivoCorregido, tuplaArchivos.archivoDesprotegido)
			manejoError(error)
			if !compararArchivos(tuplaArchivos.archivoEntrada, tuplaArchivos.archivoDesprotegido) {
				t.Errorf("-Error en el %s para la codificacion %d", tuplaArchivos.archivoEntrada, codificacion)
			} else {
				t.Logf("Exito en el %s para la codificacion %d", tuplaArchivos.archivoEntrada, codificacion)
			}
		}
	}
}

func TestProteccionArchivosTieneErrores(t *testing.T) {
	ahora := time.Now()
	fmt.Println("Test TieneErrores ", ahora)
	var archivosPrueba = []struct {
		archivoEntrada string
		archivoSalida  string
	}{
	//	{"./prueba.txt", "./prueba.ham"},
	//	{"./alicia.txt", "./alicia.ham"},
	//	{"./biblia.txt", "./biblia.ham"},
	//	{"./moby.txt", "./moby.ham"},
	//	{"./metamorfosis.txt", "./metamorfosis.ham"},
	//	{"./hyde.txt", "./hyde.ham"},
	}
	var codificacionesPosibles = []int{512, 1024, 2048} //512, 1024, 2048
	for _, parArchivos := range archivosPrueba {
		for _, codificacion := range codificacionesPosibles {
			fmt.Printf("%s\t%s\tCodificacion:%v \n", tiempo(), parArchivos.archivoEntrada, codificacion)
			erro := hamming.Proteger(parArchivos.archivoEntrada, parArchivos.archivoSalida, codificacion)
			manejoError(erro)
			error, bloque, posicion := hamming.TieneErrores(parArchivos.archivoSalida)
			if error {
				t.Errorf("Error en %s(%d):en el bloque %d en la posición %d", parArchivos.archivoEntrada, codificacion, bloque, posicion)
			} else {
				t.Logf("Exito en el %s para la codificacion %d", parArchivos.archivoEntrada, codificacion)
			}
		}
		fmt.Println("")
	}
}

func TestProteccionArchivosAgregarErroress(t *testing.T) {
	ahora := time.Now()
	fmt.Println("Test Agregar Errores ", ahora)
	var archivosPrueba = []struct {
		archivoEntrada  string
		archivoSalida   string
		archivoConError string
	}{
	//{"./prueba.txt", "./prueba.ham", "./pruebaConError.ham"},
	//{"./alicia.txt", "./alicia.ham", "./aliciaConError.ham"},
	//{"./biblia.txt", "./biblia.ham", "./bibliaConError.ham"},
	}
	var codificacionesPosibles = []int{1024} //512,, 2048
	for _, parArchivos := range archivosPrueba {
		for _, codificacion := range codificacionesPosibles {
			fmt.Printf("%s\t%s\tCodificacion:%v \n", tiempo(), parArchivos.archivoEntrada, codificacion)
			erro := hamming.Proteger(parArchivos.archivoEntrada, parArchivos.archivoSalida, codificacion)
			manejoError(erro)
			error, bloque, posicion := hamming.TieneErrores(parArchivos.archivoSalida)
			if error {
				t.Errorf("-\tError:%s(%d)-(%d,%d)", parArchivos.archivoEntrada, codificacion, bloque, posicion)
			} else {
				t.Logf("+\tExito:%s(%d)", parArchivos.archivoEntrada, codificacion)
			}
			hamming.IntroducirErrorNR(parArchivos.archivoSalida, parArchivos.archivoConError, 0, 9)

			error, bloque, posicion = hamming.TieneErrores(parArchivos.archivoConError)
			fmt.Println("Tiene errores: ", error, bloque, posicion)
			if !error {
				t.Errorf("-\t(IE)Error:%s(%d)", parArchivos.archivoEntrada, codificacion)
			} else {
				t.Logf("+\t(IE)Exito:%s(%d)-(%d,%d)", parArchivos.archivoEntrada, codificacion, bloque, posicion)
			}
		}
		fmt.Println("")
	}
}
func manejoError(err error) {
	if err != nil {
		panic(fmt.Sprintf("hubo un error, %v", err))
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
