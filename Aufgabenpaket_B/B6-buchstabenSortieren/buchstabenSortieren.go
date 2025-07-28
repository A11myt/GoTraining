package main

import (
	"sort"
	"strings"
)

// buchstabenSortieren gibt einen String zurück, in dem alle Buchstaben
// von wort alphabetisch sortiert sind. In wort mehrfach vorkommende
// Buchstaben kommen im resultierendem String genauso oft vor. Großbuchstaben
// in wort kommen im resultierendem String als Kleinbuchstaben vor.
func buchstabenSortieren(wort string) string {
	
	wort = strings.ToLower(wort) //kleinbuchstan
	letters := []rune(wort) // buchstaben zu char array
	sort.Slice(letters, func(i, j int) bool { // sortieren der buchstaben
		return letters[i] < letters[j]
	})
	return string(letters)
}
