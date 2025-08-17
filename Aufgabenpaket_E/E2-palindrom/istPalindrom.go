package main

import ( "strings")
// istPalindrom prüft, ob zeichenfolge ein Palindrom ist.
func istPalindrom(zeichenfolge string) bool {

	zeichenfolge = strings.ToLower(zeichenfolge)
	zeichenfolge = strings.ReplaceAll(zeichenfolge, " ", "")
	// Vergleiche das Zeichen am Anfang mit dem Zeichen am Ende.
	if len(zeichenfolge) <= 1 { // Wenn die Länge der Zeichenfolge 0 oder 1 ist, ist sie ein Palindrom.
		return true
	}
	if zeichenfolge[0] != zeichenfolge[len(zeichenfolge)-1] { // Wenn das erste Zeichen nicht gleich dem letzten Zeichen ist, ist es kein Palindrom.
		return false
	}
	// Bitte korrigieren Sie den Rumpf dieser Funktion.
	return istPalindrom(zeichenfolge[1 : len(zeichenfolge)-1]) // Rufe die Funktion rekursiv auf, um die inneren Zeichen zu überprüfen.
}
