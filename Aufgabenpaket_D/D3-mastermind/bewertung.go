package main

// bewertung gibt die Bewertung von versuch an Hand von code zurück.
// Die zurückgegebene Bewertung ist ein Paar (schwarze, weiße).
// versuch und code sind Arrays der Länge 4, die Farben repräsentieren.

func bewertung(versuch, code [anzStellen]string) (int, int) {
	// bitte korrigieren Sie den Rumpf dieser Funktion.
	schwarze := 0
	weiße := 0

	// Zähle schwarze Treffer
	for i, v := range versuch{
		if v == code[i] {
			schwarze++
			// Markiere die Position als bereits gezählt
			code[i] = "" // oder ein anderes Zeichen, das nicht in den Farben vorkommt
			versuch[i] = "" // um zu verhindern, dass es später als weißer Treffer gezählt wird
		}
	}
	// Zähle weiße Treffer
	for _, v := range versuch {
		for j, c := range code {
			if v == c {
				weiße++
				// Markiere die Position als bereits gezählt
				code[j] = "" // oder ein anderes Zeichen, das nicht in den Farben vorkommt
				break // Beende die innere Schleife, um doppelte Zählung zu vermeiden
			}
		}
	}
	return schwarze, weiße
}

// ich habe ein Array von strings die stellen 6 verschiedene Farben da, A - F
// und ich habe ein 4stelliges Array, das die Farben des Codes enthält
// nun darf der spieler 4 farben raten, die in meinem Array gespeichert sind
// ich muss nun die Farben vergleichen und die Anzahl der schwarzen und weißen Treffer zählen
// schwarze Treffer sind die Farben, die an der richtigen Stelle geraten wurden
// weiße Treffer sind die Farben, die zwar im Code enthalten sind, aber an der falschen
// Stelle geraten wurden
// ich muss also die Farben in den Arrays vergleichen und die Anzahl der Treffer zählen//
	// auserdem gibt es nur eine bestimmte anzahl an versuchen, die zaehle ich
