package ya2

import "fmt"

var Global = 5

func globalDeferManip() {
	fmt.Println("Start.")
	fmt.Println("Global = ", Global)
	oldGlobal := Global

	defer func() { Global = oldGlobal }()

	Global++
	fmt.Println("Global++; Global = ", Global)
	Global++
	fmt.Println("Global++; Global = ", Global)
	Global++
	fmt.Println("Global++; Global = ", Global)

	fmt.Println("End.")
}
