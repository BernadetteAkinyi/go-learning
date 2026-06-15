package main

import "fmt"

func Fibonacci(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(n int) int {
	if n <= 2 {
		return 2
	}
	for !IsPrime(n) {
		n++
	}
	return n
}

func PrimeFibonacci(index int) int {
	if index < 0 {
		return -1
	}

	fibo := Fibonacci(index)
	return FindNextPrime(fibo)
}

func main() {
	fmt.Println(PrimeFibonacci(0))
	fmt.Println(PrimeFibonacci(4))
	fmt.Println(PrimeFibonacci(6))
	fmt.Println(PrimeFibonacci(8))
}
