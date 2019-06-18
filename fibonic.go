package main
import "fmt"

func Fibonacci() {
	var n, next int
	var first , second int = 0,1

	fmt.Println("Enter the number of terms:");
	fmt.Scanf("%d", &n);

	fmt.Printf("First %d terms of Fibonacci series are:\n", n);

	for i := 0; i < n; i++ {
		if (i <= 1){
			next = i
		}else {
			next = first + second
			first = second
			second = next
		}
		fmt.Printf("%d   ", next);
	}
	fmt.Println()

}