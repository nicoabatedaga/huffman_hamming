package huffman

func SliceDelete(slice []interface{}, i int) []interface{} {
	aux := append(slice[:i], slice[i+1:]...)
	return aux
}
