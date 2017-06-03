package main

import "fmt"

var pgkScope = 0

func main() {
<<<<<<< HEAD
	for count := 10; count >= 0; count-- {
		fmt.Printf("%v\n", count)
		if count == 0 {			
			fmt.Printf("HitBoi!")
		}
		time.Sleep(time.Second)
=======
	incPgkScope := func() int {
		pgkScope++
		return pgkScope
	}
	funcScope := incPgkScope()
	count := func() int {
		funcScope++
		return funcScope
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("pgkScope value is: %v\n", pgkScope)
		fmt.Printf("funcScope value is: %v\n", funcScope)
		fmt.Printf("%v\n", count())

>>>>>>> 748fcef5a6ceab8479e415f26bcbae737bd4858a
	}

}
