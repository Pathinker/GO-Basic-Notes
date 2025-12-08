package main

import("fmt"
		"time")

func counter(ch chan bool){

	for i := 0; i < 10; i++{

		fmt.Println("We are couting with this one! 🗣️🔥 ", i)
	}

	ch <- true // If not returned this value a deadlock happens, since it does not return nothing.
}

func update(ch chan bool){
	ch <- true
}

func main(){

	ch := make(chan bool) // It is possible to get a buffered channel with storaged values.
	// Buffer channels does not need to be insta read it to keep going, since has a buffer.
	// Buffered channel only stops when the buffer is full waiting for a reda that can cause a deadlock if not handle correctly.

	ch2 := make(chan bool)
	go counter(ch)
	<-ch 
	fmt.Println("Counter Finished With Channel")

	go update(ch2)
	<- ch2
	fmt.Println("Changed Done")

	fmt.Println("--- Select Example ---")

	go counter(ch)
	go update(ch2)

	// Select allows to run all the routines and wait for each one to be done, instead of running each other and wait for it, like the previous example.

	for i := 0; i < 2; i++{
		select{
		case <- ch:
			fmt.Println("First Execution")
		case <- ch2:
			fmt.Println("Second Execution")
		}
	}

	fmt.Println("--- Timeout Example ---")

	go update(ch2)

	select{
		case <- ch2:
			fmt.Println("Done")
		case <- time.After(0 * time.Second): // Deny wait.
			fmt.Println("Time Out")
	}
}

// Go routines allow to have lighweight threads managed by go runtime.
// This lightweight threads share the same memory adress instead of having their own one.
// Need to secure this shared memory with mutex or locks.
// Channels are meant to be used to syncronize go routines it waits.