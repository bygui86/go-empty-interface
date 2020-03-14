package pointer_empty_interface_test

import (
	"testing"
	
	"github.com/bygui86/go-empty-interface/benchmarking"
)

var pointerEmptyInterfaceStructure *benchmarking.MultipleFieldStructure

func BenchmarkWithEmptyInterfaceToPointer(b *testing.B) {
	s := benchmarking.MultipleFieldStructure{A: 1, H: new(string)}
	for i := 0; i < b.N; i++ {
		emptyInterfaceToPointer(&s)
	}
}

//go:noinline
func emptyInterfaceToPointer(i interface{}) {
	s := i.(*benchmarking.MultipleFieldStructure)
	pointerEmptyInterfaceStructure = s
}
