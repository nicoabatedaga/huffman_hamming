package huffman

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func Huffman() {
	fmt.Println("empiezo el huffman")
	archivo, err := os.Open("./archivo")
	if err != nil {
		panic(fmt.Sprintf("hubo un error, %v", err))
	}
	scanner := bufio.NewScanner(archivo)
	wg := new(sync.WaitGroup)
	for scanner.Scan() {
		wg.Add(1)
		imprimirLinea(scanner.Text())
		wg.Done()
	}
	wg.Wait()
	return
}

func imprimirLinea(linea string) {
	fmt.Println(linea)
}
