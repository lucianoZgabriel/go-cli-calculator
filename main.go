package main

import (
	"calculator/calculator"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Uso: ./calculator <num1> <operação> <num2>")
		fmt.Println("Operações: +, -, *, /")
		os.Exit(1)
	}

	num1, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Printf("Número inválido: %s\n", os.Args[1])
		os.Exit(1)
	}

	operation := os.Args[2]

	num2, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Printf("Número inválido: %s\n", os.Args[3])
		os.Exit(1)
	}

	calc := calculator.New()
	switch operation {
	case "+":
		result := calc.Add(num1, num2)
		fmt.Printf("%.2f\n", result)
	case "-":
		result := calc.Subtract(num1, num2)
		fmt.Printf("%.2f\n", result)
	case "*":
		result := calc.Multiply(num1, num2)
		fmt.Printf("%.2f\n", result)
	case "/":
		result, err := calc.Divide(num1, num2)
		if err != nil {
			fmt.Printf("Erro: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%.2f\n", result)
	default:
		fmt.Printf("Operação inválida: %s\n", operation)
		fmt.Println("Operações válidas: +, -, *, /")
		os.Exit(1)
	}
}
