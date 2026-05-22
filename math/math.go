package math

// Add складывает два числа.
func Add(a, b int) int {
    return a + b
}

// IsEven проверяет, является ли число четным.
func IsEven(n int) bool {
    return n%2 == 0
}

// Factorial вычисляет факториал числа (n!).
func Factorial(n int) int {
    if n < 0 {
        return 0 // Для отрицательных чисел возвращаем 0
    }
    if n == 0 {
        return 1
    }
    result := 1
    for i := 1; i <= n; i++ {
        result *= i
    }
    return result
}
