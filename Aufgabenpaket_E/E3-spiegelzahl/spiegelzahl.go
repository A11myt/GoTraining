package main

// spiegelzahl gibt die Spiegelzahl von zahl zurück.
// zahl muss eine natürliche Zahl sein.
func spiegelzahl(zahl int) int {
	if zahl < 10 {
		return zahl
	}

		return (zahl%10)*zehnerpotenz(zahl) + spiegelzahl(zahl/10)
}

func zehnerpotenz(zahl int) int {
	result := 1
	for zahl >= 10 {
		result = result * 10
		zahl = zahl / 10
	}
	return result
}
