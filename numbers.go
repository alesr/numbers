package numbers

import "fmt"

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
