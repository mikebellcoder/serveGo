package main

import "fmt"
import "time"

func main() {
	for count := 20; count >= 0; count-- {
		fmt.Printf("%v\n", count)
		fmt.Printf("%b\n", count)
		fmt.Printf("%#x\n", count)
		if count == 0 {
			fmt.Printf("HitBoi")
		}
		time.Sleep(time.Second)
	}
}
