package main

// istPerfekt bestimmt, ob zahl eine perfekte Zahl ist.
// zahl muss eine natürliche Zahl sein.
func istPerfekt(zahl int) bool {
	// bitte korrigieren Sie den Rumpf dieser Funktion.
	// 28 = 1 + 2 + 4 + 7 + 14
	if zahl < 1 {
		return false
	}

	if zahl == 1 {
		return false
	}
	summe := 0
	// Schleife von 1 bis zur Hälfte der Zahl
	for i := 1; i <= zahl/2; i++ {
		// Prüfen, ob i ein Teiler der Zahl ist
		if zahl%i == 0 {
			// Teiler zur Summe hinzufügen
			summe += i
		}
	}
	return summe == zahl
}
