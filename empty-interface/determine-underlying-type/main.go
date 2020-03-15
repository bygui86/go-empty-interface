package main

import "fmt"

func main() {
	var t interface{}

	t = functionOfSomeType("struct")
	determineUnderlyingType(t)
	t = functionOfSomeType("string")
	determineUnderlyingType(t)
	t = functionOfSomeType("int")
	determineUnderlyingType(t)
	t = functionOfSomeType("bool")
	determineUnderlyingType(t)
	t = functionOfSomeType("bool-pointer")
	determineUnderlyingType(t)
	t = functionOfSomeType("int-pointer")
	determineUnderlyingType(t)
}

func determineUnderlyingType(t interface{}) {
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", t) // t has type bool
	case int:
		fmt.Printf("integer %d\n", t) // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	}
}

type Animal struct {
	kind         string
	numberOfArms int
}

func functionOfSomeType(inputType string) interface{} {
	switch inputType {
	case "struct":
		return Animal{"cat", 4}
	case "string":
		return "this is a string"
	case "int":
		return 3
	case "bool":
		return false
	case "bool-pointer":
		value := true
		return &value
	case "int-pointer":
		value := 42
		return &value
	default:
		return "default string"
	}
}
