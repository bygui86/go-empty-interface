package _type_test

import (
	"testing"
	
	"github.com/bygui86/go-empty-interface/benchmarking"
)

var typeStructure benchmarking.MultipleFieldStructure

func BenchmarkWithType(b *testing.B) {
	s := benchmarking.MultipleFieldStructure{A: 1, H: new(string)}
	for i := 0; i < b.N; i++ {
		typed(s)
	}
}

//go:noinline
func typed(s benchmarking.MultipleFieldStructure) {
	typeStructure = s
}
