package main
import (
	"fmt"
)

func main()  {
	hello()
	var pogramnum int32
	fmt.Println("Enter to Quit.  1: Fibonacci . 2: Reverse an Integer, 3: Palindrome, 4: Star/number Pyramid ")
	fmt.Scanf("%d",&pogramnum)
	switch pogramnum {
	case 1:
		Fibonacci()
	case 2:
		ReverseInteger()
	case 3:
		Palindrome()
	case 4:
		Pyramid()
	default:
		fmt.Println("Bye. Have a good day.")

	}

}