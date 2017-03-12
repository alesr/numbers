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

// FibonacciGen emits an infinite stream of Fibonacci values.
func FibonacciGen() chan int {
	c := make(chan int)
	go func() {
		for i, j := 0, 1; ; i, j = i+j, i {
			c <- i
		}
	}()
	return c
}

// NthFibonacci returns the nth Fibonacci value for given a int position.
func NthFibonacci(pos int) int {
	c := FibonacciGen()
	f := make([]int, pos)
	for i := 0; i <= pos; i++ {
		f = append(f, <-c)
	}
	return f[len(f)-1]
}

// EvenNumbers returns a list of even numbers from A to B.
func EvenNumbers(a, b int) ([]int, error) {
	even, _, err := SliceByParity(a, b)
	if err != nil {
		return nil, err
	}
	return even, nil
}

// OddNumbers returns a list of even numbers from A to B.
func OddNumbers(a, b int) ([]int, error) {
	_, odd, err := SliceByParity(a, b)
	if err != nil {
		return nil, err
	}
	return odd, nil
}

// NthEven returns the nth even number for a given position.
func NthEven(pos int) (int, error) {
	if pos < 0 {
		return 0, fmt.Errorf("input must be a positive number. received '%d'", pos)
	}
	even, _ := EvenNumbers(0, (pos*2)-1)
	return even[len(even)-1], nil
}
