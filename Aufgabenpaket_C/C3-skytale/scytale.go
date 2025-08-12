package main

// encrypt gibt die verschlüsselte Version von klartext zurück.
// Verschlüsselung mit dem Scytale-Verfahren, Schlüssel ist das Argument scytale,
// der Durchmesser des Stabes in Buchstaben.
// scytale muss eine natürliche Zahl sein.
func encrypt(klartext string, scytale int) string {
	// scytale ist die Länge des Stabes in Buchstaben.
	//count auf 0 setzen, um die Anzahl der Buchstaben zu zählen
	if scytale < 2 {
		return "Der Durchmesser muss größer als 1 sein."
	}
	// Hier sollte die Logik für die Verschlüsselung mit dem Scytale-Verfahren implementiert werden.

	return "quatsch"
}

// decrypt gibt die entschlüsselte Version von geheimtext zurück.
// Verschlüsselung mit dem Scytale-Verfahren, Schlüssel ist das Argument scytale,
// der Durchmesser des Stabes in Buchstaben.
// scytale muss eine natürliche Zahl sein.
func decrypt(geheimtext string, scytale int) string {
	return "quark"
}
