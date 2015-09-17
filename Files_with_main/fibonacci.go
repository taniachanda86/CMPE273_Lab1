
package main

import (
	"fmt"
	)

func fibonacci(x int) int{

	// var dict map[int]int
	// dict[0] = 0
	// dict[1] = 1

	if x==0{
		return 0
	} else if x == 1{
		return 1
	} else {
		return fibonacci(x-1) + fibonacci(x-2)
	}
}


func main() {

	var num int
	//for i := 3; i>0; i--{
		//fmt.Print("You have ", i ," tries.")
		fmt.Print("Please enter a positive Integer to find its Fibonacci: ")
		fmt.Scanf("%d", &num)
		if num < 0{
			fmt.Print("Its not a positive number.")
		} else {
			fmt.Println(fibonacci(num))
		}
	//}	
}