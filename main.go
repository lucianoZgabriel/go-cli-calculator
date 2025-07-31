package main

import (
	"calculator/cli"
	"calculator/history"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Definir flags
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

	// Inicializar gerenciador de histórico
	historyMgr := history.NewManager("calculator_history.json")

	// Processar flags especiais
	if *version {
		fmt.Println("Calculator v1.0.0")
		return
	}

	if *interactive {
		cli.RunInteractiveMode(historyMgr)
		return
	}

	if *showHist {
		historyMgr.Show()
		return
	}

	if *exportJSON != "" {
		if err := historyMgr.ExportToJSON(*exportJSON); err != nil {
			fmt.Printf("Export failed: %v\n", err)
			os.Exit(1)
		}
		return
	}

	if *exportCSV != "" {
		if err := historyMgr.ExportToCSV(*exportCSV); err != nil {
			fmt.Printf("Export failed: %v\n", err)
			os.Exit(1)
		}
		return
	}

	if *exportTXT != "" {
		if err := historyMgr.ExportToTXT(*exportTXT); err != nil {
			fmt.Printf("Export failed: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// Modo CLI normal
	cli.RunCLIMode(historyMgr, *num1, *num2, *operation, *precision, *verbose)
}
