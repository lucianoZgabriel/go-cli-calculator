package main

import (
	"bufio"
	"calculator/calculator"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type HistoryEntry struct {
	Expression string  `json:"expression"`
	Result     float64 `json:"result"`
	Timestamp  string  `json:"timestamp"`
}

var calculationHistory []HistoryEntry

const historyFile = "calculator_history.json"

func loadHistory() {
	data, err := os.ReadFile(historyFile)
	if err != nil {
		// Arquivo não existe ou erro ao ler - começar com histórico vazio
		calculationHistory = []HistoryEntry{}
		return
	}

	err = json.Unmarshal(data, &calculationHistory)
	if err != nil {
		fmt.Printf("Warning: Could not load history: %v\n", err)
		calculationHistory = []HistoryEntry{}
	}
}

func saveHistory() {
	data, err := json.MarshalIndent(calculationHistory, "", " ")
	if err != nil {
		fmt.Printf("Warning: Could not save history: %v\n", err)
		return
	}

	err = os.WriteFile(historyFile, data, 0644)
	if err != nil {
		fmt.Printf("Warning: Could not write history file: %v\n", err)
	}
}

func addToHistory(expression string, result float64) {
	entry := HistoryEntry{
		Expression: expression,
		Result:     result,
		Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
	}
	calculationHistory = append(calculationHistory, entry)
	saveHistory()
}

func showHistory() {
	if len(calculationHistory) == 0 {
		fmt.Println("No calculation in history")
		return
	}

	fmt.Println("\n--- Calculation History ---")
	for i, entry := range calculationHistory {
		fmt.Printf("%d. %s = %.2f [%s]\n", i+1, entry.Expression, entry.Result, entry.Timestamp)
	}
	fmt.Println()
}

func exportToJSON(filename string) {
	if len(calculationHistory) == 0 {
		fmt.Println("No calculations to export.")
		return
	}

	data, err := json.MarshalIndent(calculationHistory, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}

	fmt.Printf("History exported to %s (%d calculations)\n", filename, len(calculationHistory))
}

func exportToCSV(filename string) {
	if len(calculationHistory) == 0 {
		fmt.Println("No calculations to export.")
		return
	}

	var content strings.Builder
	content.WriteString("Expression,Result,Timestamp\n")

	for _, entry := range calculationHistory {
		content.WriteString(fmt.Sprintf("\"%s\",%.2f,\"%s\"\n",
			entry.Expression, entry.Result, entry.Timestamp))
	}

	err := os.WriteFile(filename, []byte(content.String()), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}

	fmt.Printf("History exported to %s (%d calculations)\n", filename, len(calculationHistory))
}

func exportToTXT(filename string) {
	if len(calculationHistory) == 0 {
		fmt.Println("No calculations to export.")
		return
	}

	var content strings.Builder
	content.WriteString("Calculator History Report\n")
	content.WriteString(fmt.Sprintf("Generated: %s\n\n", time.Now().Format("2006-01-02 15:04:05")))

	for i, entry := range calculationHistory {
		content.WriteString(fmt.Sprintf("%d. %s = %.2f [%s]\n",
			i+1, entry.Expression, entry.Result, entry.Timestamp))
	}

	content.WriteString(fmt.Sprintf("\nTotal calculations: %d\n", len(calculationHistory)))

	err := os.WriteFile(filename, []byte(content.String()), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}

	fmt.Printf("History exported to %s (%d calculations)\n", filename, len(calculationHistory))
}

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

func printResult(result float64, precision int, verbose bool, num1, num2 float64, operation string) {
	if verbose {
		fmt.Printf("Calculating: %.*f %s %.*f = %.*f\n",
			precision, num1, getOperationSymbol(operation), precision, num2, precision, result)
	} else {
		fmt.Printf("%.*f\n", precision, result)
	}
}

func convertSymbolToOperation(symbol string) string {
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

func parseExpression(input string) (float64, string, float64, error) {
	parts := strings.Fields(input)

	// Operações que precisam de 2 números
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

	// Operações que precisam de 1 número
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

func runInteractiveMode() {
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
			showHistory()
			continue
		}

		num1, operation, num2, err := parseExpression(input)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		operation = convertSymbolToOperation(operation)
		switch operation {
		case "add":
			result := calc.Add(num1, num2)
			addToHistory(input, result)
			fmt.Printf("%.2f\n", result)
		case "sub":
			result := calc.Subtract(num1, num2)
			addToHistory(input, result)
			fmt.Printf("%.2f\n", result)
		case "mul":
			result := calc.Multiply(num1, num2)
			addToHistory(input, result)
			fmt.Printf("%.2f\n", result)
		case "div":
			result, err := calc.Divide(num1, num2)
			addToHistory(input, result)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}
			fmt.Printf("%.2f\n", result)
		case "sqrt":
			result, err := calc.Sqrt(num1)
			addToHistory(input, result)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}
			fmt.Printf("%.2f\n", result)
		case "pow":
			result := calc.Pow(num1, num2)
			addToHistory(input, result)
			fmt.Printf("%.2f\n", result)
		case "sin":
			result := calc.Sin(num1)
			addToHistory(input, result)
			fmt.Printf("%.2f\n", result)
		case "cos":
			result := calc.Cos(num1)
			addToHistory(input, result)
			fmt.Printf("%.2f\n", result)
		default:
			fmt.Printf("Unknown operation: %s\n", operation)
			fmt.Println("Supported: +, -, *, /, sqrt, pow, sin, cos")
		}
	}
}

