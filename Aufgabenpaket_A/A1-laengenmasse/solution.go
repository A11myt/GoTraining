package main

import "fmt"

func main() {
	ConvertLength()
}

func ConvertLength() {

  var meter float64;
	fmt.Println("Geben Sie eine Länge (in Metern) ein:")
  fmt.Scan(&meter)


    centimeters := meter * 100

    kilometers := float64(meter) / 1000
    fmt.Printf("meter in kilometers: %v\n", kilometers);

    millimeters := meter * 1000
    fmt.Printf("meter in millimeters: %v\n", millimeters);

    inches := float64(centimeters) / 2.54
    fmt.Printf("centimeters in inches: %v\n", inches);

    nauticalMiles := kilometers / 1.852
    fmt.Printf("kilometers in nautical miles: %v\n", nauticalMiles);

    lightYears := kilometers / 9.461730472580800
    fmt.Printf("kilometers in light years: %v\n", lightYears);
  }
