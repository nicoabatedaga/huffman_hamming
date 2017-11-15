package huffman

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/jlabath/bitarray"
)

var (
	mapita map[string]string
)

func manejoError(err error) {
	if err != nil {
		panic(fmt.Sprintf("hubo un error, %v", err))
	}
}

func Huffman() {
	fmt.Println(" - empiezo el huffman - ")
	hacerBitArray()
	var listaDeCaracteres []Caracter
	listaDeCaracteres = getListaDeCaracteres(recorroArchivoYCuento("./archivoNuevo"))
	// fmt.Println(fmt.Sprintf("Lista de caracteres obtenida: %v", listaDeCaracteres))
	raizDelArbol := generarArbol(listaDeCaracteres)
	// fmt.Println(fmt.Sprintf("La raiz de arbol que obtenemos es: %v", raizDelArbol))
	// imprimirEnProfundidad(*raizDelArbol)
	treeAsMap(raizDelArbol)
	fmt.Println(mapita)
	recorroArchivoYEscriboArchivoCodificado("./archivoNuevo", "./archivoNuevoComprimido.huff")
	return
}

// Leo linea a linea e itero en cada linea por cada caracter y guardo en un mapa
// que ingreso por caracter y obtengo la cantidad de apariciones
func recorroArchivoYCuento(nfr string) map[string]int {
	archivoLectura, err := os.Open(nfr)
	manejoError(err)
	scanner := bufio.NewScanner(archivoLectura)
	ocurrencias := map[string]int{}
	for scanner.Scan() {
		caracteresDeLinea := strings.Split(scanner.Text(), "")
		for i := 0; i < len(caracteresDeLinea); i++ {
			ocurrencias[caracteresDeLinea[i]]++
		}
		ocurrencias["newLine"]++ //Sumo el salto de linea
	}
	ocurrencias["newLine"]-- //Saco el enter de sobra si es que no hay otra linea
	archivoLectura.Close()
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
		list = append(list[:0], list[0+1:]...) // Elimino el primer elemento de la lista
		c2 := list[0]
		list = append(list[:0], list[0+1:]...) // Elimino el primer elemento de la lista
		list = append(list, crearPadre(c1, c2))
		if len(list) > 1 {
			sort.Slice(list, func(i, j int) bool {
				return list[i].Ocurrencias < list[j].Ocurrencias
			})
		}
	}
	arbol := evaluarNodos(list[0]) // segundo parametro nil, ya que no tiene padre
	return arbol
	// return &list[0]
}

// Agarra dos nodos, crea un nodo padre y los pone como sus hijos
// y a sus ocurrencias las setea como la suma de los hijos
func crearPadre(c1, c2 Caracter) Caracter {
	padre := Caracter{}
	padre.Caracter = "padre"
	padre.Codigo = nil
	padre.CodigoString = ""
	padre.Ocurrencias = c1.Ocurrencias + c2.Ocurrencias
	padre.HijoIzquierdo = &c1
	padre.HijoDerecho = &c2
	return padre
}

func evaluarNodos(raiz Caracter) *Caracter {
	if raiz.HijoIzquierdo != nil {
		raiz.HijoIzquierdo.CodigoString = "0"
		raiz.HijoIzquierdo = evaluarSubNodos(*raiz.HijoIzquierdo)
	}
	if raiz.HijoDerecho != nil {
		raiz.HijoDerecho.CodigoString = "1"
		raiz.HijoDerecho = evaluarSubNodos(*raiz.HijoDerecho)
	}
	return &raiz
}

func evaluarSubNodos(raiz Caracter) *Caracter {
	if raiz.HijoIzquierdo != nil {
		raiz.HijoIzquierdo.CodigoString = raiz.CodigoString + "0"
		raiz.HijoIzquierdo = evaluarSubNodos(*raiz.HijoIzquierdo)
	}
	if raiz.HijoDerecho != nil {
		raiz.HijoDerecho.CodigoString = raiz.CodigoString + "1"
		raiz.HijoDerecho = evaluarSubNodos(*raiz.HijoDerecho)
	}
	return &raiz
}

func imprimirNodo(c *Caracter) {
	ocurrencias := c.Ocurrencias
	caracter := c.Caracter
	codigo := c.CodigoString
	hd := c.HijoDerecho
	hi := c.HijoIzquierdo
	if caracter == "padre" {
		fmt.Println(fmt.Sprintf("Es PADRE con %v ocurrencias y codigo %v, y los hijos son\n	HIJO IZQUIERDO: %v	Ocurrencias:%v\n	HIJO DERECHO: %v	Ocurrencias:%v\n", ocurrencias, codigo, hi.CodigoString, hi.Ocurrencias, hd.CodigoString, hd.Ocurrencias))
	} else {
		fmt.Println(fmt.Sprintf("Es un nodo HOJA del caracter %v, con %v ocurrencias y codigo %v\n", caracter, ocurrencias, codigo))
	}
	return
}

func imprimirArbol(raiz *Caracter) {
	imprimirNodo(raiz)
	if raiz.HijoIzquierdo != nil {
		imprimirArbol(raiz.HijoIzquierdo)
	}
	if raiz.HijoDerecho != nil {
		imprimirArbol(raiz.HijoDerecho)
	}
	return
}

