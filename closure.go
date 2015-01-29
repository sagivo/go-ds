package main

import "fmt"

func Mycounter(initialCount int) func() {
	theCount := initialCount
	incriment := func() {
		theCount++
		fmt.Println("The count is", theCount)
	}
	return incriment
}

func main() {
	myFunc := Mycounter(10)
	for i := 0; i < 5; i++ {
		myFunc() //use () to exicute incriment
	}
}
