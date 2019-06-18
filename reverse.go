package main

import (
	"fmt"
)


func ReverseInteger() {
	var number, rnumber uint64 = 0,0
	fmt.Println("Enter the Number want to Reverse: ")
	fmt.Scanf("%d",&number)

	for ; number>0 ; {
		mod := number%10
		number = number/10
		rnumber = (rnumber*10) + mod
	}
	fmt.Println("After Reverse: ",rnumber)
}