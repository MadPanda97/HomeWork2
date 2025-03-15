package main

import (
	"fmt"
	_ "fmt"
	_ "math"
)

func main() {
	if A := 4; A == 1 {
		fmt.Println("one")
	} else if A == 2 {
		fmt.Println("two")
	} else if A == 3 {
		fmt.Println("three")
	} else if A == 4 {
		fmt.Println("four")
	} else if A == 5 {
		fmt.Println("five")
	} else if A == 6 {
		fmt.Println("six")
	} else if A == 7 {
		fmt.Println("seven")
	} else if A == 8 {
		fmt.Println("eight")
	} else if A == 9 {
		fmt.Println("nine")
	} else if A == 10 {
		fmt.Println("ten")
	} else if A == 0 {
		fmt.Println("zero")
	}

	W1 := 5
	L1 := 10
	W2 := 6
	L2 := 9

	if W1 > W2 && L1 > L2 {
		println("Вмещает")
	} else {
		println("Не вмещает")
	}

	pow := -5
	if pow > 0 {
		var N = 10
		var result = N
		for P := 0; P < pow-1; P++ {
			result = result * N

		}
		{
			println(result)
		}
	} else if pow == 0 {
		var N = 1
		var result = N
		for P := 0; P < pow-1; P++ {
			result = result * N
		}
		{
			println(result)
		}
	} else if pow < 0 {
		var N float64 = 10
		var result float64 = N
		for P := 0; P > pow-1; P-- {
			result = result / N
		}
		{
			println(result)
		}
	}

}
