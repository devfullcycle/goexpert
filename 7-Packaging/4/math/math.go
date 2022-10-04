package math

var X string = "Hello World"

type math struct {
    a int
    b int
    C int
}

func NewMath(a, b int) math {
    return math{a: a, b: b}
}

func (m math) Add() int {
    return m.a + m.b
}