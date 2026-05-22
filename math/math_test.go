package math

import "testing"

// TestAdd проверяет сложение.
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    if result != expected {
        t.Errorf("Add(2,3) = %d; want %d", result, expected)
    }
}

// TestIsEven проверяет четность.
func TestIsEven(t *testing.T) {
    if !IsEven(4) {
        t.Error("IsEven(4) = false; want true")
    }
    if IsEven(5) {
        t.Error("IsEven(5) = true; want false")
    }
}

// TestFactorial проверяет факториал.
func TestFactorial(t *testing.T) {
    tests := []struct {
        name string
        n    int
        want int
    }{
        {"Factorial of 0", 0, 1},
        {"Factorial of 5", 5, 120},
        {"Factorial of 3", 3, 6},
        {"Negative number", -1, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Factorial(tt.n); got != tt.want {
                t.Errorf("Factorial(%d) = %d; want %d", tt.n, got, tt.want)
            }
        })
    }
}
