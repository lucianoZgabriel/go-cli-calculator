# CLI Calculator em Go

Uma calculadora de linha de comando desenvolvida em Go para aprendizado da linguagem.

## 🚀 Como usar

```bash
# Build do projeto
go build -o calculatorV1

# Executar operações
./calculatorV1 10 + 5
./calculatorV1 15.5 / 3
./calculatorV1 8 * 2
./calculatorV1 20 - 7
```

## 🧪 Executar testes

```bash
go test ./...
```

## 📁 Estrutura do projeto

- `main.go` - Ponto de entrada da aplicação
- `calculator/` - Package com lógica de cálculo
- `go_calculator_plan.md` - Planejamento do desenvolvimento

## 🎯 Status do projeto

- ✅ Fase 1: Operações básicas implementadas
- ✅ Fase 2: Interface CLI básica
- ⏳ Fase 3: CLI avançada com flags
- ⏳ Fase 4: Funcionalidades extras

## 🛠️ Tecnologias utilizadas

- Go 1.24
- Testing nativo do Go
- Standard library (fmt, os, strconv)