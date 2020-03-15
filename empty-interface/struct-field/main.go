package main

import "fmt"

type Dog struct {
	Age interface{}
}

func main() {
	example1()
	
	fmt.Println()
	
	example2()
}

func example1() {
	dog := Dog{}
	dog.Age = "3"
	fmt.Printf("%#v %T\n", dog.Age, dog.Age)
	
	dog.Age = 3
	fmt.Printf("%#v %T\n", dog.Age, dog.Age)
	
	dog.Age = "not really an age"
	fmt.Printf("%#v %T\n", dog.Age, dog.Age)
}

func example2() {
	dog := Dog{Age: "3"}
	fmt.Printf("%#v %T\n", dog.Age, dog.Age)
	
	dog.Age = 3
	fmt.Printf("%#v %T\n", dog.Age, dog.Age)
	
	dog.Age = "not really an age"
	fmt.Printf("%#v %T\n", dog.Age, dog.Age)
}