func imprimirEnProfundidad(raiz Caracter) {
	imprimirNodo(&raiz)
	listaImprimirEnProfundidad := []*Caracter{}
	if raiz.HijoIzquierdo != nil {
		listaImprimirEnProfundidad = append(listaImprimirEnProfundidad, raiz.HijoIzquierdo)
	}
	if raiz.HijoDerecho != nil {
		listaImprimirEnProfundidad = append(listaImprimirEnProfundidad, raiz.HijoDerecho)
	}
	var next *Caracter
	for e := listaImprimirEnProfundidad[0]; e != nil; e = next {
		imprimirNodo(e)
		if e.HijoIzquierdo != nil {
			listaImprimirEnProfundidad = append(listaImprimirEnProfundidad, e.HijoIzquierdo)
		}
		if e.HijoDerecho != nil {
			listaImprimirEnProfundidad = append(listaImprimirEnProfundidad, e.HijoDerecho)
		}
		if len(listaImprimirEnProfundidad) <= 1 {
			break
		}
		listaImprimirEnProfundidad = append(listaImprimirEnProfundidad[:0], listaImprimirEnProfundidad[0+1:]...)
		next = listaImprimirEnProfundidad[0]
	}
}

func treeAsMap(c *Caracter) {
	mapita = map[string]string{}
	treeAsMapAux(c)
}
func treeAsMapAux(c *Caracter) {
	if c.Caracter != "padre" {
		mapita[c.Caracter] = c.CodigoString
	}
	if c.HijoIzquierdo != nil {
		treeAsMapAux(c.HijoIzquierdo)
	}
	if c.HijoDerecho != nil {
		treeAsMapAux(c.HijoDerecho)
	}
}

//
func recorroArchivoYEscriboArchivoCodificado(nfr, nfw string) {
	archivoLectura, err := os.Open(nfr)
	manejoError(err)
	archivoEscritura, err := os.Create(nfw)
	manejoError(err)
	scanner := bufio.NewScanner(archivoLectura)
	writer := bufio.NewWriter(archivoEscritura)
	enc := gob.NewEncoder(writer)
	var wg sync.WaitGroup
	channelToWrite := make(chan *bitarray.BitArray)
	go func(channelToWrite chan *bitarray.BitArray) {
		ba := bitarray.New(8)
		iterba := 0

		for scanner.Scan() {
			// spliteo en todos los caracteres de la linea
			caracteresDeLinea := strings.Split(scanner.Text(), "")
			// itero por cada caracter
			for i := 0; i < len(caracteresDeLinea); i++ {
				// spliteo cada codigo de 0 y 1 de cada caracter
				splitCerosUnos := strings.Split(mapita[caracteresDeLinea[i]], "")
				// recorro sobre este split de 0 y 1
				for iSplit := 0; iSplit < len(splitCerosUnos); iSplit++ {
					// Si es un 1 entonces le hago un set al bitarray
					if splitCerosUnos[iSplit] == "1" {
						ba.Set(iterba)
					}
					// incremento el contador para recorrer el bitarray
					iterba++
					// compruebo que no se me pase del tamaño 8, si se pasa entonces
					// meto el bitarray en el canal y reseteo el bitarray y el iterador
					if iterba == 8 {
						/*---------------*/
						// Agregar un contador de las cosas que agrego al canal
						// wait group
						wg.Add(1)
						channelToWrite <- ba
						/*---------------*/
						ba.Fill(0)
						iterba = 0
					}
				}
			}
			//TODO: Insertar new line
			splitNewLine := strings.Split(mapita["newLine"], "")
			for iSplit := 0; iSplit < len(splitNewLine); iSplit++ {
				if splitNewLine[iSplit] == "1" {
					ba.Set(iterba)
				}
				iterba++
				if iterba == 8 {
					/*---------------*/
					// Agregar un contador de las cosas que agrego al canal
					// wait group
					wg.Add(1)
					channelToWrite <- ba
					/*---------------*/
					ba.Fill(0)
					iterba = 0
				}
			}
		}

		return
	}(channelToWrite)

	for scanner.Scan() {
		caracteresDeLinea := strings.Split(scanner.Text(), "")
		for i := 0; i < len(caracteresDeLinea); i++ {
			enc.Encode(mapita[caracteresDeLinea[i]])
			// writer.WriteString(mapita[caracteresDeLinea[i]])
			fmt.Println("imprimo", mapita[caracteresDeLinea[i]])
		}
		writer.WriteString(mapita["newLine"])
	}
	writer.Flush()
	archivoEscritura.Close()
	archivoLectura.Close()
}

func hacerBitArray() {
	ba := bitarray.New(8)
	ba.Set(0)
	ba.Set(2)
	ba.Set(4)
	ba.Set(5)
	ba.Set(6)
	ba.Set(7)
	bi := ba.GetData()
	fmt.Println(ba)
	ba.Fill(0)
	fmt.Println(ba)
	fmt.Println(bi)
}
