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
		//{"./alicia.txt", "./alicia.ham", "./alicia.haminfo", "./aliciaDesprotegido.txt"},
		///{"./biblia.txt", "./biblia.ham", "./biblia.haminfo", "./bibliaDesprotegido.txt"},
	}
	var codificacionesPosibles = []int{522, 1035, 2060} //
	for _, tuplaArchivos := range archivosPrueba {
		for _, codificacion := range codificacionesPosibles {
			error := hamming.Proteger(tuplaArchivos.archivoEntrada, tuplaArchivos.archivoProtegidoInfo, tuplaArchivos.archivoProtegido, codificacion)
			manejoError(error)
			error = hamming.Desproteger(tuplaArchivos.archivoProtegido, tuplaArchivos.archivoProtegidoInfo, tuplaArchivos.archivoDesprotegido)
			manejoError(error)
			if !compararArchivos(tuplaArchivos.archivoEntrada, tuplaArchivos.archivoDesprotegido) {
				t.Errorf("-Error en el %s para la codificacion %d", tuplaArchivos.archivoEntrada, codificacion)
			} else {
				t.Logf("Exito en el %s para la codificacion %d", tuplaArchivos.archivoEntrada, codificacion)
			}
		}
	}
}

/*
func TestCorregirError(t *testing.T) {
	var archivosPrueba = []struct {
		archivoEntrada       string
		archivoProtegido     string
		archivoConError      string
		archivoProtegidoInfo string
		archivoCorregido     string
	}{
		{"./prueba.txt", "./prueba.ham", "./pruebaConError.ham", "./prueba.haminfo", "./pruebaDesprotegido.txt"},
		//{"./alicia.txt", "./alicia.ham","./aliciaConError.ham", "./alicia.haminfo", "./aliciaDesprotegido.txt"},
		///{"./biblia.txt", "./biblia.ham", "./bibliaConError.ham","./biblia.haminfo", "./bibliaDesprotegido.txt"},
	}
	var codificacionesPosibles = []int{522, 1035, 2060} //
	for _, tuplaArchivos := range archivosPrueba {
		for _, codificacion := range codificacionesPosibles {
			error := hamming.Proteger(tuplaArchivos.archivoEntrada, tuplaArchivos.archivoProtegidoInfo, tuplaArchivos.archivoProtegido, codificacion)
			manejoError(error)
			hamming.IntroducirError(tuplaArchivos.archivoProtegido, tuplaArchivos.archivoProtegidoInfo, tuplaArchivos.archivoConError)
			hamming.CorregirError(tuplaArchivos.archivoConError, tuplaArchivos.archivoProtegidoInfo, tuplaArchivos.archivoCorregido)
			manejoError(error)
			if !compararArchivos(tuplaArchivos.archivoEntrada, tuplaArchivos.archivoCorregido) {
				t.Errorf("-Error en el %s para la codificacion %d", tuplaArchivos.archivoEntrada, codificacion)
			} else {
				t.Logf("Exito en el %s para la codificacion %d", tuplaArchivos.archivoEntrada, codificacion)
			}
		}
	}
}*/

func TestProteccionArchivosTieneErrores(t *testing.T) {
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
			erro := hamming.Proteger(parArchivos.archivoEntrada, parArchivos.archivoSalida+"info", parArchivos.archivoSalida, codificacion)
			manejoError(erro)
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
