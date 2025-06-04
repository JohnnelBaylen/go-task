package main

import "fmt"

func main() {

	var number int

	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
