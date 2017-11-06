package huffman

type Caracter struct {
	Caracter      string
	Ocurrencias   int
	CodigoString  string
	Codigo        *[]byte
	HijoDerecho   *Caracter
	HijoIzquierdo *Caracter
}
