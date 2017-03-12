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
