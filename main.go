package main

import (
	"flag"
	"fmt"
	h "huffman_hamming/hamming"
	"huffman_hamming/huffman"
	"os"
	"runtime/trace"
)

var (
	pathIn      = flag.String("in", "./prueba.txt", " direccion del archivo de entrada")
	pathOut     = flag.String("out", "./prueba", " direccion del archivo de salida")
	codifiacion = flag.Int("cod", 512, " codificacion correspondiente al cifrado de archivos(512,1024,2048)")
	operacion   = flag.String("op", "", " define que operacion se realizara:\n\tc:comprimir,\n\td:descomprimir,\n\ta:ver arbol de codigos,\n\tp:proteger,\n\tdp:desproteger,\n\te:comprobar error,\n\ti: ingresar error,\n\tr: reparar error,\n\tm: generar archivos matrices.")
	t           = flag.Bool("t", false, " al setearlo se genera trace de la ejecucción")
)

func main() {
	flag.Parse()
	if *t {
		f, err := os.Create("trace.out")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		err = trace.Start(f)
		if err != nil {
			panic(err)
		}
		defer trace.Stop()
	}
	if *operacion == "c" {
		fmt.Println("Comprimir archivo")
		fmt.Println(huffman.Comprimir(*pathIn, *pathOut))
		return
	}

	if *operacion == "d" {
		fmt.Println("Descomprimir archivo")
		fmt.Println(huffman.Descomprimir(*pathIn, *pathOut))
		return
	}
	if *operacion == "a" {
		fmt.Println("Abrir arbol")
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
		return

	}
	if *operacion == "dp" {
		err = h.DesprotegerB(*pathIn, *pathOut)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}
	if *operacion == "e" {
		e, b, p := h.TieneErroresB(*pathIn)
		fmt.Printf("%v\n%v\n%v\n", e, b, p)
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
	if *operacion == "m" {
		h.Hamming(*codifiacion, *pathOut)
		return
	}
	fmt.Println("La operacion ingresada no es valida.")
	return
}
