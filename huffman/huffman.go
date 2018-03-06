package huffman

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/jlabath/bitarray"
)

var (
	mapita                      map[string]string
	mapitaInvertido             map[string]string
	listaDeCaracteres           []Caracter
	stringParaMatchearLectura   string
	counterBits                 int64
	counterReadBits             int64
	tamanioArchivoCodificado    int64
	tamanioArchivo              int64
	tamanioArchivoDescomprimido int64
)

func manejoError(err error) {
	if err != nil {
		panic(fmt.Sprintf("hubo un error, %v", err))
	}
}

func Huffman() {
	fmt.Println(" - empiezo el huffman - ")
	nombreArchivo := "./archivo"
	nombreArchivoComprimido := Comprimir(nombreArchivo)
	fmt.Println(fmt.Sprintf("El archivo comprimido se llama: %v", nombreArchivoComprimido))
	leoArchivoEnBytes(nombreArchivoComprimido)
	// mapita = map[string]string{}
	// mapitaInvertido = map[string]string{}
	// stringParaMatchearLectura = ""
	// counterBits = 0
	// counterReadBits = 0
	// Comprimir("archivo.comprimido.huffDescomprimido")
	if tamanioArchivo != tamanioArchivoDescomprimido {
		panic(fmt.Sprintf("Los tamaños no coinciden %d - %d", tamanioArchivo, tamanioArchivoDescomprimido))
	}
	fmt.Println(fmt.Sprintf("OK"))
	return
}

func Comprimir(nombreArchivo string) string {
	nombreArchivoComprimido := nombreArchivo + ".comprimido.huff"
	listaDeCaracteres = getListaDeCaracteres(recorroArchivoYCuento(nombreArchivo))
	raizDelArbol := generarArbol(listaDeCaracteres)
	treeAsMap(raizDelArbol)
	fmt.Println(fmt.Sprintf("Mapa del arbol: %v", mapita))
	guardoListaDeCaracteresEnArchivo(nombreArchivo)
	recorroArchivoYEscriboArchivoCodificado(nombreArchivo, nombreArchivoComprimido)
	return nombreArchivoComprimido
}

func guardoListaDeCaracteresEnArchivo(nombreArchivo string) {
	nombreArchivoListaCaracteres := nombreArchivo + ".listaCaracteres"
	archivoEscrituraListaCaracteres, err := os.Create(nombreArchivoListaCaracteres)
	defer archivoEscrituraListaCaracteres.Close()
	manejoError(err)
	arr := make([]string, len(mapita))
	i := 0
	for k := range mapita {
		arr[i] = k
		i++
	}
	sort.Strings(arr)
	writer := bufio.NewWriter(archivoEscrituraListaCaracteres)
	for _, v := range arr {
		writer.WriteString(fmt.Sprintf("%v-cod-%v\n", v, mapita[v]))
	}
	//for k, v := range mapita {
	//	writer.WriteString(fmt.Sprintf("%v-cod-%v\n", k, v))
	//}
	err = writer.Flush()
	manejoError(err)
}

