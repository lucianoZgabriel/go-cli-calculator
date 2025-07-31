package cli

import (
	"bufio"
	"calculator/calculator"
	"fmt"
	"os"
	"strings"
)

// HistoryManager é uma interface para gerenciar histórico
// (vamos implementar isso depois quando criarmos o package history)
type HistoryManager interface {
	Add(expression string, result float64)
	Show()
}

// RunInteractiveMode executa o modo interativo da calculadora
func RunInteractiveMode(historyMgr HistoryManager) {
	fmt.Println("Calculator v1.0.0 - Interactive Mode")
	fmt.Println("Type 'exit' to quit, 'history' to show calculation history")
	fmt.Println("Usage: number operation number (e.g., 10 + 5)")
	fmt.Println()

	calc := calculator.New()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "exit" {
			fmt.Println("Goodbye")
			break
		}
		if input == "" {
			continue
		}
		if input == "history" {
			historyMgr.Show()
			continue
		}

		num1, operation, num2, err := ParseExpression(input)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		operation = ConvertSymbolToOperation(operation)
		switch operation {
		case "add":
			result := calc.Add(num1, num2)
			historyMgr.Add(input, result)
			fmt.Printf("%.2f\n", result)
		case "sub":
			result := calc.Subtract(num1, num2)
			historyMgr.Add(input, result)
			fmt.Printf("%.2f\n", result)
		case "mul":
			result := calc.Multiply(num1, num2)
			historyMgr.Add(input, result)
			fmt.Printf("%.2f\n", result)
		case "div":
			result, err := calc.Divide(num1, num2)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}
			historyMgr.Add(input, result)
			fmt.Printf("%.2f\n", result)
		case "sqrt":
			result, err := calc.Sqrt(num1)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}
			historyMgr.Add(input, result)
			fmt.Printf("%.2f\n", result)
		case "pow":
			result := calc.Pow(num1, num2)
			historyMgr.Add(input, result)
			fmt.Printf("%.2f\n", result)
		case "sin":
			result := calc.Sin(num1)
			historyMgr.Add(input, result)
			fmt.Printf("%.2f\n", result)
		case "cos":
			result := calc.Cos(num1)
			historyMgr.Add(input, result)
			fmt.Printf("%.2f\n", result)
		default:
			fmt.Printf("Unknown operation: %s\n", operation)
			fmt.Println("Supported: +, -, *, /, sqrt, pow, sin, cos")
		}
	}
}
