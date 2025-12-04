package main

import ("fmt")

// Go can return multiples values for once this is more easier to archive in this language.
// Other would requiere to get make a data stucture to return multiple data.

func fibonnaci(n int, values map[int] int) int{

	if n < 2{
		values[n] = 1
	} 

	if _, exist := values[n]; exist{
		return values[n]
	} 

	value := fibonnaci(n - 1, values) + fibonnaci(n - 2, values)
	values[n] = value
	return value
}

func overload(n ... int)int{ // Function overload

	var result int

	for _, data := range n{
		result += data
	}

	return result
}

func main(){

	var objetive int = 15
	storage := make(map[int] int)
	result := fibonnaci(15, storage)

	fmt.Printf("The result from %d is %d ", objetive, result)

	sum := overload(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("\nOverload addition: ", sum)
}