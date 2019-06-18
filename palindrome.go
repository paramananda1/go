package main

import (
	"fmt"
	"strings"
)

func Palindrome() {

	var number uint64 = 0
	var str string = ""

	fmt.Println("Enter to Quit. Enter a Number to check if it is a palindrome")
	fmt.Scanf("%d",&number)

	if number != 0 {
		var rnumber , temp uint64 = 0,number
		for ; temp>0 ; {
			mod := temp%10
			temp = temp/10
			rnumber = (rnumber*10) + mod
		}
		if number == rnumber {
			fmt.Println("Number is a Palindrome ")
		}else {
			fmt.Println("Number is NOT Palindrome ")
		}

	}

	fmt.Println("Enter to Quit. Enter a string to check if it is a palindrome")
	fmt.Scanf("%s",&str)
	if str != "" {
		var rstr string


		for _,v := range str {
			rstr = string(v) + rstr
		}

		if  strings.Compare(str,rstr) == 0 {
			fmt.Println("String is a Palindrome ")
		}else {
			fmt.Println("String is NOT Palindrome ")
		}

	}

}

