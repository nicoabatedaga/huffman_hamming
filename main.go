package main

import (
	"flag"
	"fmt"
	h "huffman_hamming/hamming"
	"huffman_hamming/huffman"
	"time"
)

var (
	pathIn      = flag.String("in", "./prueba.txt", " direccion del archivo de entrada")
	pathOut     = flag.String("out", "./prueba", " direccion del archivo de salida")
	codifiacion = flag.Int("cod", 512, " codificacion correspondiente al cifrado de archivos(512,1024,2048)")
	operacion   = flag.String("op", "", " define que operacion se realizara:\n\tc:comprimir,\n\td:descomprimir,\n\tp:proteger,\n\tdp:desproteger,\n\te:comprobar error,\n\ti: ingresar error,\n\tr: reparar error.")
)

func main() {
	flag.Parse()
	ahora := time.Now()
	if *operacion == "c" {
		fmt.Println(huffman.Comprimir(*pathIn, *pathOut))
		return
	}

	if *operacion == "d" {
		fmt.Println(huffman.Descomprimir(*pathIn, *pathOut))
		return
	}
	var err error
	if *operacion == "p" {
		if *codifiacion == 512 || *codifiacion == 1024 || *codifiacion == 2048 {
			err = h.ProtegerB(*pathIn, *pathOut, *codifiacion)
		} else {
			fmt.Println("No corresponde a una codifiacion valida ", *codifiacion)
			err = h.ProtegerB(*pathIn, *pathOut, 512)
		}
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(time.Now().Sub(ahora).Seconds(), " segundos")
		return

	}
	if *operacion == "dp" {
		err = h.DesprotegerB(*pathIn, *pathOut)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(time.Now().Sub(ahora).Seconds(), " segundos")
		return
	}
	if *operacion == "e" {
		arrErrores := h.TieneErroresB(*pathIn)
		for i := 0; i < len(arrErrores); i += 2 {
			fmt.Printf("%v\n%v\n", arrErrores[i], arrErrores[i+1])
		}
		return
	}
	if *operacion == "i" {
		h.IntroducirError(*pathIn, *pathOut)
		return
	}
	if *operacion == "r" {
		h.CorregirError(*pathIn, *pathOut)
		return
	}
	fmt.Println("La operacion ingresada no es valida.")
	return
}
