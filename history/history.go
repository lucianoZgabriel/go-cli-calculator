package history

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Entry representa uma entrada no histórico de cálculos
type Entry struct {
	Expression string  `json:"expression"`
	Result     float64 `json:"result"`
	Timestamp  string  `json:"timestamp"`
}

// Manager gerencia o histórico de cálculos
type Manager struct {
	entries  []Entry
	filename string
}

// NewManager cria um novo gerenciador de histórico
func NewManager(filename string) *Manager {
	m := &Manager{
		filename: filename,
		entries:  []Entry{},
	}
	m.Load() // Carrega histórico existente
	return m
}

// Add adiciona uma nova entrada ao histórico
func (m *Manager) Add(expression string, result float64) {
	entry := Entry{
		Expression: expression,
		Result:     result,
		Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
	}
	m.entries = append(m.entries, entry)
	m.Save() // Salva automaticamente
}

// Show exibe o histórico
func (m *Manager) Show() {
	if len(m.entries) == 0 {
		fmt.Println("No calculation in history")
		return
	}

	fmt.Println("\n--- Calculation History ---")
	for i, entry := range m.entries {
		fmt.Printf("%d. %s = %.2f [%s]\n", i+1, entry.Expression, entry.Result, entry.Timestamp)
	}
	fmt.Println()
}

// GetEntries retorna todas as entradas (para export)
func (m *Manager) GetEntries() []Entry {
	return m.entries
}

// Load carrega o histórico do arquivo
func (m *Manager) Load() {
	data, err := os.ReadFile(m.filename)
	if err != nil {
		// Arquivo não existe ou erro ao ler - começar com histórico vazio
		m.entries = []Entry{}
		return
	}

	err = json.Unmarshal(data, &m.entries)
	if err != nil {
		fmt.Printf("Warning: Could not load history: %v\n", err)
		m.entries = []Entry{}
	}
}

// Save salva o histórico no arquivo
func (m *Manager) Save() {
	data, err := json.MarshalIndent(m.entries, "", "  ")
	if err != nil {
		fmt.Printf("Warning: Could not save history: %v\n", err)
		return
	}

	err = os.WriteFile(m.filename, data, 0644)
	if err != nil {
		fmt.Printf("Warning: Could not write history file: %v\n", err)
	}
}