// Leo linea a linea e itero en cada linea por cada caracter y guardo en un mapa
// que ingreso por caracter y obtengo la cantidad de apariciones
func recorroArchivoYCuento(nfr string) map[string]int {
	archivoLectura, err := os.Open(nfr)
	defer archivoLectura.Close()
	manejoError(err)
	fileInfo, _ := archivoLectura.Stat()
	tamanioArchivo = fileInfo.Size()
	fmt.Println(fmt.Sprintf("file size: %vbytes - %vMb", fileInfo.Size(), fileInfo.Size()/1048576.0))
	// scanner := bufio.NewScanner(archivoLectura)

	// reader := bufio.NewReader(archivoLectura)
	// linea, err := reader.ReadString('\n')

	ocurrencias := map[string]int{}

	bf := bufio.NewReader(archivoLectura)

	// there are a few possible loop termination
	// conditions, so just start with an infinite loop.
	for {
		// reader.ReadLine does a buffered read up to a line terminator,
		// handles either /n or /r/n, and returns just the line without
		// the /r or /r/n.
		line, isPrefix, err := bf.ReadLine()

		// loop termination condition 1:  EOF.
		// this is the normal loop termination condition.
		if err == io.EOF {
			ocurrencias["newLine"]--
			if ocurrencias["newLine"] == 0 {
				delete(ocurrencias, "newLine")
			}
			fmt.Println(fmt.Sprintf("\nse leyo un EOF %v\n", err.Error()))
			break
		}

		// loop termination condition 2: some other error.
		// Errors happen, so check for them and do something with them.
		if err != nil {
			log.Fatal(err)
		}

		// loop termination condition 3: line too long to fit in buffer
		// without multiple reads.  Bufio's default buffer size is 4K.
		// Chances are if you haven't seen a line terminator after 4k
		// you're either reading the wrong file or the file is corrupt.
		if isPrefix {
			log.Fatal("Error: Unexpected long line reading", archivoLectura.Name())
		}

		// success.  The variable line is now a byte slice based on on
		// bufio's underlying buffer.  This is the minimal churn necessary
		// to let you look at it, but note! the data may be overwritten or
		// otherwise invalidated on the next read.  Look at it and decide
		// if you want to keep it.  If so, copy it or copy the portions
		// you want before iterating in this loop.  Also note, it is a byte
		// slice.  Often you will want to work on the data as a string,
		// and the string type conversion (shown here) allocates a copy of
		// the data.  It would be safe to send, store, reference, or otherwise
		// hold on to this string, then continue iterating in this loop.
		// fmt.Println(string(line))

		caracteresDeLinea := strings.Split(string(line), "")
		// fmt.Println(fmt.Sprintf("caracteres de linea:%v", caracteresDeLinea))
		for i := 0; i < len(caracteresDeLinea); i++ {
			// fmt.Println(fmt.Sprintf("Sumo: %v", caracteresDeLinea[i]))
			ocurrencias[caracteresDeLinea[i]]++
		}
		// fmt.Println(fmt.Sprintf("Sumo: newLine"))
		ocurrencias["newLine"]++ //Sumo el salto de linea

	}

	/*scan := scanner.Scan()
	for scan {
		caracteresDeLinea := strings.Split(scanner.Text(), "")
		// fmt.Println(fmt.Sprintf("caracteres de linea:%v", caracteresDeLinea))
		for i := 0; i < len(caracteresDeLinea); i++ {
			// fmt.Println(fmt.Sprintf("Sumo: %v", caracteresDeLinea[i]))
			ocurrencias[caracteresDeLinea[i]]++
		}
		scan = scanner.Scan()
		if scan {
			// fmt.Println(fmt.Sprintf("Sumo: newLine"))
			ocurrencias["newLine"]++ //Sumo el salto de linea
		}
	}*/

	// fmt.Println(fmt.Sprintf("Resto: newLine"))
	// ocurrencias["newLine"]-- //Saco el enter de sobra si es que no hay otra linea
	// fmt.Println(fmt.Sprintf("Ocurrencias: %v", ocurrencias))
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
		if listaDeCaracteres[i].Ocurrencias == listaDeCaracteres[j].Ocurrencias {
			return listaDeCaracteres[i].Caracter < listaDeCaracteres[j].Caracter
		}
		return listaDeCaracteres[i].Ocurrencias < listaDeCaracteres[j].Ocurrencias
	})
	fmt.Println(fmt.Sprintf("lista de caracteres: %v", listaDeCaracteres))
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
	invertMap()
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

func invertMap() {
	mapitaInvertido = map[string]string{}
	for x, y := range mapita {
		mapitaInvertido[y] = x
	}
}

