package main

// Angenommen, anzOrte ist woanders als Konstante definiert.
// const anzOrte = 3

// rundreise wendet die Nearest-Neighbor-Heuristik an:
// Starte bei Ort 0 (ausgabeseitig später +1) und gehe immer
// zum nächstgelegenen noch unbesuchten Ort. Zum Schluss zurück zum Start.
// Rückgabe: Route in 1-basiger Notation und Gesamtstrecke.
func rundreise(entfernung [anzOrte][anzOrte]int) ([anzOrte + 1]int, int) {
	var routeIdx [anzOrte + 1]int // interner Routenpuffer mit 0-basierten Indizes
	besucht := make([]bool, anzOrte) // merkt, welche Orte schon besucht sind

	routeIdx[0] = 0          // Start bei Ort-Index 0 (entspricht "Ort 1" in der Ausgabe)
	besucht[0] = true        // Startort direkt als besucht markieren
	total := 0               // aufsummierte Gesamtdistanz

	// Fülle routeIdx[1] .. routeIdx[anzOrte-1] (die Zwischenstationen).
	for i := 1; i < anzOrte; i++ {
		curr := routeIdx[i-1] // aktueller Ort ist der vorher gewählte
		minEntfernung := -1   // Sentinel: -1 bedeutet "noch nichts gefunden"
		naechster := -1       // Index des besten Kandidaten

		// Gehe die komplette Zeile 'curr' durch: Distanzen von curr -> j
		for j, d := range entfernung[curr] {
			// Überspringe bereits besuchte Orte und dich selbst
			if besucht[j] || j == curr {
				continue
			}
			// Falls erster Kandidat oder bessere (kleinere) Distanz: aktualisieren
			if minEntfernung == -1 || d < minEntfernung {
				minEntfernung = d
				naechster = j
			}
		}

		// Den gefundenen nächsten Ort in die Route übernehmen
		routeIdx[i] = naechster
		besucht[naechster] = true
		// Distanz vom aktuellen Ort zum nächsten addieren
		total += entfernung[curr][naechster]
	}

	// Am Ende zur Startposition zurückkehren
	routeIdx[anzOrte] = routeIdx[0]
	total += entfernung[routeIdx[anzOrte-1]][routeIdx[anzOrte]]

	// Für die Ausgabe 1-basiert machen (Aufgabenvorgabe)
	var routeOut [anzOrte + 1]int
	for i := 0; i <= anzOrte; i++ {
		routeOut[i] = routeIdx[i] + 1
	}
	return routeOut, total
}
