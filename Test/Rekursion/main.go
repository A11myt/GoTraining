package main

import (
	"fmt"
)

func main() {
	t1()
	// t2()
}

// - Legt eine int-Variable x an und einen Zeiger p auf int.
// - p zeigt auf x (`p = &x`).
// - Gibt die Adresse von x (p), den Wert von x über den Zeiger (*p), und direkt x aus.
// - Ändert den Wert von x über den Zeiger (`*p = 33`).
// - Zeigt mit `Printf` die Typen und Werte von x, &x, p, &p.
// - Demonstriert, wie Zeiger und Adressen in Go funktionieren und wie man mit ihnen arbeitet.

func t1() {
	var x int = 42
	var p *int
	p = &x

	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(x)
	*p = 33

	fmt.Printf("%T, %[1]v\n", x)
	fmt.Printf("%T, %[1]v\n", &x)
	// fmt.Printf("%T, %[1]v\n", *&x)

	fmt.Printf("%T, %[1]v\n", p)
	fmt.Printf("%T, %[1]v\n", &p)
	// fmt.Printf("%T, %[1]v\n", &*p)
}

// - Legt eine lokale int-Variable v an und gibt deren Adresse zurück.
// - Zeigt, dass man aus einer Funktion einen Zeiger auf eine lokale Variable zurückgeben kann (Go managed die Lebensdauer).

func neueIntVar() *int {
	var v int
	return &v
}

// - Legt eine int-Variable x und zwei Zeiger p, q an.
// - p zeigt auf eine neue int-Variable (über neueIntVar).
// - q übernimmt die Adresse von p.
// - p zeigt dann auf x, q wird auf nil gesetzt.
// - Gibt die Adressen von p und q aus.
// - Zeigt, wie Zeiger kopiert, neu zugewiesen und auf nil gesetzt werden können.

func t2() {
	var x int = 8
	var p, q *int
	p = neueIntVar()
	fmt.Println(p)
	q = p
	fmt.Println(q)
	p = &x
	q = nil
	fmt.Println(p)
	fmt.Println(q)
}

// rekursion
func fakultaetIte(n int) int {
	fac := 1
	for i := 2; i <= n; i++ {
		fac = fac * i
	}
	return fac
}
