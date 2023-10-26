package main

import (
	"fmt"
)

func calculate(a, b int, operator string, resultChan chan int) {
	var result int
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	}

	resultChan <- result
}

func main() {
	operations := []struct {
		a        int
		b        int
		operator string
	}{
		{2, 2, "+"},
		{3, 6, "+"},
		{7, 7, "*"},
		{9, 3, "/"},
	}

	resultChan := make(chan int)

	for _, operation := range operations {
		go calculate(operation.a, operation.b, operation.operator, resultChan)
	}

	for range operations {
		result := <-resultChan
		fmt.Println("Result:", result)
	}
}
