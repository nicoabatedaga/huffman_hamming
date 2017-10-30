package huffman

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func manejoError(err error) {
	if err != nil {
		panic(fmt.Sprintf("hubo un error, %v", err))
	}
}

func Huffman() {
	fmt.Println("empiezo el huffman")
	archivoLectura, err := os.Open("./archivoNuevo")
	manejoError(err)
	defer archivoLectura.Close()
	scanner := bufio.NewScanner(archivoLectura)
	var listaDeCaracteres []Caracter
	listaDeCaracteres = getListaDeCaracteres(recorroArchivoYCuento(scanner))
	fmt.Println(fmt.Sprintf("Lista de caracteres obtenida: %v", listaDeCaracteres))
	raizDelArbol := generarArbol(listaDeCaracteres)
	fmt.Println(fmt.Sprintf("La raiz de arbol que obtenemos es: %v", raizDelArbol))
	return
}

// Leo linea a linea e itero en cada linea por cada caracter y guardo en un mapa
// que ingreso por caracter y obtengo la cantidad de apariciones
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

// Dado el mapa generado en la lectura del archivo devuelvo una lista
// de la estructura Caracter
func getListaDeCaracteres(mapa map[string]int) []Caracter {
	listaDeCaracteres := []Caracter{}
	for k, v := range mapa {
		newCaracter := Caracter{}
		newCaracter.Caracter = k
		newCaracter.Ocurrencias = v
		listaDeCaracteres = append(listaDeCaracteres, newCaracter)
	}
	sort.Slice(listaDeCaracteres, func(i, j int) bool {
		return listaDeCaracteres[i].Ocurrencias < listaDeCaracteres[j].Ocurrencias
	})
	return listaDeCaracteres
}

// Dada una lista de caracteres con ocurrencias, genero el arbol.
// Tomo los dos nodos con menos ocurrencias, los uno en un nodo padre
// el cual su numero de ocurrencia es la suma de los dos hijos
func generarArbol(list []Caracter) *Caracter {
	for len(list) >= 2 {
		c1 := list[0]
		list = append(list[:0], list[0+1:]...)
		c2 := list[0]
		list = append(list[:0], list[0+1:]...)
		fmt.Println(fmt.Sprintf("Tenemos el caracter c1: %v", c1))
		fmt.Println(fmt.Sprintf("Tenemos el caracter c2: %v", c2))
		fmt.Println(fmt.Sprintf("Tenemos la lista: %v", list))
		list = append(list, crearPadre(c1, c2))
		if len(list) > 1 {
			sort.Slice(list, func(i, j int) bool {
				return list[i].Ocurrencias < list[j].Ocurrencias
			})
		}
		fmt.Println(fmt.Sprintf("Tenemos la lista despues de meter el padre: %v \n\n", list))
	}
	return &list[0]
}

// Agarra dos nodos, crea un nodo padre y los pone como sus hijos
// y a sus ocurrencias las setea como la suma de los hijos
func crearPadre(c1, c2 Caracter) Caracter {
	padre := Caracter{}
	padre.Caracter = "padre"
	padre.Ocurrencias = c1.Ocurrencias + c2.Ocurrencias
	padre.HijoIzquierdo = &c1
	padre.HijoDerecho = &c2
	return padre
}
