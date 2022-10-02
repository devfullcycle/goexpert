package main

func main() {
	// Memória -> Endreço -> Valor
	a := 10
	var ponteiro *int = &a
	*ponteiro = 20
	b := &a
	*b = 30
	println(a)
}
