# 🧮 Calculator CLI em Go

Uma calculadora avançada de linha de comando desenvolvida em Go, implementando tanto **modo CLI** quanto **modo interativo**, com histórico persistente e múltiplos formatos de export.

## ✨ Funcionalidades

### 🎯 **Dois Modos de Operação**
- **Modo CLI**: Execução direta via flags (`./calculatorV1 -a 10 -b 5 -op add`)  
- **Modo Interativo**: Interface conversacional (`./calculatorV1 -interactive`)

### 🔢 **Operações Suportadas**
- **Básicas**: Adição, Subtração, Multiplicação, Divisão
- **Científicas**: Raiz quadrada, Potenciação, Seno, Cosseno
- **Error Handling**: Divisão por zero, números negativos em raiz quadrada

### 📊 **Histórico e Export**
- **Persistência**: Histórico salvo automaticamente em JSON
- **Visualização**: Comando `history` e flag `-show-history`
- **Export**: Múltiplos formatos (JSON, CSV, TXT)

### ⚙️ **Configurações Avançadas**
- **Precisão**: Controle de casas decimais (`-precision`)
- **Modo Verbose**: Cálculos detalhados (`-verbose`)
- **Versioning**: Flag `-version`

## 🚀 Como usar

### **Instalação**
```bash
git clone https://github.com/seu-usuario/calculator-go
cd calculator-go
go build -o calculatorV1
```

### **Modo CLI**
```bash
# Operações básicas
./calculatorV1 -a 10 -b 5 -op add
./calculatorV1 -a 15.5 -b 3 -op div -precision 4

# Operações científicas  
./calculatorV1 -a 16 -op sqrt
./calculatorV1 -a 2 -b 3 -op pow
./calculatorV1 -a 90 -op sin

# Com modo verbose
./calculatorV1 -a 10 -b 5 -op add -verbose
# Output: Calculating: 10.00 + 5.00 = 15.00
```

### **Modo Interativo**
```bash
./calculatorV1 -interactive

> 10 + 5
15.00
> 16 sqrt
4.00  
> 2 ^ 3
8.00
> history
--- Calculation History ---
1. 10 + 5 = 15.00 [2024-01-15 14:30:25]
2. 16 sqrt = 4.00 [2024-01-15 14:31:10] 
3. 2 ^ 3 = 8.00 [2024-01-15 14:32:05]

> exit
Goodbye!
```

### **Histórico e Export**
```bash
# Ver histórico
./calculatorV1 -show-history

# Export em diferentes formatos
./calculatorV1 -export-json "meu_historico.json"
./calculatorV1 -export-csv "planilha.csv"  
./calculatorV1 -export-txt "relatorio.txt"
```

### **Flags Disponíveis**
```bash
-a float          First number
-b float          Second number  
-op string        Operation: add, sub, mul, div, sqrt, pow, sin, cos
-precision int    Number of decimal places (default: 2)
-verbose          Show detailed calculation
-interactive      Start interactive mode
-show-history     Show calculation history
-export-json      Export history to JSON file
-export-csv       Export history to CSV file
-export-txt       Export history to TXT report
-version          Show version information
```

## 🧪 Executar testes

```bash
# Todos os testes
go test ./...

# Com verbose
go test -v ./...

# Com coverage
go test -cover ./...

# Testes específicos
go test ./calculator
```

## 📁 Arquitetura do Projeto

```
calculator/
├── main.go                    # Ponto de entrada mínimo (80 linhas)
├── calculator/
│   ├── calculator.go          # Lógica matemática core
│   └── calculator_test.go     # Testes unitários abrangentes
├── cli/
│   ├── parser.go             # Parsing de expressões ("10 + 5")
│   ├── interactive.go        # Modo interativo com loop
│   └── flags.go              # Modo CLI com validação de flags
├── history/
│   ├── history.go            # Gerenciamento de histórico
│   └── export.go             # Export JSON/CSV/TXT
├── go.mod                    # Módulo Go
└── README.md                 # Esta documentação
```

