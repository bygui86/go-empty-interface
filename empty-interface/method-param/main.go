package main

import "fmt"

// This conversion to an empty interface and then to the original type back has a cost for your program.
func main() {
	var i interface{}
	describe(i)
	
	i = 42
	describe(i)
	
	i = "hello"
	describe(i)
	
	var i int8 = 1
	printer1(i) // regular
	printer2(i) // panic
}

//go:noinline
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

//go:noinline
func printer1(i interface{}) {
	fmt.Println(i)
}

// Any conversion from the inner type of an empty interface should be done after the conversion of the original type.
//go:noinline
func printer2(i interface{}) {
	n := i.(int16)
	fmt.Println(n)
}
