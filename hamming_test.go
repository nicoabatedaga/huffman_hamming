package main

import (
	"bytes"
	"fmt"
	"huffman_hamming/hamming"
	"io"
	"os"
	"testing"
)

func TestProteccionDesproteccionArchivo(t *testing.T) {
	var archivosPrueba = []struct {
		archivoEntrada       string
		archivoProtegido     string
		archivoProtegidoInfo string
		archivoDesprotegido  string
	}{
		{"./prueba.txt", "./prueba.ham", "./prueba.haminfo", "./pruebaDesprotegido.txt"},
		{"./alicia.txt", "./alicia.ham", "./alicia.haminfo", "./aliciaDesprotegido.txt"},
		//{"./biblia.txt", "./biblia.ham", "./biblia.haminfo", "./bibliaDesprotegido.txt"},
	}
	var codificacionesPosibles = []int{2060} //522, 1035,
	for _, tuplaArchivos := range archivosPrueba {
		for _, codificacion := range codificacionesPosibles {
			hamming.Proteger(tuplaArchivos.archivoEntrada, tuplaArchivos.archivoProtegidoInfo, tuplaArchivos.archivoProtegido, codificacion)
			hamming.Desproteger(tuplaArchivos.archivoProtegido, tuplaArchivos.archivoProtegidoInfo, tuplaArchivos.archivoDesprotegido)
			if !compararArchivos(tuplaArchivos.archivoEntrada, tuplaArchivos.archivoDesprotegido) {
				t.Errorf("-Error en el %s para la codificacion %d", tuplaArchivos.archivoEntrada, codificacion)
			} else {
				t.Logf("Exito en el %s para la codificacion %d", tuplaArchivos.archivoEntrada, codificacion)
			}
		}
	}
}

func TestProteccionArchivosTieneErrores(t *testing.T) {
	var archivosPrueba = []struct {
		archivoEntrada string
		archivoSalida  string
	}{
		{"./prueba.txt", "./prueba.ham"},
		{"./alicia.txt", "./alicia.ham"},
		//	{"./biblia.txt", "./biblia.ham"},
	}
	var codificacionesPosibles = []int{2060} //522, 1035,
	for _, parArchivos := range archivosPrueba {
		for _, codificacion := range codificacionesPosibles {
			hamming.Proteger(parArchivos.archivoEntrada, parArchivos.archivoSalida+"info", parArchivos.archivoSalida, codificacion)
			error, bloque, posicion := hamming.TieneErrores(parArchivos.archivoSalida, parArchivos.archivoSalida+"info")
			if error {
				t.Errorf("Error en %s(%d):en el bloque %d en la posición %d", parArchivos.archivoEntrada, codificacion, bloque, posicion)
			} else {
				t.Logf("Exito en el %s para la codificacion %d", parArchivos.archivoEntrada, codificacion)
			}
		}
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