### **Princípios Arquiteturais**
- **Separação de Responsabilidades**: Cada package tem uma função específica
- **Interface-Based Design**: `HistoryManager` permite desacoplamento  
- **Error Handling Idiomático**: Retorno de erros seguindo padrões Go
- **Testabilidade**: Arquitetura permite testes unitários eficazes

## 🎯 Jornada de Desenvolvimento

Este projeto foi desenvolvido em **4 fases** para aprendizado progressivo de Go:

### **✅ Fase 1: Fundação**
- Struct `Calculator` com operações básicas
- Testes unitários com table-driven tests  
- Error handling para divisão por zero
- Package organization inicial

### **✅ Fase 2: Interface CLI Básica** 
- Parsing de argumentos com `os.Args`
- Validação de entrada robusta
- Formatação de output consistente
- Help messages informativos

### **✅ Fase 3: CLI Avançada**
- Refatoração para `flag` package
- Flags obrigatórias e opcionais (`-precision`, `-verbose`)
- Validação avançada de argumentos
- Flag `-version` e help automático

### **✅ Fase 4: Funcionalidades Avançadas**
- **Operações científicas** usando `math` package
- **Modo interativo** com `bufio.Scanner` e parsing customizado
- **Histórico persistente** com JSON marshaling/unmarshaling  
- **Export multi-formato** (JSON, CSV, TXT reports)

### **✅ Refatoração Profissional**
- Reorganização em packages bem definidos
- Redução do `main.go` de 350+ para 80 linhas
- Implementação de interfaces para desacoplamento
- Arquitetura seguindo best practices Go

## 🛠️ Tecnologias e Conceitos

### **Standard Library Go**
- `fmt` - Formatação e I/O
- `os` - Argumentos CLI e file operations  
- `flag` - Parsing avançado de argumentos
- `strconv` - Conversão de tipos
- `math` - Operações científicas
- `encoding/json` - Serialização de dados
- `bufio` - I/O bufferizado para modo interativo
- `strings` - Manipulação de strings
- `time` - Timestamps e formatação

### **Conceitos Go Praticados** 
- **Structs e Methods** com pointer receivers
- **Interfaces** para design desacoplado
- **Error Handling** idiomático com múltiplos retornos
- **Package Organization** e encapsulamento
- **Testing** com table-driven tests e coverage
- **JSON Marshaling/Unmarshaling** para persistência
- **File I/O** para histórico e exports
- **CLI Parsing** com flag package
- **String Processing** para modo interativo

## 🏆 Resultados de Aprendizado

Após completar este projeto, você dominou:

- ✅ **Sintaxe Go** completa e idiomática
- ✅ **Arquitetura de Software** com separação clara de responsabilidades  
- ✅ **Testing** abrangente e metodologias TDD
- ✅ **CLI Development** profissional com múltiplas interfaces
- ✅ **Data Persistence** com JSON e múltiplos formatos
- ✅ **Error Handling** robusto seguindo best practices
- ✅ **Package Design** para reutilização e manutenibilidade

## 📈 Estatísticas do Projeto

- **Linhas de Código**: ~800 linhas
- **Packages**: 4 (main, calculator, cli, history)  
- **Funções/Métodos**: 25+
- **Testes Unitários**: Cobertura completa das operações
- **Funcionalidades**: 12+ recursos implementados
- **Formatos Suportados**: 3 tipos de export

## 🎓 Próximos Passos

Possíveis extensões para continuar aprendendo:

- **Testes de Integração** com cenários end-to-end
- **Benchmarking** para otimização de performance  
- **CLI Frameworks** (Cobra/Viper) para funcionalidades avançadas
- **Concorrência** com goroutines para operações paralelas
- **APIs REST** para interface web
- **Containerização** com Docker
- **CI/CD** com GitHub Actions

---

**Desenvolvido com ❤️ em Go para aprendizado prático da linguagem**

*Este projeto demonstra uma progressão estruturada do básico ao avançado em Go, seguindo as melhores práticas da comunidade e servindo como referência para desenvolvimento profissional.*