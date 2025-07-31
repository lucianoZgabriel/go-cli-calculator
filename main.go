package main

import (
	"calculator/calculator"
	"flag"
	"fmt"
	"os"
)

func getOperationSymbol(op string) string {
	switch op {
	case "add":
		return "+"
	case "sub":
		return "-"
	case "mul":
		return "*"
	case "div":
		return "/"
	default:
		return op
	}
}

func printResult(result float64, precision int, verbose bool, num1, num2 float64, operation string) {
	if verbose {
		fmt.Printf("Calculating: %.*f %s %.*f = %.*f\n",
			precision, num1, getOperationSymbol(operation), precision, num2, precision, result)
	} else {
		fmt.Printf("%.*f\n", precision, result)
	}
}

func main() {
	num1 := flag.Float64("a", 0.0, "First number")
	num2 := flag.Float64("b", 0.0, "Second number")
	operation := flag.String("op", "", "Operation: add, sub, mul, div")
	precision := flag.Int("precision", 2, "Number of decimal places (default: 2)")
	verbose := flag.Bool("verbose", false, "Show detailed calculation")
	version := flag.Bool("version", false, "Show version information")

	flag.Parse()

	if *version {
		fmt.Println("Calculator v1.0.0")
		os.Exit(0)
	}

	if flag.NFlag() < 3 {
		fmt.Println("Error: all flags are required")
		fmt.Println("Usage: -a <number> -b <number> -op <operation>")
		flag.Usage()
		os.Exit(1)
	}

	if *operation == "" {
		fmt.Println("Error: operation is required")
		flag.Usage()
		os.Exit(1)
	}

	calc := calculator.New()
	switch *operation {
	case "add":
		result := calc.Add(*num1, *num2)
		printResult(result, *precision, *verbose, *num1, *num2, *operation)
	case "sub":
		result := calc.Subtract(*num1, *num2)
		printResult(result, *precision, *verbose, *num1, *num2, *operation)
	case "mul":
		result := calc.Multiply(*num1, *num2)
		printResult(result, *precision, *verbose, *num1, *num2, *operation)
	case "div":
		result, err := calc.Divide(*num1, *num2)
		if err != nil {
			fmt.Printf("Erro: %v\n", err)
			os.Exit(1)
		}
		printResult(result, *precision, *verbose, *num1, *num2, *operation)
	default:
		fmt.Printf("Operação inválida: %s\n", *operation)
		fmt.Println("Operações válidas: add, sub, mul, div")
		os.Exit(1)
	}
}
