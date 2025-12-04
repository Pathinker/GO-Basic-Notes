package main

import ("fmt"
		"math/rand/v2"
		"time"		
)

// Go allows to declare variables at if staments

func main(){

	if random := rand.IntN(10); random > 5{

		fmt.Println("Better than average: ", random)
	} else {
		fmt.Println("Lower than average: ", random)
	}

	var m int = 0;

	for m < 10{
		fmt.Println("M value: ", m)
		m++
	}

	for j := 0; j < 1; j++{
		fmt.Println("Execution Time: ", j)
	}

	for l := range 1 {
		fmt.Printf("This is Amazing (%d) isn't it?", l)
	}

	p := fmt.Println
	now := time.Now()
	p("\nNo idea this was even possible, congrats for people that made this possible: ", now)

	switch{

		case now.Hour() < 20:
			fmt.Println("Lock In")
		case now.Hour() > 20 && now.Hour() < 22:
			fmt.Println("Lock Off")
		default:
			p("Sleep")
	}
}

// Go is flexible with the declaration of if and for statments it could use diferent common sintaxis from declarative and non declaritive languages
// Printf can output a value %d stands for %i on C, printf gives the string the specified format, PrintLn does not but makes a \n.