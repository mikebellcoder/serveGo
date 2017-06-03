package main

import "fmt"
import "time"

func main() {
	for count := 10; count >= 0; count-- {
		fmt.Printf("%v\n", count)
		if count == 0 {			
			fmt.Printf("HitBoi!")
		}
		time.Sleep(time.Second)
	}
}
