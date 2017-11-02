package hamming

import (
	"fmt"
)

//Hamming es el programa de prueba
func Hamming() {
	fmt.Println("-Hamming-")

	codificacion := 512
	fmt.Println(fmt.Sprintf("Codificacion: %v", codificacion))
<<<<<<< HEAD
	
	
=======
	pruebaMatriz()
>>>>>>> develop_joaco
	return
}

func pruebaMatriz() {
	k := [][]bool{{false, true, true}, {true, true, false}}
	ma := Matriz{datos: k}
	r := [][]bool{{false, true}, {true, true}, {false, true}}
	mb := Matriz{datos: r}
	var m Matriz
	m = ma.Multiplicar(mb)
	fmt.Println("Matriz Resultante:")
	fmt.Println(m.ToString())
}

func esPotenciaDeDos( ent int)bool {
	return ((ent != 0) && ((ent & (ent-1)) == 0))
}

func bitsInformacion() {

}

func bitsParidad() {

}

func h() {

}

func g() {

}

//Codificar  Cifra el archivo de entrada
func Codificar() {

}

func tieneError() {

}

//Decodificar descrifra el archivo de entrada
func Decodificar() {

}

//AgregarError a un archivo
func AgregarError() {

}
