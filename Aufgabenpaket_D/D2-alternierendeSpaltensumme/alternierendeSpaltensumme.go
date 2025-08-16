package main

// alternierendeSpaltensumme überprüft arr daraufhin, ob in der letzten
// Zeile die alternierenden Spaltensummen stehen, korrigiert diese ggf.
// und gibt die Anzahl der fehlerhaften Summen zurück.
func alternierendeSpaltensumme(arr *[anzZeilen][anzSpalten]int) int {
	var fehlerhafteSummen int = 0

	// Iteriere über die Spalten, außer der letzten (die soll ja die Summe sein)
	// for i := 0; i < anzSpalten-1; i++ {
	for i := range arr[0][:anzSpalten] {
		spaltensumme := 0 // Summe der aktuellen Spalte
		// for j := 0; j < anzZeilen-1; j++ { // Letzte Zeile ignorieren
		for j := range arr[:anzZeilen-1] {
			if j%2 == 0 {
				spaltensumme += arr[j][i]
			} else {
				spaltensumme -= arr[j][i]
			}
		}

		// Prüfe, ob die Summe in der letzten Zeile korrekt ist
		if spaltensumme != arr[anzZeilen-1][i] {
			fehlerhafteSummen++                // Zähle fehlerhafte Summen
			
			//error hier
			arr[anzZeilen-1][i] = spaltensumme // Korrigiere die Summe
		}
	}
	// Bitte korrigieren Sie den Rumpf der Funktion.
	return fehlerhafteSummen
}
