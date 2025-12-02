// Go only allow compilation when all the import packages are used, if not is not possible.
// This is mainly done to reduce resource 
package main

import ("fmt"
		"os")

func main(){
	
	env := os.Getenv("HOME")

	if env == " "{
		fmt.Println("Local Enviroment does not exist")
	} else {
		fmt.Println("Local Enviroment", env)
	}
}

// If staments does not requiere () and variable asignation is :=