package main

import "fmt"

func main() {
	printFibonacciSequence(10)
}

func printFibonacciSequence(quantity int) {
	var firstNumBefore int = 0
	var secondNumBefore int = 1

	var bufferSecondNumBefore int = secondNumBefore

	var counter int = 0
	for x := 0; x < quantity; x++ {
		bufferSecondNumBefore = secondNumBefore

		counter = firstNumBefore + secondNumBefore
		fmt.Print(counter)
		if quantity > 1 && x < (quantity-1) {
			fmt.Print(" - ")
		}
		firstNumBefore = bufferSecondNumBefore
		secondNumBefore = counter
	}
}
