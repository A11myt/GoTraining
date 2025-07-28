package main

// potenziere gibt den Potenzwert basis hoch exponent zurück.
// exponent muss >= 0 sein.

func potenziere(basis, exponent int) int {

	var result = 1
	if exponent == 0 {
		return 1 // Jede Zahl hoch 0 ist 1
	}
	for i := 1; i <= exponent; i++ {
		result *= basis
	}
	// bitte korrigieren Sie den Rumpf dieser Funktion.
	return result
}
