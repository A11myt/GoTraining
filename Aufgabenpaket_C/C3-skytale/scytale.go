package main

// encrypt gibt die verschlüsselte Version von klartext zurück.
// Verschlüsselung mit dem Scytale-Verfahren, Schlüssel ist das Argument scytale,
// der Durchmesser des Stabes in Buchstaben.
// scytale muss eine natürliche Zahl sein.
func encrypt(klartext string, scytale int) string {
	// scytale ist die Länge des Stabes in Buchstaben.
	//count auf 0 setzen, um die Anzahl der Buchstaben zu zählen
	if scytale < 1 {
		return "Der Durchmesser muss größer als 0 sein."
	}
	//klartext in array umwwandeln und leerzeicehn entfernen
	klartext = removeSpaces(klartext)

	// Text in Array schreiben (Zeilen = scytale, Spalten = berechnet)
	// Länge des Klartexts bestimmen
	textLen := len(klartext)
	// Anzahl der Spalten berechnen (aufgerundet)
	cols := (textLen + scytale - 1) / scytale
	// Matrix mit 'scytale' Zeilen erstellen
	matrix := make([][]rune, scytale)
	for i := range matrix {
		// Jede Zeile bekommt 'cols' Spalten
		matrix[i] = make([]rune, cols)
	}

	for i, c := range klartext {
		row := i % scytale   // Zeile bestimmen
		col := i / scytale   // Spalte bestimmen
		matrix[row][col] = c // Zeichen eintragen
	}

	encrypted := ""
	for row := range make([]int, scytale) {
		for col := range make([]int, cols) {
			if matrix[row][col] != 0 {
				encrypted += string(matrix[row][col])
			}
		}
	}

	return encrypted
}

// decrypt gibt die entschlüsselte Version von geheimtext zurück.
// Verschlüsselung mit dem Scytale-Verfahren, Schlüssel ist das Argument scytale,
// der Durchmesser des Stabes in Buchstaben.
// scytale muss eine natürliche Zahl sein.
func decrypt(geheimtext string, scytale int) string {
	// Prüfe, ob der Durchmesser gültig ist
	if scytale < 2 {
		return "Der Durchmesser muss größer als 1 sein."
	}
	// Entferne Leerzeichen aus dem Geheimtext
	geheimtext = removeSpaces(geheimtext)
	textLen := len(geheimtext)
	// Berechne die Anzahl der Spalten der Matrix
	cols := (textLen + scytale - 1) / scytale

	// Initialisiere die Matrix mit der passenden Größe
	matrix := make([][]rune, scytale)
	for i := range matrix {
		matrix[i] = make([]rune, cols)
	}

	// Schreibe den Geheimtext zeilenweise in die Matrix
	idx := 0
	for row := range make([]int, scytale) {
		for col := range make([]int, cols) {
			// Fülle die Matrix nur solange noch Zeichen im Geheimtext übrig sind
			if idx < textLen {
				matrix[row][col] = rune(geheimtext[idx])
				idx++
			}
		}
	}

	// Lese die Matrix spaltenweise aus, um den Klartext zu rekonstruieren
	var klartext string
	for col := range make([]int, cols) {
		for row := range make([]int, scytale) {
			// Füge nur belegte Felder zum Klartext hinzu
			if matrix[row][col] != 0 {
				klartext += string(matrix[row][col])
			}
		}
	}
	// Gib den entschlüsselten Klartext zurück
	return klartext
}

func removeSpaces(text string) string {
	result := ""
	for _, c := range text {
		if c != ' ' {
			result += string(c)
		}
	}
	return result
}

