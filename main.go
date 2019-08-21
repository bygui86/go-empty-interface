package main

// This conversion to an empty interface and then to the original type back has a cost for your program.
func main() {

	var i int8 = 1
	sample1(i) // regular
	sample2(i) // panic
}

//go:noinline
func sample1(i interface{}) {
	println(i)
}

// Any conversion from the inner type of an empty interface should be done after the conversion of the original type.
//go:noinline
func sample2(i interface{}) {
	n := i.(int16)
	println(n)
}
