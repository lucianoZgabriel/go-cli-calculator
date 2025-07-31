# CLI Calculator em Go - Planejamento Didático

## 🎯 Objetivos de Aprendizado

### Conceitos Fundamentais de Go
- **Tipos básicos**: int, float64, string, bool
- **Funções**: declaração, parâmetros, retorno múltiplo, error handling
- **Structs**: definição e uso
- **Packages**: organização de código, imports
- **Error handling**: padrão idiomático do Go
- **Testing**: testes unitários com testing package
- **CLI**: argumentos de linha de comando

### Packages da Standard Library
- `fmt` - formatação e I/O
- `os` - argumentos da linha de comando
- `strconv` - conversão de strings
- `flag` - parsing de argumentos CLI
- `errors` - criação de erros customizados
- `testing` - testes unitários

## 📁 Estrutura do Projeto

```
calculator/
├── main.go           # Ponto de entrada
├── calculator/       # Package principal
│   ├── calculator.go # Lógica de cálculo
│   └── calculator_test.go # Testes
├── cli/             # Package para interface CLI
│   ├── parser.go    # Parsing de argumentos
│   └── parser_test.go
├── go.mod           # Módulo Go
└── README.md        # Documentação
```

## 🚀 Fases de Desenvolvimento

### Fase 1: Estrutura Básica e Operações Simples
**Objetivo**: Criar a base do projeto com operações básicas

#### O que implementar:
- [x] Inicializar módulo Go (`go mod init calculator`)
- [x] Criar struct `Calculator` 
- [x] Implementar operações básicas (Add, Subtract, Multiply, Divide)
- [x] Error handling para divisão por zero
- [x] Testes unitários básicos

#### Conceitos praticados:
- Structs e methods
- Error handling idiomático
- Package organization
- Unit testing

#### Código exemplo da estrutura:
```go
// calculator/calculator.go
package calculator

import "errors"

type Calculator struct{}

func New() *Calculator {
    return &Calculator{}
}

func (c *Calculator) Add(a, b float64) float64 {
    return a + b
}

func (c *Calculator) Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

### Fase 2: Interface CLI Básica
**Objetivo**: Criar interface de linha de comando simples

#### O que implementar:
- [x] Parsing de argumentos com `os.Args`
- [x] Validação de entrada
- [x] Formatação de output
- [x] Help message

#### Conceitos praticados:
- Command line arguments
- String manipulation
- Input validation
- User experience

### Fase 3: CLI Avançada com Flag Package
**Objetivo**: Implementar interface mais robusta

#### O que implementar:
- [ ] Usar `flag` package para argumentos nomeados
- [ ] Suporte a diferentes formatos de entrada
- [ ] Verbose mode
- [ ] Version flag

#### Conceitos praticados:
- Flag package
- Boolean flags
- String flags
- Program configuration

### Fase 4: Melhorias e Extensões
**Objetivo**: Adicionar funcionalidades avançadas

#### O que implementar:
- [ ] Operações científicas (sqrt, pow, sin, cos)
- [ ] Modo interativo
- [ ] Histórico de cálculos
- [ ] Export para arquivo

#### Conceitos praticados:
- Math package
- File I/O
- JSON marshaling
- Interactive programs

## 🛠️ Boas Práticas Enfatizadas

### 1. **Naming Conventions**
- Packages: lowercase, single word
- Functions: CamelCase (exported), camelCase (private)
- Constants: CamelCase or ALL_CAPS

### 2. **Error Handling**
```go
// ✅ Correto - sempre verificar erros
result, err := calculator.Divide(10, 0)
if err != nil {
    log.Fatal(err)
}

// ❌ Incorreto - ignorar erros
result, _ := calculator.Divide(10, 0)
```

### 3. **Testing**
- Um test file para cada source file
- Nome dos testes: `TestFunctionName`
- Use table-driven tests quando apropriado

### 4. **Documentation**
- Comentários para funções exportadas
- Package-level documentation

## 📚 Libs Externas Recomendadas (Opcionais)

### Para CLI Avançada:
- **cobra**: Framework para CLI applications
- **viper**: Configuration management

### Para Testing:
- **testify**: Assertions e mocks mais expressivos

### Para Logging:
- **logrus**: Structured logging

## 🎯 Exercícios Práticos por Fase

### Fase 1 - Exercícios:
1. Implementar todas as operações básicas com testes
2. Criar error customizado para cada tipo de erro
3. Adicionar método `String()` ao Calculator (Stringer interface)

### Fase 2 - Exercícios:
1. Validar que pelo menos 3 argumentos foram passados
2. Implementar parsing manual dos argumentos
3. Adicionar support para números negativos

### Fase 3 - Exercícios:
1. Converter CLI para usar flags nomeadas: `-a 10 -b 5 -op add`
2. Adicionar flag `-precision` para controlar casas decimais
3. Implementar `-batch` mode para múltiplos cálculos

### Fase 4 - Exercícios:
1. Adicionar operações científicas usando `math` package
2. Implementar modo interativo com loop
3. Salvar histórico em JSON

## 🧪 Casos de Teste Importantes

```go
func TestCalculator_Divide(t *testing.T) {
    tests := []struct {
        name    string
        a, b    float64
        want    float64
        wantErr bool
    }{
        {"normal division", 10, 2, 5, false},
        {"division by zero", 10, 0, 0, true},
        {"negative numbers", -10, 2, -5, false},
    }
    
    calc := New()
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := calc.Divide(tt.a, tt.b)
            if (err != nil) != tt.wantErr {
                t.Errorf("Divide() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("Divide() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

## 📖 Comandos Importantes

```bash
# Inicializar projeto
go mod init calculator

# Executar testes
go test ./...

# Executar com verbose
go test -v ./...

# Build
go build

# Executar
./calculator 10 + 5

# Com flags (Fase 3)
./calculator -a 10 -b 5 -op add
```

## 🎓 Próximos Passos

Após completar este projeto, você terá praticado:
- Organização de código em packages
- Error handling idiomático
- Testing em Go
- CLI development
- Estruturas de dados básicas

**Pronto para começar? Vamos implementar a Fase 1!** 🚀