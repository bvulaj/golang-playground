package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var num int
	flag.IntVar(&num, "num", -1, "List all prime numbers up to this number. Must be a positive number. Required.")
	flag.Parse()

	if num < 1 {
		fmt.Println("Invalid input.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Println("Getting primes (1) for", num, "...")
	fmt.Println(getPrimes(num))

	fmt.Println("Getting primes (2) for", num, "...")
	fmt.Println(getPrimes2(num))
}

func getPrimes(num int) []int {
	primes := []int{2}
	for i := 3; i < num; i += 2 {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func isPrime(num int) bool {
	for i := 2; i < num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func getPrimes2(num int) []int {
	primes := []int{2}
	for i := 3; i < num; i += 2 {
		if isPrime2(i, primes) {
			primes = append(primes, i)
		}
	}
	return primes
}
func isPrime2(num int, primes []int) bool {
	for _, prvPrime := range primes {
		if num%prvPrime == 0 {
			return false
		}
	}
	return true
}
