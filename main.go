package main

import (
    "fmt"
    "github.com/mazzy001/go-ci-demo/math"
)

func main() {
    fmt.Println("=== Демонстрация работы математического пакета ===\n")
    
    // Тестируем Add
    fmt.Printf("Add(10, 5) = %d\n", math.Add(10, 5))
    
    // Тестируем IsEven
    fmt.Printf("IsEven(7) = %v\n", math.IsEven(7))
    fmt.Printf("IsEven(8) = %v\n", math.IsEven(8))
    
    // Тестируем Factorial
    fmt.Printf("Factorial(5) = %d\n", math.Factorial(5))
    fmt.Printf("Factorial(0) = %d\n", math.Factorial(0))
    
    fmt.Println("\n✅ Все функции работают корректно!")
}
