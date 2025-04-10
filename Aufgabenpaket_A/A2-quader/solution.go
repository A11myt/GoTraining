package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Println("Geben Sie die Länge des Rechtecks ein (in Metern):")
	fmt.Scan(&a)
	fmt.Println("Geben Sie die Breite des Rechtecks ein (in Metern):")
	fmt.Scan(&b)
	fmt.Println("Geben Sie die Höhe des Rechtecks ein (in Metern):")
	fmt.Scan(&c)
	fmt.Printf("Die eingegebenen Maße sind: Länge=%.2f m, Breite=%.2f m, Höhe=%.2f m\n", a, b, c)

	fmt.Printf("Der Volumen beträgt: %.2f m\n", rectangularVolume(a, b, c));
	fmt.Printf("Die Kantensumme beträgt: %.2f m\n", rectangularTotalEdgeWidth(a, b, c));;
	fmt.Printf("Die Oberfläche beträgt: %.2f m\n", rectangularSurfaceArea(a, b, c));
	fmt.Printf("Der Radius der Umkugel beträgt: %.2f m\n", rectangularCircumsphereRadius(a, b, c));;
	fmt.Printf("Die Länge der Raumdiagonale beträgt: %.2f m\n", rectangularBodyDiagonal(a, b, c));;
}

func rectangularVolume(length, width, height float64) float64 {
	return (length * width * height)
}

func rectangularTotalEdgeWidth(length, width, height float64) float64 {
	return (4 * (length + width + height))
}

func rectangularSurfaceArea(length, width, height float64) float64 {
	return (2 * (length * width + length * height + width * height))
}

func rectangularCircumsphereRadius(length, width, height float64) float64 {
	return 0.5 * math.Sqrt(length*length + width*width + height*height)
}

func rectangularBodyDiagonal(length, width, height float64) float64 {
	return math.Sqrt(length*length + width*width + height*height)
}


