package numbers

import (
	"errors"
	"reflect"
	"testing"
)

// TestFactorsList unit test FactorsList function.
func TestFactorsList(t *testing.T) {
	testCases := []struct {
		input    int
		expected []int
		err      error
	}{
		{2, []int{2}, nil},
		{15, []int{3, 5}, nil},
		{26, []int{2, 13}, nil},
		{37, []int{37}, nil},
		{42, []int{2, 3, 7}, nil},
		{-1, []int(nil), errors.New("no prime factors for numbers below two. received '-1'")},
	}
	for _, test := range testCases {
		observed, err := FactorsList(test.input)
		if err != nil {
			if test.err.Error() != err.Error() {
				t.Error(err)
			}
		}
		if !reflect.DeepEqual(observed, test.expected) {
			t.Errorf("for input '%d', expected '%v', got '%v'",
				test.input, test.expected, observed)
		}
	}
}

// BenchmarkFactorList benchmark FactorsList function.
func BenchmarkFactorList(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		_, err := FactorsList(1000)
		if err != nil {
			b.Error(err)
		}
	}
}

// TestIsPrime unit test IsPrime function.
func TestIsPrime(t *testing.T) {
	testCases := []struct {
		input    int
		expected bool
	}{
		{-5, false},
		{0, false},
		{2, true},
		{6, false},
		{13, true},
	}
	for _, test := range testCases {
		observed := IsPrime(test.input)
		if observed != test.expected {
			t.Errorf("for input '%d', expected '%t', got '%t'",
				test.input, test.expected, observed)
		}
	}
}

// BenchmarkIsPrime benchmark IsPrime function.
func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		_ = IsPrime(1728459)
	}
}

// TestFibonacciGen unit test FibonacciGen function.
func TestFibonacciGen(t *testing.T) {
	testCases := []struct {
		input, expected int
	}{
		{1, 0},
		{2, 1},
		{3, 1},
		{4, 2},
		{5, 3},
		{8, 13},
		{10, 34},
		{12, 89},
	}
	for _, test := range testCases {
		var observed int
		c := FibonacciGen()
		for n := 0; n < test.input; n++ {
			observed = <-c
		}
		if observed != test.expected {
			t.Errorf("for input '%d', expected '%d', got '%d'", test.input, test.expected, observed)
		}
	}
}

// BenchmarkFibonacciGen benchmark FibonacciGen function.
func BenchmarkFibonacciGen(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		c := FibonacciGen()
		for n := 0; n < 10000; n++ {
			_ = <-c
		}
	}
}

// TestNthFibonacci unit test NthFibonacci function.
func TestNthFibonacci(t *testing.T) {
	var testCases = []struct {
		input, expected int
	}{
		{1, 1},
		{5, 5},
		{7, 13},
		{10, 55},
		{19, 4181},
	}
	for _, test := range testCases {
		observed := NthFibonacci(test.input)
		if observed != test.expected {
			t.Errorf("for input '%d', expected '%d', got '%d'",
				test.input, test.expected, observed)
		}
	}
}

// BenchmarkNthFibonacci benchmark NthFibonacci function.
func BenchmarkNthFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NthFibonacci(10000)
	}
}

// TestEvenNumbers unit test EvenNumbers function.
func TestEvenNumbers(t *testing.T) {
	testCases := []struct {
		inputA, inputB int
		expected       []int
		err            error
	}{
		{1, 1, []int(nil), nil},
		{1, 3, []int{2}, nil},
		{4, 8, []int{4, 6, 8}, nil},
		{-4, 8, []int{-4, -2, 0, 2, 4, 6, 8}, nil},
		{-1, -2, []int(nil), errors.New("a must be greater than b. received a '-1' and b '-2'")},
	}
	for _, test := range testCases {
		observed, err := EvenNumbers(test.inputA, test.inputB)
		if err != nil {
			if err.Error() != test.err.Error() {
				t.Error(err)
			}
		}
		if !reflect.DeepEqual(observed, test.expected) {
			t.Errorf("for a '%d' and b '%d', expected '%v', got '%v'",
				test.inputA, test.inputB, test.expected, observed)
		}
	}
}

// BenchmarkEvenNumbers benchmark EvenNumbers function.
func BenchmarkEvenNumbers(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		_, err := EvenNumbers(1, 100000)
		if err != nil {
			b.Error(err)
		}
	}
}

// TestOddNumbers unit test OddNumbers function.
func TestOddNumbers(t *testing.T) {
	testCases := []struct {
		inputA, inputB int
		expected       []int
		err            error
	}{
		{1, 1, []int{1}, nil},
		{1, 3, []int{1, 3}, nil},
		{4, 8, []int{5, 7}, nil},
		{-4, 8, []int{-3, -1, 1, 3, 5, 7}, nil},
		{-1, -2, []int(nil), errors.New("a must be greater than b. received a '-1' and b '-2'")},
	}
	for _, test := range testCases {
		observed, err := OddNumbers(test.inputA, test.inputB)
		if err != nil {
			if err.Error() != test.err.Error() {
				t.Error(err)
			}
		}
		if !reflect.DeepEqual(observed, test.expected) {
			t.Errorf("for a '%d' and b '%d', expected '%v', got '%v'",
				test.inputA, test.inputB, test.expected, observed)
		}
	}
}

// BenchmarkOddNumbers benchmark OddNumbers function.
func BenchmarkOddNumbers(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		_, err := OddNumbers(1, 100000)
		if err != nil {
			b.Error(err)
		}
	}
}
