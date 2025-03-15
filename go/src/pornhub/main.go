package main

import (
	_ "fmt"
	_ "math"
)

func main() {
	
	var (
		nFib1 = 1
		nFib2 = 1 + 1
	)

	println(nFib1)
	println(nFib2)

	for n := 0; n < 30; n++ {
		nFib1 = nFib2 + nFib1
		println(nFib1)
		nFib2 = nFib2 + nFib1
		println(nFib2)
	}
}
