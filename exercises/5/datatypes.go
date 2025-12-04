package main

import("fmt")

func main(){

	p := fmt.Println

	var array [5]int
	p("First Declared Array: ", array)

	array2 := [2]string{"Bye", "World"}
	p("Second Declared Array: ", array2)

	dinamic_lenght := [...]float64{1.23, 4.23, 3.14}
	p("Third Dinamic Lenght: ", len(dinamic_lenght))

	dinamic_vector_python := make([]int, 0)

	for i := range(6){
		dinamic_vector_python = append(dinamic_vector_python, i) 
		// Since is not an object like python need to be reassingn and not called as a function, like dinamic_vector_python.append(i)
	}

	p(dinamic_vector_python)

	storage := make(map[int]int)
	
	for i:= range(10){
		if i < 2{
			storage[i] = 1
			continue 
		}
		storage[i] = storage[i - 1] + storage[i - 2] // Fibonaci
	}

	p(storage)

}

// Go does not allow to overwrite an previous array, need to change the identifier.