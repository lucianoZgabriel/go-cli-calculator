package history

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// ExportToJSON exporta o histórico para arquivo JSON
func (m *Manager) ExportToJSON(filename string) error {
	if len(m.entries) == 0 {
		return fmt.Errorf("no calculations to export")
	}

	data, err := json.MarshalIndent(m.entries, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %v", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	fmt.Printf("History exported to %s (%d calculations)\n", filename, len(m.entries))
	return nil
}

// ExportToCSV exporta o histórico para arquivo CSV
func (m *Manager) ExportToCSV(filename string) error {
	if len(m.entries) == 0 {
		return fmt.Errorf("no calculations to export")
	}

	var content strings.Builder
	content.WriteString("Expression,Result,Timestamp\n")

	for _, entry := range m.entries {
		content.WriteString(fmt.Sprintf("\"%s\",%.2f,\"%s\"\n",
			entry.Expression, entry.Result, entry.Timestamp))
	}

	err := os.WriteFile(filename, []byte(content.String()), 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	fmt.Printf("History exported to %s (%d calculations)\n", filename, len(m.entries))
	return nil
}

// ExportToTXT exporta o histórico para relatório TXT
func (m *Manager) ExportToTXT(filename string) error {
	if len(m.entries) == 0 {
		return fmt.Errorf("no calculations to export")
	}

	var content strings.Builder
	content.WriteString("Calculator History Report\n")
	content.WriteString(fmt.Sprintf("Generated: %s\n\n", time.Now().Format("2006-01-02 15:04:05")))

	for i, entry := range m.entries {
		content.WriteString(fmt.Sprintf("%d. %s = %.2f [%s]\n",
			i+1, entry.Expression, entry.Result, entry.Timestamp))
	}

	content.WriteString(fmt.Sprintf("\nTotal calculations: %d\n", len(m.entries)))

	err := os.WriteFile(filename, []byte(content.String()), 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	fmt.Printf("History exported to %s (%d calculations)\n", filename, len(m.entries))
	return nil
}
