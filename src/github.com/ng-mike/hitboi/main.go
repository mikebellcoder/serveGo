package main

import (
	"fmt"
	"time"
)

func main() {
	hour := time.Now().Hour()

	var getGreeting = func(t int) string {
		var greet string
		if t < 12 {
			greet = "Good morning."
		} else if t > 12 {
			greet = "Good Afternoon."

		} else if t < 18 {
			greet = "Good Evening."
		}
		return greet
	}

	greeting := getGreeting(hour)
	fmt.Println(greeting)

}
