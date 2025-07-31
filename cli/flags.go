package cli

import (
	"calculator/calculator"
	"fmt"
	"os"
)

// HistoryManager interface (já definida em interactive.go, mas boa prática ter aqui também)
// Você pode remover de interactive.go se quiser

// RunCLIMode executa um cálculo único via flags de linha de comando
func RunCLIMode(historyMgr HistoryManager, num1, num2 float64, operation string, precision int, verbose bool) {
	// Validação básica
	if operation == "" {
		fmt.Println("Error: operation is required")
		os.Exit(1)
	}

	// Criar calculadora
	calc := calculator.New()

	// Executar operação
	switch operation {
	case "add":
		result := calc.Add(num1, num2)
		expression := fmt.Sprintf("%.2f + %.2f", num1, num2)
		historyMgr.Add(expression, result)
		printResult(result, precision, verbose, num1, num2, operation)

	case "sub":
		result := calc.Subtract(num1, num2)
		expression := fmt.Sprintf("%.2f - %.2f", num1, num2)
		historyMgr.Add(expression, result)
		printResult(result, precision, verbose, num1, num2, operation)

	case "mul":
		result := calc.Multiply(num1, num2)
		expression := fmt.Sprintf("%.2f * %.2f", num1, num2)
		historyMgr.Add(expression, result)
		printResult(result, precision, verbose, num1, num2, operation)

	case "div":
		result, err := calc.Divide(num1, num2)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		expression := fmt.Sprintf("%.2f / %.2f", num1, num2)
		historyMgr.Add(expression, result)
		printResult(result, precision, verbose, num1, num2, operation)

	case "sqrt":
		result, err := calc.Sqrt(num1)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		expression := fmt.Sprintf("sqrt(%.2f)", num1)
		historyMgr.Add(expression, result)
		printResult(result, precision, verbose, num1, 0, operation)

	case "pow":
		result := calc.Pow(num1, num2)
		expression := fmt.Sprintf("%.2f ^ %.2f", num1, num2)
		historyMgr.Add(expression, result)
		printResult(result, precision, verbose, num1, num2, operation)

	case "sin":
		result := calc.Sin(num1)
		expression := fmt.Sprintf("sin(%.2f)", num1)
		historyMgr.Add(expression, result)
		printResult(result, precision, verbose, num1, 0, operation)

	case "cos":
		result := calc.Cos(num1)
		expression := fmt.Sprintf("cos(%.2f)", num1)
		historyMgr.Add(expression, result)
		printResult(result, precision, verbose, num1, 0, operation)

	default:
		fmt.Printf("Invalid operation: %s\n", operation)
		fmt.Println("Valid operations: add, sub, mul, div, sqrt, pow, sin, cos")
		os.Exit(1)
	}
}

// printResult formata e exibe o resultado
func printResult(result float64, precision int, verbose bool, num1, num2 float64, operation string) {
	if verbose {
		fmt.Printf("Calculating: %.*f %s %.*f = %.*f\n",
			precision, num1, GetOperationSymbol(operation), precision, num2, precision, result)
	} else {
		fmt.Printf("%.*f\n", precision, result)
	}
}
