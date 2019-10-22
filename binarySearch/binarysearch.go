package main

import "fmt"

var (
	guess int8
	low   int8
	high  int8
)

func main() {
	low = 0
	high = 100
	arePlaying := true
	for arePlaying {
		var validGame = performSearch()
		if !validGame {
			arePlaying = false
			fmt.Println("We don't play with cheaters. Goodbye >:(")
		}
		else{
			//TODO: take in user input to play again
		}
	}
}

func performSearch() bool {
	areCorrect := false
	fmt.Printf("Let's begin. Please choose a number between %d and %d", low, high)
	for !areCorrect && low <= high {
		guess = int8((high + low) / 2)
		fmt.Printf("My guess is %d.\nIs this 'c'orrect, too 'h', or too 'l'ow", guess)
		//TODO: take in user input to determine guess validity
	}
	return areCorrect
}
