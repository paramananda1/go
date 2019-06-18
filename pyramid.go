package main

import "fmt"

func Pyramid() {
	var number int = 0
	fmt.Println("Enter to Quit. Enter a number of rows of Pyramid")

	fmt.Scanf("%d",&number)

	fmt.Println("Star Pyramid")
	for i:=0 ; i<=number ; i++ {
		for j:=0 ; j<=i; j++{
			fmt.Printf("*  ")
		}
		fmt.Println()
	}

	fmt.Println("Number Pyramid 1")
	for i:=1 ; i<=number ; i++ {
		for j:=1 ; j<=i; j++{
			fmt.Printf("%d  ",j)
		}
		fmt.Println()
	}

	fmt.Println("Number Pyramid 2 ")
	k := 1
	for i:=1 ; i<=number ; i++ {
		for j:=1 ; j<=i; j++{
			fmt.Printf("%d  ",k)
			k++
		}
		fmt.Println()
	}
}
