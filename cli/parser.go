package cli

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseExpression converte uma string como "10 + 5" em números e operação
func ParseExpression(input string) (float64, string, float64, error) {
	// Separar por espaços
	parts := strings.Fields(input)

	// Operações que precisam de 2 números: "10 + 5"
	if len(parts) == 3 {
		num1, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			return 0, "", 0, fmt.Errorf("invalid first number: %v", err)
		}

		operation := parts[1]

		num2, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			return 0, "", 0, fmt.Errorf("invalid second number: %v", err)
		}

		return num1, operation, num2, nil
	}

	// Operações que precisam de 1 número: "sqrt(16)" ou "16 sqrt"
	if len(parts) == 2 {
		num, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			return 0, "", 0, fmt.Errorf("invalid number: %v", err)
		}

		operation := parts[1]
		return num, operation, 0, nil
	}

	return 0, "", 0, fmt.Errorf("invalid expression format")
}

// ConvertSymbolToOperation converte símbolos (+, -, *, /) para nomes (add, sub, mul, div)
func ConvertSymbolToOperation(symbol string) string {
	switch symbol {
	case "+":
		return "add"
	case "-":
		return "sub"
	case "*":
		return "mul"
	case "/":
		return "div"
	case "^":
		return "pow"
	case "√":
		return "sqrt"
	case "add", "sub", "mul", "div", "sqrt", "pow", "sin", "cos":
		return symbol
	default:
		return symbol
	}
}

// GetOperationSymbol converte nomes de operação para símbolos para display
func GetOperationSymbol(op string) string {
	switch op {
	case "add":
		return "+"
	case "sub":
		return "-"
	case "mul":
		return "*"
	case "div":
		return "/"
	case "sqrt":
		return "√"
	case "pow":
		return "^"
	case "sin":
		return "sin"
	case "cos":
		return "cos"
	default:
		return op
	}
}
