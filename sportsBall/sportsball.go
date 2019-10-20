package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var teamA int32 = rand.Int31n(100)
	var teamB int32 = rand.Int31n(100)
	if teamA > teamB {
		fmt.Printf("Team A won with %d points!", teamA)
	} else if teamB > teamA {
		fmt.Printf("Team B won with %d points!", teamB)
	} else {
		fmt.Printf("It's a tie with %d points!", teamA)
	}
}
