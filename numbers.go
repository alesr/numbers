package numbers

import (
	"fmt"
	"math"
)

// FactorsList lists all prime factors for a positive integer position.
// Does not include repeated factors.
func FactorsList(n int) ([]int, error) {
	if n < 2 {
		return nil, fmt.Errorf("no prime factors for numbers below two. received '%d'", n)
	}
	if n == 2 {
		return []int{2}, nil
	}
	var factors = make([]int, 0)
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
			i++
		}
	}
	return factors, nil
}

// IsPrime checks if a given positive int is a prime number.
func IsPrime(n int) bool {
	// Considering that primes are greater than 1
	if n <= 0 {
		return false
	}
	// For even, only true if equals to two
	if n%2 == 0 {
		return n == 2
	}
	root := int(math.Sqrt(float64(n)))
	// Tries to divide n for all odd numbers from 3 to n
	for i := 3; i <= root; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
