package hamming

import (
	"fmt"
)

//Hamming es el programa de prueba
func Hamming() {
	fmt.Println("-Hamming-")

	codificacion := 512
	fmt.Println(fmt.Sprintf("Codificacion: %v, Bits Paridad: %v, Bits Información: %v", 
		codificacion,
		bitsParidad(codificacion),
		bitsInformacion(codificacion)))
			

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

func bitsInformacion(ent int) int {
	return(ent - (bitsParidad(ent)))
}

//bitsParidad Devuelve que candidad de bit corresponderian a bits de paridad para 
//determinada codifciacion, si se desea para evitar iterar, podria hacerse un mapeo
//con los 5 valores de codificacion que exiten.
func bitsParidad(ent int) int{
	n:=0
	for i:=2;i<=ent;i=i*2{
		n++;
	}
	if(ent == 7){
		return 3;}
	if(ent == 31){
		return 5;
	}
	return n
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
