package main
import (
	// "fmt"
	// "testing"		
)

// istRechtwinklig prüft, ob das durch die Seitenlängen x, y, z gebildete
// Dreieck rechtwinklig ist.
func istRechtwinklig(x, y, z int) bool {
	if x <= 0 || y <= 0 || z <= 0 {
		return false
	}

	a, b, c := x*x, y*y, z*z
	return a+b == c || a+c == b || b+c == a
}
