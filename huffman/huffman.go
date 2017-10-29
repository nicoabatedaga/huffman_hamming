package huffman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func manejoError(err error) {
	if err != nil {
		panic(fmt.Sprintf("hubo un error, %v", err))
	}
}

func Huffman() {
	fmt.Println("empiezo el huffman")
	archivoLectura, err := os.Open("./archivo")
	manejoError(err)
	defer archivoLectura.Close()
	scanner := bufio.NewScanner(archivoLectura)
	var listaDeCaracteres []Caracter
	ocurrencias := map[string]int{}

	//itero sobre las lineas y cargo un mapa simple
	ocurrencias = recorroArchivoYCuento(scanner)
	fmt.Println(fmt.Sprintf("Ocurrencias: %v", ocurrencias))
	//Recorro el mapa siple y genero una lista de Caracter
	listaDeCaracteres = parseToSliceOfCaracters(ocurrencias)
	fmt.Println(fmt.Sprintf("Tengo la lista: %v", listaDeCaracteres))
	return
}

func recorroArchivoYCuento(scanner *bufio.Scanner) map[string]int {
	ocurrencias := map[string]int{}
	for scanner.Scan() {
		ocurrencias["newLine"] = ocurrencias["newLine"] + 1 //Sumo el salto de linea
		caracteresDeLinea := strings.Split(scanner.Text(), "")
		for i := 0; i < len(caracteresDeLinea); i++ {
			ocurrencias[caracteresDeLinea[i]]++
		}
	}
	return ocurrencias
}

// Agarro el mapa de caracteres y genero un slice con structs
// del tipo Caracter
func parseToSliceOfCaracters(mapita map[string]int) []Caracter {
	listaDeCaracteres := []Caracter{}
	for k, v := range mapita {
		fmt.Printf("key[%s] value[%v]\n", k, v)
		newCaracter := Caracter{}
		newCaracter.Caracter = k
		newCaracter.Ocurrencias = v
		listaDeCaracteres = append(listaDeCaracteres, newCaracter)
	}
	return listaDeCaracteres
}
