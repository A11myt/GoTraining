package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var min = 1   // min := int32(1)
	var max = 100 // max := int32(100)
	var guesses int32 = 0
	var closestGuess = 0
	var closestDiff int = max - min + 1
	randNum := min + rand.Intn(max-min+1)

	for {
		guess := guessNumber(min, max)
		guesses++
		diff := abs(int(guess) - randNum)

		if guess == randNum {
			fmt.Println("Congratulations! You guessed the number after {{0}}!", guesses)
			break
		}

		if guesses == 1 {
			if guess+6 < randNum && guess-6 > randNum {
				fmt.Println("HEAT!")
			} else {
				fmt.Println("Eiskalt!")
			}
		} else {
			if diff < closestDiff {
				fmt.Println("Hotter! Das ist jetzt dein nächster Tipp.")
				closestGuess = guess
				closestDiff = diff
			} else {
				fmt.Printf("Colder! Dein bisher nächster Tipp war: %d\n", closestGuess)
			}
		}
	}
	fmt.Println(randNum)
}

// guess a number between min and max
func guessNumber(min, max int) int {
	var guess int
	fmt.Printf("Guess a number between %d and %d: ", min, max)
	fmt.Scanf("%d", &guess)
	if guess < min || guess > max {
		fmt.Printf("Please enter a number between %d and %d.\n", min, max)
		return guessNumber(min, max)
	}
	return guess
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
