package main

import ("fmt")

func modify_value(list *[10]int){
	(*list)[1] = 1e9
}

func main(){

	p := fmt.Println
	array := [...]int{1,2,3,4,5,6,7,8,9,10}
	p("Before pointer access: ", array)

	modify_value(&array)
	p("After pointer access: ", array)

	ptx := &array[3]
	*ptx = 0

	p("After single pointer access: ", array)
}

// Go allows pointer and memory address manipulation like C.