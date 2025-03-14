package main

import (
	"fmt"
	_ "math"
)

func main() {
	var A = 5
	var B = 4.5
	S := A + int(B)
	fmt.Println(S)

	var V uint = 14
	var pisyun byte = 25
	GG := V * uint(pisyun)
	fmt.Println(GG)

	H := 13
	J := 9
	HJ := H & J
	/* 13 =1101
	9=1001
	13&9 = 1001=9
	*/
	fmt.Println(HJ)
	SPR := H | J
	fmt.Println(SPR)

}
