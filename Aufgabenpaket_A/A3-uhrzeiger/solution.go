
package main

import (
	"fmt"
	"math"
)

func main() {
	var hour, minute int

	fmt.Println("Bitte geben Sie die aktuelle Uhrzeit ein.")
	fmt.Print("Stunde (0–23): ")
	fmt.Scan(&hour)

	fmt.Print("Minute (0–59): ")
	fmt.Scan(&minute)

	currentTime := float64(hour) + float64(minute)/60

	fmt.Printf("\n🕰️ Uhrzeit: %02d:%02d\n", hour, minute)
	fmt.Printf("🕒 Stundenzeiger steht auf %.2f Grad\n", getHourHandAngle(currentTime))
	fmt.Printf("🕕 Minutenzeiger steht auf %.2f Grad\n", getMinuteHandAngle(currentTime))
	fmt.Printf("↔️ Winkel zwischen den Zeigern: %.2f Grad\n", getAngleBetweenHands(currentTime))
}

func getHourHandAngle(currentTime float64) float64 {
	hour, minute := separateHourAndMinute(currentTime)
	return float64(hour%12)*30 + float64(minute)*0.5
}

func getMinuteHandAngle(currentTime float64) float64 {
	_, minute := separateHourAndMinute(currentTime)
	return float64(minute) * 6
}

func getAngleBetweenHands(currentTime float64) float64 {
	hourAngle := getHourHandAngle(currentTime)
	minuteAngle := getMinuteHandAngle(currentTime)
	diff := math.Abs(hourAngle - minuteAngle)
	if diff > 180 {
		return 360 - diff
	}
	return diff
}

func separateHourAndMinute(currentTime float64) (int, int) {
	hour := int(math.Floor(currentTime))
	minute := int((currentTime - float64(hour)) * 60)
	if minute == 60 {
		hour++
		minute = 0
	}
	return hour, minute
}

