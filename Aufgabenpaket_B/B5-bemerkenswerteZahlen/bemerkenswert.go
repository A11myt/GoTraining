package main

// bemerkenswert gibt die nächste natürliche Zahl größer start zurück, deren
// Quadrat die Summe der natürlichen Zahlen von 1 bis n ist (mit beliebigem n).

// fmt.Println("Die erste bemerkenswerte Zahl größer als", start, "ist", bemerkenswert(start))

func bemerkenswert(start int) int {
	for x := start + 1; ; x++ {
		square := potenziere(x, 2)
		sum := 0 // Summe der Zahlen von 1 bis i
		for i := 1; sum <= square; i++ {
			sum += i // Summe der Zahlen von 1 bis i
			if sum == square {
				return x
			}
		}
	}
}

// 	var n int = start
// 	var summe int = 0
// 	var square = potenziere(start, 2)
//
// 	for i := 1; i <= square; i++ {
// 		summe += i
// 	}
// 	if summe == square {
// 		return n
// 	} else {
// 		return bemerkenswert(start + 1)
// 	}
// }
