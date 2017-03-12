# numbers
GO package numbers implements useful functions to work with prime numbers, Fibonacci, and parity.

### PACKAGE DOCUMENTATION

```
package numbers

import "github.com/alesr/numbers"

Package numbers implements useful functions to work with prime numbers,
Fibonacci, and parity.
```

### FUNCTIONS

```
func EvenNumbers(a, b int) ([]int, error)
    EvenNumbers returns a list of even numbers from A to B.

func FactorsList(n int) ([]int, error)
    FactorsList lists all prime factors for a positive integer position.
    Does not include repeated factors.

func FibonacciGen() chan int
    FibonacciGen emits an infinite stream of Fibonacci values.

func IsPrime(n int) bool
    IsPrime checks if a given positive int is a prime number.

func NthEven(pos int) (int, error)
    NthEven returns the nth even number for a given position.

func NthFibonacci(pos int) int
    NthFibonacci returns the nth Fibonacci value for given a int position.

func NthOdd(pos int) (int, error)
    NthOdd returns the nth odd number for a given position.

func OddNumbers(a, b int) ([]int, error)
    OddNumbers returns a list of even numbers from A to B.

func SliceByParity(a, b int) (even, odd []int, err error)
    SliceByParity iterates over each n from A to B and add n to a slice of
    odd or even numbers.
```
