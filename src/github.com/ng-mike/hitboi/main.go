<<<<<<< HEAD
package main

import "fmt"
import "time"

func main() {
	for count := 20; count >= 0; count-- {
		fmt.Printf("%v\n", count)
		fmt.Printf("%b\n", count)
		fmt.Printf("%#x\n", count)
		if count == 0 {
			fmt.Printf("Take off!")
		}
		time.Sleep(time.Second)
	}
}
=======
package main

import "fmt"
import "time"

func main() {
	for count := 20; count >= 0; count-- {
		fmt.Printf("%v\n", count)
		fmt.Printf("%b\n", count)
		fmt.Printf("%#x\n", count)
		if count == 0 {
<<<<<<< HEAD
			fmt.Printf("Baby!")
=======
			fmt.Printf("HitBoi")
			fmt.Printf("HitBoi")
>>>>>>> save from core on Self Branch
		}
		time.Sleep(time.Second)
	}
}
>>>>>>> self
