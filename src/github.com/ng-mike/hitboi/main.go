package main

import "fmt"

func main() {
	for count := 10; count >= 0; count-- {
		fmt.Printf("%v", count)
		fmt.Printf("%b", count)
		fmt.Printf("%#x", count)
	}
}
