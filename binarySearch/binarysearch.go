package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var (
	guess int32
	low   int32
	high  int32
)

func main() {
	arePlaying := true
	for arePlaying {
		var validGame = performSearch()
		if !validGame {
			arePlaying = false
			fmt.Println("We don't play with cheaters. Goodbye >:(")
		} else {
			fmt.Print("Would you like to play again? ('y'/'n') ")
			arePlaying = getResponse() == 'y'
		}
	}
}

func getResponse() rune {
	reader := bufio.NewReader(os.Stdin)
	response, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}
	response = unicode.ToLower(response)
	return response
}

func performSearch() bool {
	low = 0
	high = 100
	areCorrect := false
	fmt.Printf("Let's begin. Please choose a number between %d and %d...\n", low, high)
	for !areCorrect && low <= high {
		guess = int32((high + low) / 2)
		fmt.Printf("My guess is %d.\nIs this 'c'orrect, too 'h', or too 'l'ow: ", guess)
		var response rune
		isValid := false
		for !isValid {
			response = getResponse()
			if response == 'c' {
				fmt.Println("Awesome! Good game :)")
				return true
			} else if response == 'h' {
				high = guess - 1
				isValid = true
			} else if response == 'l' {
				low = guess + 1
				isValid = true
			} else {
				fmt.Print("This is an invalid reponse. Please try 'c', 'h', or 'l': ")
			}
		}
	}
	return areCorrect
}
