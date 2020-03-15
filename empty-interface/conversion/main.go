package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {}
func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {}
func (c Cat) Speak() string {
	return "Meow!"
}

type Llama struct {}
func (l Llama) Speak() string {
	return "Lllllllllama!"
}

type JavaProgrammer struct {}
func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}

func main() {
	animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}}
	
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
	
	// we cannot just assign animals to beings, because the type does not match
	beings := make([]interface{}, len(animals))
	// we need to explicitly convert each Animal to an interface{} before, with a loop
	for i, v := range animals {
		beings[i] = v
	}
}
