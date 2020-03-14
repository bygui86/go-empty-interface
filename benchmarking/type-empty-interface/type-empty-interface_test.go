package type_empty_interface_test

import (
	"testing"
	
	"github.com/bygui86/go-empty-interface/benchmarking"
)

var typeEmptyInterfaceStructure benchmarking.MultipleFieldStructure

func BenchmarkWithEmptyInterfaceToType(b *testing.B) {
	s := benchmarking.MultipleFieldStructure{A: 1, H: new(string)}
	for i := 0; i < b.N; i++ {
		emptyInterfaceToType(s)
	}
}

//go:noinline
func emptyInterfaceToType(i interface{}) {
	s := i.(benchmarking.MultipleFieldStructure)
	typeEmptyInterfaceStructure = s
}
