package main

import("fmt")

func main(){

	var size int
	fmt.Println(size)

	// := Summary the expresion var and datatype, only usable inside functions not as packages.

	n := 10.23
	fmt.Println(n)

	// Ways to declara values.

	var flag bool = true
	var fly = false
	eat := true

	// Print values

	fmt.Println("1.- ", flag, "\n2.- ", fly, "\n3.- ", eat)
	fmt.Println("Basic Operations 2 + 2 = ", 2 + 2)
	fmt.Println("String + " + "Operation")
}

// Sometimes is better to specify the datatype instead that Go infere the value in case you want specific value.
// Go also does not let you compile if there is variable not used.