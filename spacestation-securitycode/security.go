package main

func Fibonacci(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func CodePower(n int) int {
	if n == 0 {
		return 1
	}
	power := n * n
	return power
}

func PrimeCode(power int) bool {
	if power < 0 {
		return false
	}
	if power <= 1 {
		return false
	}
	for i := 2; i*i <= power; i++ {
		if power%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	for PrimeCode(nb) == false {
		nb++
	}
	return nb
}

func FinalSecurityCode(n int) int {
	fib := Fibonacci(n)

	squared := CodePower(fib)

	if PrimeCode(squared) {
		return squared
	}

	return FindNextPrime(squared)
}
