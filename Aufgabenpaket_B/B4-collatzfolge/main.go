package main

import "fmt"

// collatz sequence input - int groeser null eary exist - grade zahl durch 2 - ungrade mal 3 plus

func main() {
	var n int
		fmt.Println("Geben Sie eine positive ganze Zahl ein:")
	_, err := fmt.Scan(&n)
	if err != nil || n <= 0 {
		fmt.Println("Ungültige Eingabe.")
		return
	}
	fmt.Print("Collatz-Folge: ")
	collatz(n)
}

func collatz(n int) {
	sequence := []int{n}
	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}
		sequence = append(sequence, n)
	}
	for i, v := range sequence {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(v)
	}

	fmt.Println()
}
