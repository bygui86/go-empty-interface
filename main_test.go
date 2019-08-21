package main_test

import (
	"testing"
)

var x MultipleFieldStructure
var y *MultipleFieldStructure

type MultipleFieldStructure struct {
	a int
	b string
	c float32
	d float64
	e int32
	f bool
	g uint64
	h *string
	i uint16
}

//go:noinline
func emptyInterfaceToType(i interface{}) {
	s := i.(MultipleFieldStructure)
	x = s
}

//go:noinline
func emptyInterfaceToPointer(i interface{}) {
	s := i.(*MultipleFieldStructure)
	y = s
}

//go:noinline
func typed(s MultipleFieldStructure) {
	x = s
}

func BenchmarkWithType(b *testing.B) {
	s := MultipleFieldStructure{a: 1, h: new(string)}
	for i := 0; i < b.N; i++ {
		typed(s)
	}
}

func BenchmarkWithEmptyInterfaceToType(b *testing.B) {
	s := MultipleFieldStructure{a: 1, h: new(string)}
	for i := 0; i < b.N; i++ {
		emptyInterfaceToType(s)
	}
}

func BenchmarkWithEmptyInterfaceToPointer(b *testing.B) {
	s := MultipleFieldStructure{a: 1, h: new(string)}
	for i := 0; i < b.N; i++ {
		emptyInterfaceToPointer(&s)
	}
}
