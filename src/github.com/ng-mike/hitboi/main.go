package main

import "fmt"

var pgkScope = 0

func main() {
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

	}

}
