package main

import "fmt"

func isPrime(n int) bool {

	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func findPrimes(start, end int, results chan int) {

	for i := start; i <= end; i++ {

		if isPrime(i) {
			results <- i
		}
	}

	close(results)
}

func main() {

	start, end := 1, 100
	results := make(chan int)

	go findPrimes(start, end, results)

	fmt.Println("Prime numbers : ")

	for prime := range results {
		fmt.Println(prime)
	}
}
