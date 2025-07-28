package main
import (	"fmt" )
// istArmstrong bestimmt, ob zahl eine Armstrong-Zahl ist.
// zahl muss eine natürliche Zahl sein.
func istArmstrong(zahl int) bool {
	if zahl < 0 {
		return false // Negative Zahlen können keine Armstrong-Zahlen sein.
	}
	//split int into array of digits
	digits := []int{}
	original := zahl
	result := 0
	
	for zahl > 0 {
		// Füge die letzte Ziffer an den Anfang des Arrays
		digits = append([]int{zahl % 10}, digits...)
		fmt.Println("Ziffern:", digits)
		zahl /= 10 // Entferne die letzte Ziffer
	}
	for _, digit := range digits {
		result += potenziere(digit, len(digits))
	}
	// bitte korrigieren Sie den Rumpf dieser Funktion.
	return result == original
}

//also 153 = 1^3 + 5^3 + 3^3
//mit 4 Ziffern: 1634 = 1^4 + 6^4 + 3^4 + 4^4
