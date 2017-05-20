package main

import "fmt"
import "time"

func main() {
	for count := 10; count >= 0; count-- {
		fmt.Printf("%v\n", count)
		fmt.Printf("%b\n", count)
		fmt.Printf("%#x\n", count)
		time.Sleep(time.Second)
	}
}