func main() {
	num1 := flag.Float64("a", 0.0, "First number")
	num2 := flag.Float64("b", 0.0, "Second number")
	operation := flag.String("op", "", "Operation: add, sub, mul, div, sqrt, pow, sin, cos")
	precision := flag.Int("precision", 2, "Number of decimal places (default: 2)")
	verbose := flag.Bool("verbose", false, "Show detailed calculation")
	version := flag.Bool("version", false, "Show version information")
	interactive := flag.Bool("interactive", false, "Start interactive mode")
	showHist := flag.Bool("show-history", false, "Show calculation history")
	exportJSON := flag.String("export-json", "", "Export history to JSON file")
	exportCSV := flag.String("export-csv", "", "Export history to CSV file")
	exportTXT := flag.String("export-txt", "", "Export history to TXT report")

	flag.Parse()

	loadHistory()

	if *version {
		fmt.Println("Calculator v1.0.0")
		os.Exit(0)
	}

	if *interactive {
		runInteractiveMode()
		return
	}

	if *showHist {
		showHistory()
		return
	}

	if *exportJSON != "" {
		exportToJSON(*exportJSON)
		return
	}

	if *exportCSV != "" {
		exportToCSV(*exportCSV)
		return
	}

	if *exportTXT != "" {
		exportToTXT(*exportTXT)
		return
	}

	if flag.NFlag() < 2 {
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
		addToHistory(fmt.Sprintf("%.2f + %.2f", *num1, *num2), result)
		printResult(result, *precision, *verbose, *num1, *num2, *operation)
	case "sub":
		result := calc.Subtract(*num1, *num2)
		addToHistory(fmt.Sprintf("%.2f + %.2f", *num1, *num2), result)
		printResult(result, *precision, *verbose, *num1, *num2, *operation)
	case "mul":
		result := calc.Multiply(*num1, *num2)
		addToHistory(fmt.Sprintf("%.2f + %.2f", *num1, *num2), result)
		printResult(result, *precision, *verbose, *num1, *num2, *operation)
	case "div":
		result, err := calc.Divide(*num1, *num2)
		if err != nil {
			fmt.Printf("Erro: %v\n", err)
			os.Exit(1)
		}
		addToHistory(fmt.Sprintf("%.2f + %.2f", *num1, *num2), result)
		printResult(result, *precision, *verbose, *num1, *num2, *operation)
	case "sqrt":
		result, err := calc.Sqrt(*num1)
		if err != nil {
			fmt.Printf("Erro: %v\n", err)
			os.Exit(1)
		}
		addToHistory(fmt.Sprintf("sqrt(%.2f)", *num1), result)
		printResult(result, *precision, *verbose, *num1, 0, *operation)
	case "pow":
		result := calc.Pow(*num1, *num2)
		addToHistory(fmt.Sprintf("%.2f ^ %.2f", *num1, *num2), result)
		printResult(result, *precision, *verbose, *num1, *num2, *operation)
	case "sin":
		result := calc.Sin(*num1)
		addToHistory(fmt.Sprintf("sin(%.2f)", *num1), result)
		printResult(result, *precision, *verbose, *num1, 0, *operation)
	case "cos":
		result := calc.Cos(*num1)
		addToHistory(fmt.Sprintf("cos(%.2f)", *num1), result)
		printResult(result, *precision, *verbose, *num1, 0, *operation)
	default:
		fmt.Printf("Operação inválida: %s\n", *operation)
		fmt.Println("Operações válidas: add, sub, mul, div, sqrt, pow, sin, cos")
		os.Exit(1)
	}
}