//
func recorroArchivoYEscriboArchivoCodificado(nfr, nfw string) {
	counterBits = 0
	bufferEscritura := []byte{}
	archivoLectura, err := os.Open(nfr)
	manejoError(err)
	defer archivoLectura.Close()
	archivoEscritura, err := os.Create(nfw)
	manejoError(err)
	defer archivoEscritura.Close()
	// scanner := bufio.NewScanner(archivoLectura)
	writer := bufio.NewWriter(archivoEscritura)
	ba := bitarray.New(8)
	iterba := 0

	// for scanner.Scan() {
	// 	if err := scanner.Err(); err != nil {
	// 		fmt.Println("error", err)
	// 	}
	// 	// spliteo en todos los caracteres de la linea
	// 	caracteresDeLinea := strings.Split(scanner.Text(), "")
	// 	caracteresDeLinea = append(caracteresDeLinea, "newLine")
	// 	// fmt.Println(fmt.Sprintf("Linea: %v", caracteresDeLinea))
	// 	// itero por cada caracter
	// 	for i := 0; i < len(caracteresDeLinea); i++ {
	// 		// spliteo cada codigo de 0 y 1 de cada caracter
	// 		// fmt.Println(fmt.Sprintf("Codigo del caracter %v: %v", caracteresDeLinea[i], mapita[caracteresDeLinea[i]]))
	// 		splitCerosUnos := strings.Split(mapita[caracteresDeLinea[i]], "")
	// 		// recorro sobre este split de 0 y 1
	// 		for iSplit := 0; iSplit < len(splitCerosUnos); iSplit++ {
	// 			// Si es un 1 entonces le hago un set al bitarray
	// 			// fmt.Println(fmt.Sprintf("bits: %v - iterba: %v", ba.String(), iterba))
	// 			if splitCerosUnos[iSplit] == "1" {
	// 				ba.Set(iterba)
	// 			}
	// 			// incremento el contador para recorrer el bitarray
	// 			iterba++
	// 			counterBits++
	// 			// fmt.Println(fmt.Sprintf("bits: %v - iterba: %v", ba.String(), iterba))
	// 			// compruebo que no se me pase del tamaño 8, si se pasa entonces
	// 			// meto el bitarray en el canal y reseteo el bitarray y el iterador
	// 			if iterba == 8 {
	// 				bufferEscritura = append(bufferEscritura, ba.GetData()...)
	// 				ba.Fill(0)
	// 				iterba = 0
	// 			}
	// 			if len(bufferEscritura) >= 256 {
	// 				meterEnArchivoBinario(*writer, bufferEscritura)
	// 				bufferEscritura = []byte{}
	// 			}
	// 		}
	// 	}
	// }

	bf := bufio.NewReader(archivoLectura)

	for {
		line, isPrefix, err := bf.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if isPrefix {
			log.Fatal("Error: Unexpected long line reading", archivoLectura.Name())
		}
		// fmt.Println(string(line))

		caracteresDeLinea := strings.Split(string(line), "")
		if v, err := bf.Peek(1); err != io.EOF {
			caracteresDeLinea = append(caracteresDeLinea, "newLine")
		} else {
			fmt.Println("LLEGO AL FINAL", v)
		}
		//fmt.Println(fmt.Sprintf("Linea: %v", caracteresDeLinea))
		// itero por cada caracter
		for i := 0; i < len(caracteresDeLinea); i++ {
			// spliteo cada codigo de 0 y 1 de cada caracter
			//fmt.Println(fmt.Sprintf("iracter %v: %v", caracteresDeLinea[i], mapita[caracteresDeLinea[i]]))
			splitCerosUnos := strings.Split(mapita[caracteresDeLinea[i]], "")
			// recorro sobre este split de 0 y 1
			for iSplit := 0; iSplit < len(splitCerosUnos); iSplit++ {
				// Si es un 1 entonces le hago un set al bitarray
				//fmt.Println(fmt.Sprintf("bits: %v - iterba: %v", ba.String(), iterba))
				if splitCerosUnos[iSplit] == "1" {
					ba.Set(iterba)
				}
				// incremento el contador para recorrer el bitarray
				iterba++
				counterBits++
				//fmt.Println(fmt.Sprintf("bits: %v - iterba: %v", ba.String(), iterba))
				// compruebo que no se me pase del tamaño 8, si se pasa entonces
				// meto el bitarray en el canal y reseteo el bitarray y el iterador
				if iterba == 8 {
					bufferEscritura = append(bufferEscritura, ba.GetData()...)
					ba.Fill(0)
					iterba = 0
				}
				if len(bufferEscritura) >= 256 {
					meterEnArchivoBinario(*writer, bufferEscritura)
					bufferEscritura = []byte{}
				}
			}
		}
	}
	if iterba != 0 {
		bufferEscritura = append(bufferEscritura, ba.GetData()...)
	}
	meterEnArchivoBinario(*writer, bufferEscritura)
	// writer.Flush()
	fileInfo, _ := archivoEscritura.Stat()

	tamanioArchivoCodificado = fileInfo.Size()
	fmt.Println(fmt.Sprintf("file size: %vbytes - %vMb", fileInfo.Size(), fileInfo.Size()/1048576.0))
}

func meterEnArchivoBinario(writer bufio.Writer, byteArray []byte) {
	//if writer.Available() < len(byteArray) {
	//}
	// fmt.Println(fmt.Sprintf("Guardo en el archivo: %v", ba.String()))
	_, err := writer.Write(byteArray)
	manejoError(err)
	err = writer.Flush()
	manejoError(err)
	// fmt.Println("imprimi ", i, "bytes en el archivo")
}

