
package main

import (
	"fmt"
	"math"
)


func main() {
    var currentTime float64

    fmt.Println("Bitte geben Sie die aktuelle Uhrzeit ein (z. B. 10.30 für 10:30 Uhr):")
    _, err := fmt.Scan(&currentTime)
    if err != nil {
        fmt.Println("Fehler:", err)
        return
    }

    fmt.Printf("Der Stundenzeiger steht auf %.2f Grad\n", getHourHandAngle(currentTime))
    fmt.Printf("Der Minutenzeiger steht auf %.2f Grad\n", getMinuteHandAngle(currentTime))
}


func getHourHandAngle(currentTime float64) float64 {
    hour, minute := separateHourAndMinute(currentTime)
    return float64(hour)*30 + float64(minute)*0.5
}

func getMinuteHandAngle(currentTime float64) float64 {
    _, minute := separateHourAndMinute(currentTime)
    return float64(minute) * 10
}

func separateHourAndMinute(currentTime float64) (int, int) {
    hour := int(math.Floor(currentTime)) // Extract the hour part
    minute := int((currentTime - float64(hour)) * 60) // Extract the minute part

    if minute == 60 { // Handle edge case where minute equals 60
        hour++
        minute = 0
    }

    return hour, minute
}


