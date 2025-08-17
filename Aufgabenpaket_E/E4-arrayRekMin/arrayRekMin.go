package main

// arrayRekMin gibt den kleinsten Wert von arr im Indexbereich start bis inklusive ende-1 zurück.
// Dabei muss start<ende<=arrLen gelten.
func arrayRekMin(arr [arrLen]int, start, ende int) int {
	if start+1 == ende {
		return arr[start]
	}
	minRest := arrayRekMin(arr, start+1, ende)
	if arr[start] < minRest {
		return arr[start]
	}
	return minRest
}