/*
DESCOMPRIMIR
leo el archivo byte a byte y convierto esos byte a [8]bits
agarro todos esos arreglos de bits creados los convierto a string
[string,string,string] voy recorriendo cada string del arreglo caracter por caracter
si el caracter/cadena actual no matchea con ningun codigo entonces le concateno el siguiente caracter
consulto nuevamente si matchea, si matchea, grabo en el archivo el caracter asociado y vacio la cadena
con la que voy a ir concatenando los 0's y 1's para descomprimir
*/

func leoArchivoEnBytes(nombreArchivo string) {
	counterReadBits = 0
	bufferEscritura := []string{}
	archivoLecturaBytes, err := os.Open(nombreArchivo)
	manejoError(err)
	defer archivoLecturaBytes.Close()
	archivoEscritura, err := os.Create(nombreArchivo + "Descomprimido")
	manejoError(err)
	defer archivoEscritura.Close()
	writer := bufio.NewWriter(archivoEscritura)
	scanner := bufio.NewScanner(archivoLecturaBytes)
	buf := make([]byte, 256)
	bufferReader := bufio.NewReader(archivoLecturaBytes)
	stringParaMatchearLectura = ""
	_, err = bufferReader.Read(buf)
	manejoError(err)
	for err != io.EOF {
		//for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Println("error", err)
		}
		//x := scanner.Bytes()
		// fmt.Println(fmt.Sprintf("scanner.Bytes():%v", x))
		for _, b := range buf {
			s := fmt.Sprintf("%b", b)
			// fmt.Println(fmt.Sprintf("s:%v", s))
			for it := len(s); it < 8; it++ {
				s = "0" + s
			}
			spliteado := strings.Split(s, "")
			// for caracters := 0; caracters < 8; caracters++ {
			for _, val := range spliteado {
				counterReadBits++
				if counterReadBits <= counterBits {
					stringParaMatchearLectura = stringParaMatchearLectura + val
				}
				// fmt.Println(fmt.Sprintf("Busco %v en el mapa", stringParaMatchearLectura))
				if stringParaMatchearLectura == "0100001010" {
					fmt.Println("encuentra esto")
				}
				if mapitaInvertido[stringParaMatchearLectura] != "" {
					// fmt.Println(fmt.Sprintf("Encuntro coincidencia del codigo %v y caracter %v", stringParaMatchearLectura, mapitaInvertido[stringParaMatchearLectura]))
					// meterEnArchivo(*writer, mapitaInvertido[stringParaMatchearLectura])
					bufferEscritura = append(bufferEscritura, mapitaInvertido[stringParaMatchearLectura])
					stringParaMatchearLectura = ""
				}
			}
			if len(bufferEscritura) >= 256 {
				meterEnArchivo(*writer, bufferEscritura)
				bufferEscritura = []string{}
			}
		}

		_, err = bufferReader.Read(buf)
		if err != io.EOF {
			manejoError(err)
		}

	}

	meterEnArchivo(*writer, bufferEscritura)
	fileInfo, _ := archivoEscritura.Stat()
	tamanioArchivoDescomprimido = fileInfo.Size()
	fmt.Println(fmt.Sprintf("file size: %vbytes - %vMb", fileInfo.Size(), fileInfo.Size()/1048576.0))
}

func meterEnArchivo(writer bufio.Writer, as []string) {
	// p := mpb.New(
	// 	// override default (80) width
	// 	mpb.WithWidth(100),
	// 	// override default "[=>-]" format
	// 	mpb.WithFormat("╢▌▌░╟"),
	// 	// override default 100ms refresh rate
	// 	mpb.WithRefreshRate(120*time.Millisecond),
	// )
	// name := "Escribiendo archivo descomprimido"
	// bar := p.AddBar(int64(len(as)),
	// 	// Prepending decorators
	// 	mpb.PrependDecorators(
	// 		// StaticName decorator with minWidth and no extra config
	// 		// If you need to change name while rendering, use DynamicName
	// 		decor.StaticName(name, len(name), 0),
	// 		// ETA decorator with minWidth and no extra config
	// 		decor.ETA(4, 0),
	// 	),
	// 	// Appending decorators
	// 	mpb.AppendDecorators(
	// 		// Percentage decorator with minWidth and no extra config
	// 		decor.Percentage(5, 0),
	// 	),
	// )

	for _, s := range as {
		if s == "newLine" {
			// fmt.Println("")
			writer.WriteString("\n")
		} else {
			// fmt.Printf("%s", s)
			writer.WriteString(s)
		}
		// bar.Increment()
	}
	// p.Stop()

	writer.Flush()
}
