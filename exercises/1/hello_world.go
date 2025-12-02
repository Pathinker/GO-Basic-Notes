package main

// Go uses only the most basic packages then need to import what is need.

import "fmt"

func main(){
	fmt.Println("Hello-World.")
}

// Go build {file} to run the application
// Similar compilation process that gcc, go build -o {compile.desirable.extension} {go.source.file}
// go build -o hello-world.o hello-world.go