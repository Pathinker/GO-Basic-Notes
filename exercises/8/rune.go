package main

import ("fmt"
		"unicode/utf8"
	)


type node struct{
	name string
	cost int
}

func main(){

	n := node{"áéíóú", 12} // {} Need it instead of () since is not strictly an object
	fmt.Println(n)
	fmt.Printf("Name: %s Cost: %d\n", n.name, n.cost)
	fmt.Println("Runes used:", utf8.RuneCountInString(n.name))
	fmt.Println("Memory used:", len(n.name)) 

	if 2 % 2 == 0{
		panic("Custom Error Message Similar To Assert On Python") // iota let put default values.
	}

}

// Runes storage the aditional bytes on memory need to represent chars that are not from the ASCII code.