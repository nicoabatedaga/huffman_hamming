package main

import (
	"huffman_hamming/hamming"
	"testing"
)

func TestProteccionArchivos(t *testing.T) {
	var archivosPrueba = []struct {
		archivoEntrada string
		archivoSalida  string
	}{
		{"./prueba.txt", "./prueba.ham"},
		{"./alicia.txt", "./alicia.ham"},
		{"./biblia.txt", "./biblia.ham"},
	}
	var codificacionesPosibles = []int{522, 1035, 2060}
	for _, parArchivos := range archivosPrueba {
		for _, codificacion := range codificacionesPosibles {
			hamming.Proteger(parArchivos.archivoEntrada, parArchivos.archivoSalida+"info", parArchivos.archivoSalida, codificacion)
			error, bloque, posicion := hamming.TieneErrores(parArchivos.archivoSalida, parArchivos.archivoSalida+"info")
			if error {
				t.Errorf("Error en %s(%d):en el bloque %d en la posición %d", parArchivos.archivoEntrada, codificacion, bloque, posicion)
			}
		}
	}
}
