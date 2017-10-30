package huffman

type Caracter struct {
	Caracter      string
	Ocurrencias   int
	Codigo        byte
	HijoDerecho   *Caracter
	HijoIzquierdo *Caracter
}
