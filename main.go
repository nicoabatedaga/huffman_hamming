package main

import (
	"fmt"
	"huffman_hamming/hamming"
	"huffman_hamming/huffman"
	"time"
)

func main() {
	s := time.Now().Nanosecond()
	fmt.Println("Empiezo ejecucion")
	hamming.Hamming()
	huffman.Huffman()
	f := time.Now().Nanosecond()
	attempt := f - s
	fmt.Printf("tiempo transcurrido: %vns", attempt)
}
